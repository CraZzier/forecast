package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	geojson "github.com/paulmach/go.geojson"
)

func mustGetHourlyForecast(coords []float64) []WeatherHourlyData {
	res, err := http.Get(fmt.Sprintf("http://api.weather.gov/points/%f,%f", coords[0], coords[1]))
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	f, err := geojson.UnmarshalFeature(body)
	if err != nil {
		log.Fatal(err)
	}

	//Gettting final forecast link
	fHourly, ok := f.Properties["forecastHourly"].(string)
	if !ok {
		log.Println("Could not find forecastHourly")
	}
	res, err = http.Get(fHourly)
	if err != nil {
		log.Fatal(err)
	}
	body, err = io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	f, err = geojson.UnmarshalFeature(body)
	if err != nil {
		log.Fatal(err)
	}

	periodsPlain, ok := f.Properties["periods"]
	if !ok {
		log.Fatal("Could not find periods")
	}
	jsonData, err := json.Marshal(periodsPlain)
	if err != nil {
		log.Fatal(err)
	}

	hourlyData := []WeatherHourlyData{}
	json.Unmarshal(jsonData, &hourlyData)
	return hourlyData
}

func printCurrentForecast(data []WeatherHourlyData, city string) {
	currentTime := time.Now().UTC()
	for _, data := range data {
		if data.StartTime.Before(currentTime) && data.EndTime.After(currentTime) {
			log.Println("Current forecast:", city, " ", data.StartTime, " to ", data.EndTime)
			log.Println("Tempreature: ", data.Temperature, data.TemperatureUnit)
			log.Println("Wind: ", data.WindSpeed, data.WindDirection)
			log.Println("Chance of rain:", data.ProbabilityOfPrecip.Value, strings.Split(data.ProbabilityOfPrecip.UnitCode, ":")[1])
		}
	}
}
