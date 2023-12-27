package display

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/eiannone/keyboard"
	"github.com/gosuri/uilive"
	"github.com/guptarohit/asciigraph"
)

func DisplayChart(pair string, ch <-chan Response) {
	//var wg sync.WaitGroup

	writer := uilive.New()
	writer.Start()

	data := make([]float64, 0)

	var (
		char rune
		key  keyboard.Key
		err  error
	)

	go func() {
		for {
			char, key, err = keyboard.GetKey()
			if err != nil {
				panic(err)
			}
			if key == keyboard.KeyBackspace2 {
				return
			}
		}
	}()

	for response := range ch {
		if char == 'q' {
			os.Exit(0)
		} else if key == keyboard.KeyBackspace2 {
			writer.Stop()
			return
		}
		var avg float64
		var sellPrice float64

		switch pair {
		case "BTC_USD":
			avg, _ = strconv.ParseFloat(response.BtcUsd.Avg, 64)
			sellPrice, _ = strconv.ParseFloat(response.BtcUsd.SellPrice, 64)
		case "LTC_USD":
			avg, _ = strconv.ParseFloat(response.LtcUsd.Avg, 64)
			sellPrice, _ = strconv.ParseFloat(response.LtcUsd.SellPrice, 64)
		case "ETH_USD":
			avg, _ = strconv.ParseFloat(response.EthUsd.Avg, 64)
			sellPrice, _ = strconv.ParseFloat(response.EthUsd.SellPrice, 64)
		}

		data = append(data, sellPrice)
		graph := asciigraph.Plot(data, asciigraph.Width(100), asciigraph.Height(10), asciigraph.SeriesColors(
			asciigraph.Red,
		))
		fmt.Fprintf(writer, "%s: %.2f\n", pair, avg)
		fmt.Fprintf(writer, graph+"\n")
		fmt.Fprintf(writer, "Текущее время: %d:%d:%d\n", time.Now().Hour(), time.Now().Minute(), time.Now().Second())
		fmt.Fprintf(writer, "Текущая дата: %d-%d-%d\n", time.Now().Year(), time.Now().Month(), time.Now().Day())

		time.Sleep(time.Second)
	}
}
