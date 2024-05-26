package main

import (
	"fmt"

	"github.com/lakshhtaneja/cryptoprice/api"
)

func main() {
	rate, err := api.GetRate("ETH")
	if err == nil {
		fmt.Printf("Rate for %s is %.2f USD\n", rate.Currency, rate.Price)
	} else {
		fmt.Println(err)
	}
}
