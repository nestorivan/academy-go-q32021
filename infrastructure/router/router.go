package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nestorivan/academy-go-q32021/controller"
)


func NewRouter(c controller.AppController) *gin.Engine {
  r := gin.Default()

  r.Handle(http.MethodGet, "/pokemon/", c.Pokemon.GetPokemons())
  r.Handle(http.MethodGet, "/pokemon/:id", c.Pokemon.GetPokemons())
  r.Handle(http.MethodPost, "/pokemon", c.Pokemon.CreatePokemons())

  r.Handle(http.MethodGet, "weather/:city", c.Weather.GetWeather())


  return r
}