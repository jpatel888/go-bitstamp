package main

import (
	"fmt"
	"bitstamp"
)

func main() {
	bitstamp_client := bitstamp.New()

	ticker, err := bitstamp_client.GetTicker("btcusd")

	fmt.Println(err, ticker)
	fmt.Println("Last:", ticker.Last)
}