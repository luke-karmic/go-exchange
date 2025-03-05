package main

import (
	"fmt"
	"time"
)

const MARKET_PARTICIPANTS = 10
const INITIAL_ASSET_SUPPLY = 100000
const PARTICIPANT_INIT_SUPPLY = 500
const TRADE_PER_SECOND_RANGE = 5

func main() {
	// Create Asset with (X) supply (X) USD value
	orderBook := NewOrderBook()
	participants := []*MarketUser{}
	for i := 0; i < MARKET_PARTICIPANTS; i++ {
		participants = append(participants, NewMarketUser(i+1, PARTICIPANT_INIT_SUPPLY))
	}

	go SimulateTrades(orderBook, participants)
	time.Sleep(time.Second * TRADE_PER_SECOND_RANGE)
}

func SimulateTrades(ob *Orderbook, mp []*MarketUser) {
	for range time.Tick(time.Second * TRADE_PER_SECOND_RANGE) {
		fmt.Println("Foo")
		fmt.Print(mp)
		fmt.Print(ob)

		// For every participant, create realistic bid or ask based on available orders in the order book

		// Print the order from Bid/Ask side, size, number of orders required, the USD value, price impact and spread
	}
}
