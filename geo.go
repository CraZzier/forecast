package main

import (
	"io"
	"log"
	"net/http"

	geojson "github.com/paulmach/go.geojson"
)

func mustFindCoordinatesForCity(city string) []float64 {
	res, err := http.Get("http://photon.komoot.io/api/?q=" + city)
	if err != nil {
		log.Fatalln(err)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}
	fc, err := geojson.UnmarshalFeatureCollection(body)
	if err != nil {
		log.Fatalln(err)
	}
	var cityCoords []float64
	for _, feature := range fc.Features {
		//I am getting only first entry because if i am about to show forecast for one city first on the list will be the most probable
		if feature.Properties["countrycode"] == "US" {
			cityCoords = append(cityCoords, feature.Geometry.Point...)
			cityCoords[0], cityCoords[1] = cityCoords[1], cityCoords[0]
			return cityCoords
		}
	}
	log.Fatalln("Could not find US city")
	return nil
}
