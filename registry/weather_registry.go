package registry

import (
	"github.com/nestorivan/academy-go-q32021/controller"
	"github.com/nestorivan/academy-go-q32021/interactor"
	"github.com/nestorivan/academy-go-q32021/repository"
	"github.com/nestorivan/academy-go-q32021/service"
)

func (r *registry) NewWeatherController() controller.WeatherController {
  return controller.NewWeatherController(r.NewWeatherInteractor())
}

func (r *registry) NewWeatherInteractor() interactor.WeatherInteractor {
  fs := r.NewFileService()
  wr := r.NewWeatherRepository(fs)
  return interactor.NewWeatherInteractor(wr)
}

func (r *registry) NewWeatherRepository(fs service.FileService) repository.WeatherRepository {
  return repository.NewWeatherRepository(fs)
}