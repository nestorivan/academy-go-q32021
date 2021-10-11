package repository

import (
	"encoding/csv"
	"fmt"
	"os"

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
  r := csv.NewReader(f)

  pkmList := []model.Pokemon{}

  values,err := r.ReadAll()

  if err != nil {
    return nil, err
  }

  for _, p := range values {
    pm := model.Pokemon{
      Id: p[0],
      Name: p[1],
    }
    pkmList = append(pkmList, pm)
  }

  return pkmList, nil
}

func (pr *pokemonRepo) Get(id string) ([]model.Pokemon, error){
  f, err := pr.fileService.ReadCsv("pokemon.csv")

  if err != nil {
    return nil, err
  }

  pkmList, _ := pr.getPokemonFromFile(f)

  fmt.Println("hello from repo")

  //close file after func closes
  defer f.Close()

  return pkmList, nil
}

func (pr *pokemonRepo) Save(p model.Pokemon) ([]model.Pokemon, error){
  f, err := pr.fileService.ReadCsv("pokemon.csv")

  if (err != nil){
    fmt.Println(err)
    return nil, err
  }

  pkmList, _ := pr.getPokemonFromFile(f)

  w := csv.NewWriter(f)

  for _, p := range pkmList {
    var row []string
    row = append(row, p.Id)
    row = append(row, p.Name)
    w.Write(row)
  }

  w.Flush()

  defer func(){
    f.Close()
  }()

  return pkmList, nil
}