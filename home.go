package main

import (
	"fmt"
	"sync"

	"github.com/lakshhtaneja/cryptoprice/api"
)

func main() {
	defaultCurrencies := []string{"BTC", "ETH"}
	var wg sync.WaitGroup
	fmt.Println("Welcome to Crypto Price")
	fmt.Println("=========================")

	for _, currency := range defaultCurrencies {
		wg.Add(1)
		go func(currency string) {
			api.PrintRate(currency)
			wg.Done()
		}(currency)
	}
	wg.Wait()
}
