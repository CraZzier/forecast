package main

import "time"

type WeatherHourlyData struct {
	Number              int           `json:"number"`
	Name                string        `json:"name"`
	StartTime           time.Time     `json:"startTime"`
	EndTime             time.Time     `json:"endTime"`
	IsDaytime           bool          `json:"isDaytime"`
	Temperature         int           `json:"temperature"`
	TemperatureUnit     string        `json:"temperatureUnit"`
	TemperatureTrend    interface{}   `json:"temperatureTrend"`
	ProbabilityOfPrecip Precipitation `json:"probabilityOfPrecipitation"`
	Dewpoint            Measurement   `json:"dewpoint"`
	RelativeHumidity    Measurement   `json:"relativeHumidity"`
	WindSpeed           string        `json:"windSpeed"`
	WindDirection       string        `json:"windDirection"`
	Icon                string        `json:"icon"`
	ShortForecast       string        `json:"shortForecast"`
	DetailedForecast    string        `json:"detailedForecast"`
}

type Precipitation struct {
	UnitCode string `json:"unitCode"`
	Value    int    `json:"value"`
}

type Measurement struct {
	UnitCode string  `json:"unitCode"`
	Value    float64 `json:"value"`
}
