package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type WeatherData struct {
	Name string `json:"name"`
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
}

func getWeather(apiKey, city string) (WeatherData, error) {
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city, apiKey)
	
	resp, err := http.Get(url)
	if err != nil {
		return WeatherData{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return WeatherData{}, fmt.Errorf("HTTP request failed with status code %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return WeatherData{}, err
	}

	var weatherData WeatherData
	if err := json.Unmarshal(body, &weatherData); err != nil {
		return WeatherData{}, err
	}

	return weatherData, nil
}

func main() {
	apiKey := "99b1e77c0771ded0db3ab13394f4faf0"
	city := "New York"

	weatherData, err := getWeather(apiKey, city)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("City: %s\n", weatherData.Name)
		fmt.Printf("Temperature: %.2f°C\n", weatherData.Main.Temp - 273.15)
	}
}
