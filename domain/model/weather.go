package model

type WeatherResponse struct {
  Coord struct{
    Lon string `json:"lon"`
    Lat string `json:"lat"`
  } `json:"coord"`
  Weather[] struct{
    Id int `json:"id"`
    Main string `json:"main"`
    Description string `json:"description"`
  } `json:"weather"`
  Base string `json:"stations"`
  Main struct {
    Temp float32 `json:"temp"`
    FeelsLike string `json:"feels_like"`
    TempMin float32 `json:"temp_min"`
    TempMax float32 `json:"temp_max"`
    Pressure int `json:"pressure"`
    Humidity int `json:"humidity"`
  } `json:"main"`
  Visibility string `json:"visibility"`
  Wind struct{
    Speed float32 `json:"string"`
  }`json:"wind"`
  Clouds struct{
    All int `json:"all"`
  } `json:"clouds"`
  Name string `json:"name"`
}