package main

import (
	"fmt"
	"os"

	"CLITradeChart/display"

	"github.com/eiannone/keyboard"
)

func main() {
	responses := make(chan display.Response)
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	menu := `
My menu:
	1. BTC_USD
	2. LTC_USD
	3. ETH_USD
Press 1-3 to change symbol, press q to exit
`

	go display.GetPriceData(responses)
	for {
		fmt.Print(menu)
		char, _, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}

		switch char {
		case '1':
			fmt.Print("\033[H\033[2J")
			display.DisplayChart("BTC_USD", responses)
			fmt.Print("\033[H\033[2J")
			//continue
		case '2':
			fmt.Print("\033[H\033[2J")
			display.DisplayChart("LTC_USD", responses)
			fmt.Print("\033[H\033[2J")
			//continue
		case '3':
			fmt.Print("\033[H\033[2J")
			display.DisplayChart("ETH_USD", responses)
			fmt.Print("\033[H\033[2J")
			//continue
		case 'q':
			os.Exit(0)
		}

	}
}
