package models

type Rate struct {
	Currency string
	Price    float64 //json encoder takes 64 bit float
}
