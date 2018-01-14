package bitstamp

import "encoding/json"

const (
	API_BASE                   = "https://www.bitstamp.net/api/"
	API_VERSION                = "v2"
	DEFAULT_HTTPCLIENT_TIMEOUT = 8
)

func New() *Bitstamp {
	client := NewClient()
	return &Bitstamp{client}
}

type Bitstamp struct {
	client *client
}

type Ticker struct {
	High  float64 `json:"high,string"`
	Last  float64 `json:"last,string"`
	TimeStamp uint `json:"timestamp,string"`
	Bid float64 `json:"bid,string"`
	Vwap float64 `json:"vwap,string"`
	Volume float64 `json:"volume,string"`
	Low float64 `json:"low,string"`
	Ask float64 `json:"ask,string"`
	Open float64 `json:"open,string"`
}

func (b *Bitstamp) GetTicker(market string) (ticker Ticker, err error) {
	r, err := b.client.do("GET", "ticker/"+market)
	if err != nil {
		return
	}
	if err = json.Unmarshal(r, &ticker); err != nil {
		return
	}
	return
}