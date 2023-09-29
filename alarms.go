package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gen2brain/beeep"
	geojson "github.com/paulmach/go.geojson"
)

// City Eloy tested for alarm
func getAlarmsForCoordinates(coords []float64) []*geojson.Feature {
	fmt.Println("Alarm list:")
	res, err := http.Get(fmt.Sprintf("https://api.weather.gov/alerts/active?point=%f,%f", coords[0], coords[1]))
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fc, err := geojson.UnmarshalFeatureCollection(body)
	if err != nil {
		log.Fatal(err)
	}
	return fc.Features
}

func printAlarms(alarms []*geojson.Feature) {
	for _, alarm := range alarms {
		//Displaying selected fields
		log.Println("Alarm:")
		log.Println("Headline: ", alarm.Properties["headline"])
		log.Println("Description: ", alarm.Properties["description"])
		log.Println("Instruction: ", alarm.Properties["instruction"])
		log.Println("Severity: ", alarm.Properties["severity"])
		log.Println("Urgency: ", alarm.Properties["urgency"])
		log.Println("Certainty: ", alarm.Properties["certainty"])
		log.Println("Event: ", alarm.Properties["event"])

		//System warning
		beeep.Notify("Weather Alarm", alarm.Properties["headline"].(string), "")
	}
}
