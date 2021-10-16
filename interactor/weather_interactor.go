package interactor

import (
	"github.com/nestorivan/academy-go-q32021/domain/model"
	"github.com/nestorivan/academy-go-q32021/repository"
)

type weatherInteractor struct {
  WeatherRepo repository.WeatherRepository

}

type WeatherInteractor interface {
  Get(c string) (model.WeatherResponse, error)
}

func NewWeatherInteractor(wr repository.WeatherRepository) WeatherInteractor{
  return &weatherInteractor{wr}
}

func (wi *weatherInteractor) Get(city string) (model.WeatherResponse, error) {
  w, err := wi.WeatherRepo.GetWeather(city)

  if err != nil{
    return w, err
  }

  return w, nil
}
// func (wi weatherInteractor) Save (model.WeatherResponse, error) {}
