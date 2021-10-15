package repository

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

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
}

func NewPokemokemonRepo(fi service.FileService) PokemonRepo {
  return &pokemonRepo{fi}
}

func (pr *pokemonRepo) getPokemonFromFile(f *os.File) ([]model.Pokemon, error){
  var pkmList []model.Pokemon

  fb, _ := ioutil.ReadAll(f)

	if err := csvutil.Unmarshal(fb, &pkmList); err != nil {
		return nil, err
	}

  return pkmList, nil
}

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

func (pr *pokemonRepo) Save(p model.Pokemon) ([]model.Pokemon, error){
  f, err := pr.fileService.ReadFile("pokemon.csv")

  if (err != nil){
    fmt.Println(err)
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