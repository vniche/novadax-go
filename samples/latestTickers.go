package main

import (
	"log"

	novadax "github.com/vniche/novadax-go"
)

func main() {
	client := novadax.Default()

	tickers, err := client.GetLatestTickers(&novadax.GetLatestMarketTickersFilters{})
	if err != nil {
		log.Printf("%s", err.Error())
	}

	log.Printf("%+v", tickers)
	for _, ticker := range tickers {
		log.Printf("%+v", ticker)
	}
}
