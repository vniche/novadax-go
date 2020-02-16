package main

import (
	"log"

	novadax "github.com/artemis-tech/novadax-go"
)

func main() {
	/*
	* This configuration is possible via environment variables, eg.:
	* NOVADAX_ACCESS_KEY="5388359-538583-5i9593-3596e0-6ca252484934aa4"
	* NOVADAX_SECRET_KEY="nl3KVXiOp4JN74482h4nkahiu5jDKWkKhnMumMy"
	*
	* or programmatically via:
	* novadax.New("ACCESS_KEY", "PRIVATE_KEY")
	 */
	client := novadax.Default()

	orders, err := client.ListOrders(&novadax.ListOrdersFilters{
		Symbol: "BTC_BRL",
		Limit:  10,
	})
	if err != nil {
		log.Printf("%s", err.Error())
		return
	}

	log.Printf("%+v", orders)
	for _, order := range orders {
		log.Printf("%+v", order)
	}
}
