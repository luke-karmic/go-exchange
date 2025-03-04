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
	// Create an orderbook
	orderBook := NewOrderBook()
	participants := []*MarketUser{}
	for i := 0; i < MARKET_PARTICIPANTS; i++ {
		participants = append(participants, NewMarketUser(i+1, PARTICIPANT_INIT_SUPPLY))
	}

	go SimulateTrades(orderBook, participants)
	time.Sleep(time.Second * TRADE_PER_SECOND_RANGE)
}

func SimulateTrades(ob *Orderbook, mp []*MarketUser) {
	for range time.Tick(time.Second * 1) {
		fmt.Println("Foo")
		fmt.Print(mp)
		fmt.Print(ob)

		// For every participant, create a bid or ask based on available order inside the order book
	}
}
