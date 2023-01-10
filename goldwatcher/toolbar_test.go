package main

import (
	"testing"

	"fyne.io/fyne/v2/test"
)

func TestApp_getToolBar(t *testing.T) {
	tb := testApp.getToolBar()

	if len(tb.Items) != 4 {
		t.Error("wrong number of items in toolbar")
	}
}

func TestApp_addHoldingsDialog(t *testing.T) {
	testApp.addHoldingsDialog()

	test.Type(testApp.AddHoldingsPurchaseAmountEntry, "1")
	test.Type(testApp.AddHoldingsPurchasePriceEntry, "1250")
	test.Type(testApp.AddHoldingsPurchaseDateEntry, "2020-03-20")

	if testApp.AddHoldingsPurchaseDateEntry.Text != "2020-03-20" {
		t.Error("date not correct")
	}

	if testApp.AddHoldingsPurchaseAmountEntry.Text != "1" {
		t.Error("purchase amount not correct")
	}

	if testApp.AddHoldingsPurchasePriceEntry.Text != "1250" {
		t.Error("price not correct")
	}
}
