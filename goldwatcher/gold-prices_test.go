package main

import (
	"testing"
)

func TestGold_GetPrices(t *testing.T) {
	g := Gold{
		Prices: nil,
		Client: client,
	}

	price, err := g.GetPrices()
	if err != nil {
		t.Error(err)
	}

	if price.Price != 1866.125 {
		t.Error("wrong price returned:", price.Price)
	}
}
