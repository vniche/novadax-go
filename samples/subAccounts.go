package main

import (
	"log"

	novadax "github.com/vniche/novadax-go"
)

func main() {
	client := novadax.Default()

	subAccounts, err := client.AccountSubs()
	if err != nil {
		log.Panicf("unable to list sub accounts: %s", err.Error())
	}

	for _, subAccount := range subAccounts {
		log.Printf("%+v", subAccount)
	}
}
