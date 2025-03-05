package main

import (
	"fmt"
	"testing"
)

func TestOrderbook(t *testing.T) {
	ob := NewOrderBook()
	buyOrder := NewOrder(true, 10)
	ob.PlaceOrder(18_000, buyOrder)

	fmt.Println(ob.Bids)
}

func TestLimit(t *testing.T) {
	l := NewLimitOrder(10_000)
	buyOrderA := NewOrder(true, 5)
	buyOrderB := NewOrder(true, 8)
	buyOrderC := NewOrder(true, 12)
	l.AddOrder(buyOrderA)
	l.AddOrder(buyOrderB)
	l.AddOrder(buyOrderC)
	l.DeleteOrder(buyOrderB)

	fmt.Println(l)
}
