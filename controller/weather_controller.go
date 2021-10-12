package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nestorivan/academy-go-q32021/interactor"
)

type weatherController struct{
  WeatherInteractor interactor.WeatherInteractor
}

type WeatherController interface {
  GetWeather() gin.HandlerFunc
}

// NewWeatherController - Returns a new WeatherController instance
func NewWeatherController(wi interactor.WeatherInteractor) WeatherController{
  return &weatherController{wi}
}

// GetWeather - return the weather for the provided city
func (wc *weatherController) GetWeather() gin.HandlerFunc {
  return func(c *gin.Context){
    city := c.Param("city")

    if city == "" {
      c.Status(http.StatusBadRequest)
    }

    w,err := wc.WeatherInteractor.Get(city)

    if err !=nil {
      c.Status(http.StatusInternalServerError)
    }

    c.JSON(http.StatusOK, w)

  }
}