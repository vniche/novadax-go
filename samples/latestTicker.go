package main

import (
	"log"

	novadax "github.com/vniche/novadax-go"
)

func main() {
	client := novadax.Default()

	ticker, err := client.GetMarketTicker(&novadax.GetMarketTickersFilters{})
	if err != nil {
		log.Printf("%s", err.Error())
	}

	log.Printf("%+v", ticker)
}
