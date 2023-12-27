package display

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type Response struct {
	BtcUsd BtcUsd `json:"BTC_USD"`
	LtcUsd LtcUsd `json:"LTC_USD"`
	EthUsd EthUsd `json:"ETH_USD"`
}

type BtcUsd struct {
	BuyPrice  string `json:"buy_price"`
	SellPrice string `json:"sell_price"`
	LastTrade string `json:"last_trade"`
	High      string `json:"high"`
	Low       string `json:"low"`
	Avg       string `json:"avg"`
	Vol       string `json:"vol"`
	VolCurr   string `json:"vol_curr"`
	Updated   int    `json:"updated"`
}

type LtcUsd struct {
	BuyPrice  string `json:"buy_price"`
	SellPrice string `json:"sell_price"`
	LastTrade string `json:"last_trade"`
	High      string `json:"high"`
	Low       string `json:"low"`
	Avg       string `json:"avg"`
	Vol       string `json:"vol"`
	VolCurr   string `json:"vol_curr"`
	Updated   int    `json:"updated"`
}

type EthUsd struct {
	BuyPrice  string `json:"buy_price"`
	SellPrice string `json:"sell_price"`
	LastTrade string `json:"last_trade"`
	High      string `json:"high"`
	Low       string `json:"low"`
	Avg       string `json:"avg"`
	Vol       string `json:"vol"`
	VolCurr   string `json:"vol_curr"`
	Updated   int    `json:"updated"`
}

func GetPriceData(ch chan<- Response) {
	for {
		url := "https://api.exmo.com/v1.1/ticker"
		method := "POST"

		payload := strings.NewReader("")

		client := &http.Client{}
		req, err := http.NewRequest(method, url, payload)

		if err != nil {
			fmt.Println(err)
			return
		}
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

		res, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer res.Body.Close()

		var response Response
		err = json.NewDecoder(res.Body).Decode(&response)
		//price, _ := strconv.ParseFloat(response.BtcUsd.SellPrice, 64)
		ch <- response
		time.Sleep(time.Second)
	}

}
