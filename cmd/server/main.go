package main

import (
	"log"
	"net/url"
	"strconv"

	"github.com/CrutoiAlexandru/weather-app/cmd/weather"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/weather", func(c *gin.Context) {
		var values url.Values = c.Request.URL.Query()

		latitude, err := strconv.ParseFloat(values["latitude"][0], 32)
		if err != nil {
			log.Println(err)
		}
		longitude, err := strconv.ParseFloat(values["longitude"][0], 32)
		if err != nil {
			log.Println(err)
		}

		data, err := weather.GetWeather(latitude, longitude)
		if err != nil {
			c.JSON(400, gin.H{})
		}

		c.JSON(200, gin.H{
			"data": data,
		})
	})

	r.Run("0.0.0.0:8080")
}
