package controller

type AppController struct {
  Pokemon interface{ PokemonController }
  Weather interface{ WeatherController }
}