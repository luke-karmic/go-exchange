package main

import (
	"fmt"
	"time"
)

type Orderbook struct {
	Asks []*Limit
	Bids []*Limit

	AskLimits map(float64) *Limit
	BidLimits map(float64) *Limit
}

type MarketUser struct {
  Name string
  Balance int
}

type Limit struct {
	Price       float64
	Orders      []*Order
	TotalVolume float64
}

type Order struct {
	Size      float64
	Limit     *Limit
	Timestamp int64
}

func NewOrder(bid bool, size float64) *Order {
	return &Order{
		Size:      size,
		Timestamp: time.Now().UnixNano(),
	}
}

func NewLimitOrder(price float64) *Limit {
	return &Limit{
		Price:  price,
		Orders: []*Order{},
	}
}

func (o *Order) String() string {
	return fmt.Sprintf("[size %.2f]", o.Size)
}

func (l *Limit) AddOrder(o *Order) {
	o.Limit = l
	l.Orders = append(l.Orders, o)
	l.TotalVolume += o.Size
}

func (l *Limit) DeleteOrder(o *Order) {
	for i := 0; i < len(l.Orders); i++ {
		if l.Orders[i] == o {
			l.Orders[i] = l.Orders[len(l.Orders)-1]
			l.Orders = l.Orders[:len(l.Orders)-1]
		}
	}
	o.Limit = nil
	l.TotalVolume -= o.Size

	// TODO: Re-sort all resting orders
}

func NewOrderBook() *OrderBook {
	return &OrderBook{
		Asks: []*Limit{},
		Bids: []*Limit{},
		AskLimits: make(map[float64]*Limit),
		BidLimits: make(map[float64]*Limit),
	}
}

func NewMarketUser(number int, balance int) *MarketUser {
	return &MarketUser {
    Name: "User_Name_" + iterator,
    Balance: balance,
  }
}
