package main

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"log"
	"os"
)

func main() {
	logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
	logger.Println("Start app...")

	c := colly.NewCollector()
	url := "https://api.open-meteo.com/v1/forecast" +
		"?latitude=52.52" +
		"&longitude=13.41" +
		"&current=temperature_2m,wind_speed_10m" +
		"&hourly=temperature_2m,relative_humidity_2m,wind_speed_10m"

	c.OnScraped(func(r *colly.Response) {
		logger.Println("Finished scraping data")
		body := r.Body
		fmt.Println(string(body))
	})

	err := c.Visit(url)
	if err != nil {
		logger.Fatal(err)
	}

	logger.Println("End app...")
}
