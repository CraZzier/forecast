package main

import (
	"fmt"
	"os"
	"time"

	"github.com/inancgumus/screen"
)

func main() {

	//Simple Handling for inproper input
	if len(os.Args) == 1 {
		fmt.Println("Please enter a city name")
		return
	}
	if len(os.Args) > 2 {
		fmt.Println("Please enter only one city name")
		return
	}
	chosenCity := os.Args[1]

	//STEP 1: Find coordinates for chosen city
	cityCoords := mustFindCoordinatesForCity(chosenCity)

	for {
		//STEP 2: Get hourly forecast for chosen city
		hourlyData := mustGetHourlyForecast(cityCoords)

		//STEP 3: Print forecast for current hour
		//But first clear console
		screen.Clear()
		screen.MoveTopLeft()
		printCurrentForecast(hourlyData, chosenCity)

		//STEP 3: Print alarms for current hour
		alarmList := getAlarmsForCoordinates(cityCoords)
		printAlarms(alarmList)
		time.Sleep(1 * time.Minute)
	}

}
