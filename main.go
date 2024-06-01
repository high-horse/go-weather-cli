package main

import (
	"io"
	"net/http"
)

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