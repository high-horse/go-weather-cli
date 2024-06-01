package main

import (
	"io"
	"net/http"
)

type Weather struct {
	Location struct {
		Name string 	`json:"name"`
		Country string 	`json:"country"`
	}`json:"location"`
	Current struct {
		TempC float64 	`json:"temp_c"`
		Condition struct{
			Text string `json:"text"`
		} 	`json:"condition"`
	}`json:"current"`
	Forecast struct {
		Forecastday []struct {
			Hour []struct {
				TimeEpoch int64 `json:"time_epoch"`
				TempC float64 	`json:"temp_c"`
				Condition struct{
					Text string `json:"text"`
				} 	`json:"condition"`
				ChanceOfRain float64 	`json:"chance_of_rain"`

			} `json:"hour"`
		}`json:"Forecastday"`
	} `json:"forecast"`
}

func main() {
	res, err := http.Get("http://api.weatherapi.com/v1/forecast.json?key=db03a6cec6f8441ba52173233240106&q=london")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	// println(res.StatusCode )
	if res.StatusCode != 200 {
		panic("api not available")
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
println(string(body))
	println("ready to go")
}