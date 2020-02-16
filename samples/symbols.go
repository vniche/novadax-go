package main

import (
	"log"

	novadax "github.com/artemis-tech/novadax-go"
)

func main() {
	client := novadax.Default()

	symbols, err := client.ListSymbols()
	if err != nil {
		log.Printf("%s", err.Error())
	}

	log.Printf("%+v", symbols)
	for _, symbol := range symbols {
		log.Printf("%+v", symbol)
	}
}
