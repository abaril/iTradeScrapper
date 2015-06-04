package main

import "testing"

func TestNewStock(t *testing.T) {

	var stock Stock
	var err error

	_, err = NewStock("abc123", []string{"symbol", "100"})
	if err == nil {
		t.Error("Should fail due to inadequate parameters")
	}

	_, err = NewStock("abc123", []string{"symbol", "abc", "12,000.34", "3.14", "14.43", "54.232"})
	if err == nil {
		t.Error("Should fail due to invalid parameter type")
	}

	_, err = NewStock("abc123", []string{"symbol", "100", "12,000.34", "abc", "14.43", "54.232"})
	if err == nil {
		t.Error("Should fail due to invalid parameter type")
	}

	stock, err = NewStock("abc123", []string{"symbol", "100", "12,000.34", "3.14", "14.43", "54.232"})
	if err != nil {
		t.Error("Should parse fine")
	} else if stock.Symbol != "symbol" || stock.Quantity != 100 || stock.AvgCost != 12000.34 || stock.MarketPrice != 3.14 || stock.BookValue != 14.43 || stock.MarketValue != 54.232 {
		t.Error("Invalid value parsed")
	}
}

func TestCsvOutput(t *testing.T) {
	stock, err := NewStock("abc123", []string{"symbol", "100", "12,000.34", "3.14", "14.43", "54.232"})
	if err != nil {
		t.Error("Shouldn't have a problem here")
	}

	expected := "abc123,symbol,100.000000,12000.340000,3.140000,14.430000,54.232000"
	csv := stock.CsvOutput()
	if csv != expected {
		t.Errorf("Doesn't match.\n\tExpected = %s\n\tActual = %s", expected, csv)
	}
}
