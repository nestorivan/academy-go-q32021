package registry

import (
	"github.com/nestorivan/academy-go-q32021/controller"
	"github.com/nestorivan/academy-go-q32021/interactor"
	presenter "github.com/nestorivan/academy-go-q32021/presenters"
	"github.com/nestorivan/academy-go-q32021/repository"
	"github.com/nestorivan/academy-go-q32021/service"
)

func (r *registry) NewPokemonController() controller.PokemonController {
  return controller.NewPokemonController(r.NewPokemonInteractor())
}

func (r *registry) NewPokemonInteractor() interactor.PokemonInteractor {
  fs := r.NewFileService()
  pkr := r.NewPokemonRepository(fs)
  return interactor.NewPokemonInteractor(pkr)
}

func (r *registry) NewPokemonPresenter() presenter.PokemonPresenter {
  return presenter.NewPokemonPresenter()
}

func (r *registry) NewPokemonRepository(fs service.FileService) repository.PokemonRepo {
  return repository.NewPokemonRepo(fs)
}