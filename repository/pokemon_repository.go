package repository

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/jszwec/csvutil"
	"github.com/nestorivan/academy-go-q32021/domain/model"
	"github.com/nestorivan/academy-go-q32021/service"
)

type pokemonRepo struct{
  fileService service.FileService
}

type PokemonRepo interface {
  Get(id string) ([]model.Pokemon, error)
  Save(p model.Pokemon) ([]model.Pokemon, error)
  GetAsync(ap model.AsyncParams) ([]model.Pokemon, error)
}

// NewPokemonRepo - returns a pokemonRepo pointer
func NewPokemonRepo(fi service.FileService) PokemonRepo {
  return &pokemonRepo{fi}
}

// getPokemonFromFile - reads a csv file and returns a pokemon list
func (pr *pokemonRepo) getPokemonFromFile(f *os.File) ([]model.Pokemon, error){
  var pkmList []model.Pokemon

  csvReader := csv.NewReader(f)

  dec, err := csvutil.NewDecoder(csvReader)
	if err != nil {
		log.Fatal(err)
	}

	for {
		var p model.Pokemon
		if err := dec.Decode(&p); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		pkmList = append(pkmList, p)
	}

  return pkmList, nil
}

// sortPokemonAsync - private method that checks if pokemon id is odd or even and pipes result into res channel
func (pr *pokemonRepo) sortPokemonAsync(f *os.File, jobs <- chan model.Pokemon, res chan <- model.Pokemon, ap *model.AsyncParams ){
  itemCounter := 0

	for pkm := range jobs{
    if(itemCounter == ap.ItemsPerWorker){
      break
    }
    if ap.Type == "even" && pkm.Id%2 == 0 {
      res <- pkm
      itemCounter++
    } else if ap.Type == "odd" && pkm.Id%2 == 1{
      res <- pkm
      itemCounter++
    }
  }
}

// Get - Returns pokemon list
func (pr *pokemonRepo) Get(id string) ([]model.Pokemon, error){
  f, err := pr.fileService.ReadFile("pokemon.csv")

  if err != nil {
    return nil, err
  }

  pkmList, _ := pr.getPokemonFromFile(f)

  //close file after func closes
  defer f.Close()

  return pkmList, nil
}

// Save - Saves a new pokemon into csv file
func (pr *pokemonRepo) Save(p model.Pokemon) ([]model.Pokemon, error){
  f, err := pr.fileService.ReadFile("pokemon.csv")

  if (err != nil){
    return nil, err
  }

  pkmList, _ := pr.getPokemonFromFile(f)

  w := csv.NewWriter(f)

  for _, p := range pkmList {
    var row []string
    row = append(row, strconv.Itoa(p.Id))
    row = append(row, p.Name)
    w.Write(row)
  }

  w.Flush()

  defer func(){
    f.Close()
  }()

  return pkmList, nil
}

// GetAsync - Returns Pokemon List that complies with params constraints
func (pr *pokemonRepo) GetAsync(ap model.AsyncParams) ([]model.Pokemon, error){
  f, err := pr.fileService.ReadFile("pokemon.csv")
  csvReader := csv.NewReader(f)

  if err != nil {
    return nil, err
  }

  var pkmList []model.Pokemon

  //items per worker
  numWkrs := ap.Items / ap.ItemsPerWorker
  //numWkrs
  w := make(chan model.Pokemon, numWkrs)
  res := make(chan model.Pokemon)

  var wg sync.WaitGroup


  // when channel receives pokemon execute worker function
  for i := 0; i < numWkrs; i++ {
    wg.Add(1)

    go func(){
      defer wg.Done()
      pr.sortPokemonAsync(f, w, res, &ap)
    }()
  }


  // iterate pokemons from csv
  // send each pokemon to workers channel
  go func(){
    dec, err := csvutil.NewDecoder(csvReader)

    if err != nil {
      log.Fatal(err)
    }

    // while pokemon list is less than required items send pokemon to worker channel
    for len(pkmList) < ap.Items {
      var p model.Pokemon
      if err := dec.Decode(&p); err == io.EOF {
        break
      } else if err != nil {
        log.Fatal(err)
        break
      }
      w <- p
    }
    // close worker channel when done
    close(w)
  }()

  defer f.Close()

  // wait for all workers of wait group end
  // close res channel when done
  go func(){
    wg.Wait()
    close(res)
  }()

  // add all pokemon that we get from response channel
  for pk := range res {
    pkmList = append(pkmList,pk)
  }

  return pkmList, nil
}