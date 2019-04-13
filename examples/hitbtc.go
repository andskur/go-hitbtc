package main

import (
	"fmt"
	"time"

	"github.com/davecgh/go-spew/spew"

	"github.com/andskur/go-hitbtc"
)

const (
	API_KEY    = "HzLLapvLNF0KR0NZ4lHikUSlcXCz6O8n"
	API_SECRET = "rkWSjiyP94+OzdCQKEgchYTUHOHhTVHq"
)

func main() {
	// hitbtc client
	rest := hitbtc.New(API_KEY, API_SECRET)

	// GetBalances
	balances, _ := rest.GetBalances()
	fmt.Println(len(balances))

	for _, bal := range balances {
		if bal.Currency == "BTC" {
			fmt.Println(bal.Available)
		}
	}

	// Initialize websocket connection
	client, err := hitbtc.NewWSClient()
	if err != nil {
		panic(err) // do something
	}
	defer client.Close()

	symbol, err := client.GetSymbol("ETHBTC")
	if err != nil {
		panic(err)
	}
	spew.Dump(symbol)

	ok, err := client.AuthSession("BASIC", API_KEY, API_SECRET)
	if err != nil {
		panic(err)
	}

	fmt.Println(ok)

	_, reports, err := client.SubscribeTradeReports()
	if err != nil {
		panic(err)
	}

	for {
		ord := <-reports

		spew.Dump(ord)

		time.Sleep(5 * time.Second)
	}

	/*_, _, err = client.SubscribeOrderbook("ETHBTC")

	for {
		trBalances, err := client.GetTradingBalances()
		if err != nil {
			panic(err)
		}

		for _, bal := range trBalances.Balances {
			if bal.Currency == "XVG" {
				fmt.Println(bal.Available)
			}
		}

		time.Sleep(30 * time.Millisecond)
	}*/

	// Subscribe and handle
	/*tickerFeed, _, err := client.SubscribeOrderbook("ETHBTC")
	for {
		ticker := <-tickerFeed
		fmt.Println(ticker)
	}*/

}
