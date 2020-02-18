package main

import (
	"fmt"
	"log"
	"time"

	novadax "github.com/artemis-tech/novadax-go"
)

func main() {
	client := novadax.Default()

	buyOrder, err := client.CreateOrder(&novadax.Order{
		Amount: fmt.Sprintf("%f", 0.001), // target currency amount to buy
		Type:   "LIMIT",
		Side:   "BUY",
		Price:  fmt.Sprintf("%f", 9500.00),
		Symbol: "BTC_USDT",
	})
	if err != nil {
		log.Panicf("unable to buy in: %s", err.Error())
	}

	log.Printf("created order to buy %s %s (%s)", buyOrder.Amount, "BTC", buyOrder.Price)

	success, err := client.CancelOrder(buyOrder.ID)
	if err != nil {
		log.Panicf("unable to buy in: %s", err.Error())
	}

	for !success {
		result, err := client.CancelOrder(buyOrder.ID)
		if err == nil {
			success = result
		}
		time.Sleep(3)
	}
	log.Printf("canceled order %s", buyOrder.ID)
}
