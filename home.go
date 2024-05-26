package main

import (
	"fmt"

	"github.com/lakshhtaneja/cryptoprice/api"
)

func main() {
	var code string

	fmt.Println("Welcome to Crypto Price")
	fmt.Println("=========================")
	fmt.Println("Enter Currency code")
	fmt.Scanf("%s", &code)

	rate, err := api.GetRate(code)
	if err == nil {
		fmt.Printf("Rate for %s is %.2f USD\n", rate.Currency, rate.Price)
	} else {
		fmt.Println(err)
	}
}
