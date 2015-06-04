package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var ErrInvalidParameters = errors.New("stock: Not enough parameters to create object")
var ErrInvalidParameterType = errors.New("stock: Unable to parse parameter")

type Stock struct {
	Acct        string
	Symbol      string
	Quantity    float64
	AvgCost     float64
	MarketPrice float64
	BookValue   float64
	MarketValue float64
}

func NewStock(acctNumber string, raw []string) (Stock, error) {
	if len(raw) < 6 {
		return Stock{}, ErrInvalidParameters
	}

	var stock = Stock{
		Acct:   acctNumber,
		Symbol: raw[0],
	}

	var err error
	stock.Quantity, err = strconv.ParseFloat(strings.Replace(raw[1], ",", "", -1), 64)
	if err != nil {
		return stock, ErrInvalidParameterType
	}
	stock.AvgCost, err = strconv.ParseFloat(strings.Replace(raw[2], ",", "", -1), 64)
	if err != nil {
		return stock, ErrInvalidParameterType
	}
	stock.MarketPrice, err = strconv.ParseFloat(strings.Replace(raw[3], ",", "", -1), 64)
	if err != nil {
		return stock, ErrInvalidParameterType
	}
	stock.BookValue, err = strconv.ParseFloat(strings.Replace(raw[4], ",", "", -1), 64)
	if err != nil {
		return stock, ErrInvalidParameterType
	}
	stock.MarketValue, err = strconv.ParseFloat(strings.Replace(raw[5], ",", "", -1), 64)
	if err != nil {
		return stock, ErrInvalidParameterType
	}

	return stock, nil
}

func (s *Stock) CsvOutput() string {
	return fmt.Sprintf("%s,%s,%f,%f,%f,%f,%f", s.Acct, s.Symbol, s.Quantity, s.AvgCost, s.MarketPrice, s.BookValue, s.MarketValue)
}
