package main

import (
	"bytes"
	"io"
	"net/http"
	"testing"
)

func TestGold_GetPrices(t *testing.T) {
	client := NewTestClient(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBufferString(jsonToReturn)),
			Header:     make(http.Header),
		}
	})

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
