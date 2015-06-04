package main

import (
	"errors"
	"golang.org/x/net/html"
	"io"
	"log"
	"strings"
)

type Scrapper interface {
	Scrap(in io.Reader)
}

func NewScrapper(out chan *Stock) Scrapper {
	return &itradeScrapper{
		out:           out,
		accountNumber: "UNKNOWN",
	}
}

const classWithStockTable = "var-acct-det-hst-expand-cat"
const classWithAcctNumber = "branding-header"
const expectedPrefixForAcctNumber = "Scotia iTRADE - "

type itradeScrapper struct {
	out           chan *Stock
	accountNumber string
}

func (i *itradeScrapper) Scrap(in io.Reader) {
	defer close(i.out)

	doc, err := html.Parse(in)
	if err != nil {
		log.Fatal("Unable to parse HTML", err)
		return
	}

	i.scanBody(doc)
}

func (i *itradeScrapper) scanBody(n *html.Node) {

	if n.Type == html.ElementNode {
		if n.Data == "div" {
			for _, attr := range n.Attr {
				if attr.Key == "class" && attr.Val == classWithAcctNumber {
					for c := n.FirstChild; c != nil; c = c.NextSibling {
						if accountNumber, err := i.scanForAccountNumber(c); err == nil {
							i.accountNumber = accountNumber
						}
					}
				}
			}
		} else if n.Data == "tbody" {
			for _, attr := range n.Attr {
				if attr.Key == "class" && attr.Val == classWithStockTable {
					for c := n.FirstChild; c != nil; c = c.NextSibling {
						row := i.scanTableRow(c)
						if len(row) > 0 {
							stock, err := NewStock(i.accountNumber, row)
							if err == nil {
								i.out <- &stock
							}
						}
					}
				}
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		i.scanBody(c)
	}
}

func (_ *itradeScrapper) scanForAccountNumber(n *html.Node) (string, error) {

	if n.Type == html.ElementNode && n.Data == "h3" {
		for _, attr := range n.Attr {
			if attr.Key == "title" {
				data := strings.TrimSpace(attr.Val)
				if strings.HasPrefix(data, expectedPrefixForAcctNumber) {
					acctNumber := strings.TrimPrefix(data, expectedPrefixForAcctNumber)
					return strings.Split(acctNumber, " ")[0], nil
				}
			}
		}
	}
	return "", errors.New("No account number found")
}

func (i *itradeScrapper) scanTableRow(n *html.Node) []string {
	var values []string

	if n.Type == html.ElementNode && n.Data == "tr" {
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			if c.Type == html.ElementNode && c.Data == "td" {
				if c.FirstChild != nil && c.FirstChild.Data == "div" {
					value, err := i.parseRawValue(c.FirstChild)
					if err == nil {
						values = append(values, value)
					}
				}
			}
		}
	}

	return values
}

func (i *itradeScrapper) parseRawValue(n *html.Node) (string, error) {

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.TextNode {
			trimmed := strings.TrimSpace(c.Data)
			if len(trimmed) > 0 {
				return trimmed, nil
			}
		} else if c.Type == html.ElementNode {
			return i.parseRawValue(c)
		}
	}

	return "", errors.New("No value available")
}
