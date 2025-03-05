package main

import (
	"fmt"
	"time"
)

type Match struct {
	Ask        *Order
	Bid        *Order
	SizeFilled float64
	Price      float64
}

type MarketUser struct {
	Name    string
	Balance int
}

type Limit struct {
	Price       float64
	Orders      []*Order
	TotalVolume float64
}

type Order struct {
	Size      float64
	Bid       bool
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

type Orderbook struct {
	Asks []*Limit
	Bids []*Limit

	AskLimits map[float64]*Limit
	BidLimits map[float64]*Limit
}

func NewOrderBook() *Orderbook {
	fmt.Println("TEST")
	return &Orderbook{
		Asks:      []*Limit{},
		Bids:      []*Limit{},
		AskLimits: make(map[float64]*Limit),
		BidLimits: make(map[float64]*Limit),
	}
}

func NewMarketUser(number int, balance int) *MarketUser {
	return &MarketUser{
		Name:    fmt.Sprintf("User_Name_%v", number),
		Balance: balance,
	}
}

// 10 BTC @15k
func (ob *Orderbook) PlaceOrder(price float64, o *Order) []Match {
	// 1. Try to match the order

	// Decrease remaining size of the order after matching

	// 2. Add the remaining size of the order to the books
	if o.Size > 0.0 {
		ob.Add(price, o)
	}

	return []Match{}
}

func (ob *Orderbook) Add(price float64, o *Order) {
	// PreRequisites before we can successfully add an order
	var limit *Limit

	// 1. There must be enough volume available
	if o.Bid {
		limit = ob.BidLimits[price]
	} else {
		limit = ob.AskLimits[price]
	}

	if limit == nil {
		limit = NewLimitOrder(price)
		if o.Bid {
			ob.Bids = append(ob.Bids, limit)
			ob.BidLimits[price] = limit
		} else {
			ob.Asks = append(ob.Asks, limit)
			ob.AskLimits[price] = limit
		}
	}
	// Spread & Price impact should be acceptable
}
