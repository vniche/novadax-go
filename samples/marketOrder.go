package main

import (
	"fmt"
	"log"

	novadax "github.com/artemis-tech/novadax-go"
)

func main() {
	client := novadax.Default()

	buyOrder, err := client.CreateMarketBuyOrder(&novadax.MarketBuyOrder{
		Value:  fmt.Sprintf("%.2f", 15.00), // target currency amount to buy
		Type:   "MARKET",
		Side:   "BUY",
		Symbol: "TRX_USDT",
	})
	if err != nil {
		log.Panicf("unable to buy TRX with USDT: %s", err.Error())
	}

	log.Printf("bought %s %s", buyOrder.Value, "TRX")
}
