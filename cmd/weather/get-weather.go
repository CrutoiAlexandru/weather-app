package weather

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func GetWeather(latitude float64, longitude float64) (string, error) {
	apiEndpoint := "https://api.open-meteo.com/v1/forecast?"
	apiEndpoint = apiEndpoint + fmt.Sprintf("latitude=%f&", latitude)
	apiEndpoint = apiEndpoint + fmt.Sprintf("longitude=%f&", longitude)
	apiEndpoint = apiEndpoint + "temperature_unit=celsius&" + "current=temperature_2m&" + "hourly=temperature_2m"

	log.Println(apiEndpoint)

	response, err := http.Get(apiEndpoint)
	if err != nil {
		log.Println(err)
	}

	jsonData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
	}

	return string(jsonData), nil
}
