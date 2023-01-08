package main

import (
	"bytes"
	"goldwatcher/repository"
	"io"
	"log"
	"net/http"
	"os"
	"testing"

	"fyne.io/fyne/v2/test"
)

var testApp Config

func TestMain(m *testing.M) {
	a := test.NewApp()
	testApp.App = a
	testApp.MainWindow = a.NewWindow("")
	log.Println("setting client to test client")
	testApp.HTTPClient = client
	testApp.DB = repository.NewTestRepository()
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

var client = NewTestClient(func(req *http.Request) *http.Response {
	return &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewBufferString(jsonToReturn)),
		Header:     make(http.Header),
	}
})
