package repository

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/nestorivan/academy-go-q32021/domain/model"
	"github.com/nestorivan/academy-go-q32021/service"
)

type weatherRepository struct{
  FileService service.FileService
}

type WeatherRepository interface{
  GetWeather(city string) (model.WeatherResponse, error)
  saveWeather(model.WeatherResponse) (error)
}

func NewWeatherRepository(fs service.FileService) WeatherRepository{
  return &weatherRepository{fs}
}

func (wr *weatherRepository) GetWeather(city string) (model.WeatherResponse, error){
  client := &http.Client{}
    req, _ := http.NewRequest(http.MethodGet, "http://api.openweathermap.org/data/2.5/weather" , nil)

    q := req.URL.Query()
    q.Add("q", city)
    q.Add("appid", "77d5640c224071461a75e8d8cd373ad9")
    req.URL.RawQuery = q.Encode()

    fmt.Println("url",req.URL.String())

    res, err := client.Do(req)

    var w model.WeatherResponse

    if err != nil{
      return w , err
    }

    rb, _ := ioutil.ReadAll(res.Body)

    json.Unmarshal(rb, &w)

    wr.saveWeather(w)

    return w, nil
}

func (wr *weatherRepository) saveWeather(w model.WeatherResponse) (error){
  fp := "weather.json"
  data, err := json.MarshalIndent(w, "", " ")

  if err != nil {
    return err
  }

  ioutil.WriteFile(fp, data, 0666)

  return nil
}