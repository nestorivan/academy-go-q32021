package interactor

import (
	"github.com/nestorivan/academy-go-q32021/domain/model"
	"github.com/nestorivan/academy-go-q32021/repository"
)

type pokemonInteractor struct {
	PokemonRepo repository.PokemonRepo
}

type PokemonInteractor interface{
  Get(id string) ([]model.Pokemon, error)
  Create(pkm model.Pokemon) error
}

func NewPokemonInteractor(pr repository.PokemonRepo) PokemonInteractor {
  return &pokemonInteractor{pr}
}

func (pi *pokemonInteractor) Get(id string) ([]model.Pokemon, error) {
  pkmList, err := pi.PokemonRepo.Get(id)

  if err != nil{
    return nil, err
  }

  return pkmList, nil
}

func (pi *pokemonInteractor) Create(pkm model.Pokemon) error {
  _, err := pi.PokemonRepo.Save(pkm)

  if err != nil{
    return err
  }

  return nil
}