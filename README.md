# go-bitstamp

go-bitstamp is a go implementation of the bitstamp API (public) in Golang.

## Import
	import "github.com/jpatel888/go-bitstamp"

## Usage
~~~ go
package main

import (
	"fmt"
	"github.com/jpatel888/go-bitstamp"
)

func main() {
	bitstamp_client := bitstamp.New()

	ticker, err := bitstamp_client.GetTicker("btcusd")

	fmt.Println(err, ticker)
	fmt.Println("Last:", ticker.Last)
}
~~~