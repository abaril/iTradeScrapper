package main

import (
	"strings"
	"testing"
)

const (
	testInput = "<div class=\"branding-header\"><h3 title=\"Scotia iTRADE - abc123 Cash\"/></div><table><tbody class=\"var-acct-det-hst-expand-cat\"><tr><td><div>symbol</div></td><td><div>100</div></td><td><div>1.23</div></td><td><div>3.14</div></td><td><div>10000</div></td><td><div>20000</div></td><td><div>30000</div></td><td><div>40000</div></td></tr></tbody></table>"
)

func TestNew(t *testing.T) {
	out := make(chan *Stock)
	scrapper := NewScrapper(out)
	if scrapper == nil {
		t.Error("Unable to create scrapper")
	} else {
		in := strings.NewReader(testInput)
		go scrapper.Scrap(in)

		result := <-out

		expected := Stock{
			Acct:        "abc123",
			Symbol:      "symbol",
			Quantity:    100,
			AvgCost:     1.23,
			MarketPrice: 3.14,
			BookValue:   10000,
			MarketValue: 20000,
		}
		if *result != expected {
			t.Errorf("Output incorrect. Expected %v, Received %v", expected, *result)
		}
	}
}
