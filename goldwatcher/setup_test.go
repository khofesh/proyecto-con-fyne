package main

import (
	"net/http"
	"os"
	"testing"
)

var testApp Config

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

var jsonToReturn = `
{
	"ts": 1673177564137,
	"tsj": 1673177555642,
	"date": "Jan 8th 2023, 06:32:35 am NY",
	"items": [
	  {
		"curr": "USD",
		"xauPrice": 1866.125,
		"xagPrice": 23.841,
		"chgXau": 33.29,
		"chgXag": 0.618,
		"pcXau": 1.8163,
		"pcXag": 2.6612,
		"xauClose": 1832.835,
		"xagClose": 23.223
	  }
	]
  }
`

type RoundTripFunc func(req *http.Request) *http.Response

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: fn,
	}
}
