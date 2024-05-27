package api

import "fmt"

type Response struct {
	Timestamp             string  `json:"timestamp"`
	Low                   string  `json:"low"`
	High                  string  `json:"high"`
	Last                  string  `json:"last"`
	Volume                string  `json:"volume"`
	Volume30D             string  `json:"volume30d"`
	Bid                   float64 `json:"bid"`
	Ask                   float64 `json:"ask"`
	PriceChange           string  `json:"priceChange"`
	PriceChangePercentage string  `json:"priceChangePercentage"`
	Pair                  string  `json:"pair"`
}

func PrintRate(code string) {
	rate, err := GetRate(code)
	if err == nil {
		fmt.Printf("Rate for %s is %.2f USD\n", rate.Currency, rate.Price)
	} else {
		fmt.Println(err)
	}
}
