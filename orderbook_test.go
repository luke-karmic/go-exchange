package main

import (
	"fmt"
	"testing"
)

func TestOrderbook(t *testing.T) {

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
