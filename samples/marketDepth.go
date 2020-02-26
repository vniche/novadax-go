package main

import (
	"log"

	novadax "github.com/artemis-tech/novadax-go"
)

func main() {
	client := novadax.Default()

	depth, err := client.GetMarketDepth(&novadax.GetMarketDepthFilters{
		Symbol: "ADA_ETH",
	})
	if err != nil {
		log.Printf("%s", err.Error())
	}

	log.Printf("%+v", depth)
	log.Println("asks")
	for _, ask := range depth.Asks {
		log.Printf("price: %s amount: %s", ask[0], ask[1])
	}
	log.Println("bids")
	for _, bid := range depth.Bids {
		log.Printf("price: %s amount: %s", bid[0], bid[1])
	}
}
