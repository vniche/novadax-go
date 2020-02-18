package main

import (
	"log"

	novadax "github.com/artemis-tech/novadax-go"
)

func main() {
	client := novadax.Default()

	tickers, err := client.GetLatestTickers(&novadax.GetLatestTickersFilters{})
	if err != nil {
		log.Printf("%s", err.Error())
	}

	log.Printf("%+v", tickers)
	for _, ticker := range tickers {
		log.Printf("%+v", ticker)
	}
}
