package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nestorivan/academy-go-q32021/domain/model"
	"github.com/nestorivan/academy-go-q32021/interactor"
)


type pokemonController struct {
  PokemonInteractor interactor.PokemonInteractor
}

type PokemonController interface {
  GetPokemons() gin.HandlerFunc
  CreatePokemons() gin.HandlerFunc
}

func NewPokemonController(pi interactor.PokemonInteractor) PokemonController{
  return &pokemonController{pi}
}

func (pk *pokemonController) GetPokemons() gin.HandlerFunc {
  return func(c *gin.Context) {

    // p := c.Request.URL.Query()

    id := c.Param("id")
    pkml, err := pk.PokemonInteractor.Get(id)

    if (err != nil){
      c.AbortWithStatus(http.StatusInternalServerError)
    }

    if (id == ""){
      c.JSON(http.StatusOK, pkml)
      return
    }

    pkm := model.Pokemon{};

    for _, p := range pkml{
      if (strconv.Itoa(p.Id) == id){
        pkm = p
      }
    }

    c.JSON(http.StatusOK, pkm)
  }
}

func (pk *pokemonController) CreatePokemons() gin.HandlerFunc {
  return func(c *gin.Context) {
    var pkmn model.Pokemon

    err := c.Bind(&pkmn)

    if (err != nil){
      c.Status(http.StatusBadRequest)
    }

    pk.PokemonInteractor.Create(pkmn)

    c.Status(http.StatusOK)
  }
}