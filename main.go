package main

import (
	"fmt"
	"log"
	"os"
	"flag"
)

func main() {

	inputFilename := flag.String("i", "", "HTML input filename")
	flag.Parse()

	if (len(*inputFilename) <= 0) {
		log.Fatalln("An input file must be specified. Use -h to view command usage.")
		return
	}

	file, err := os.Open(*inputFilename)
	if err != nil {
		log.Fatalln("Unable to open file: ", err)
		return
	}
	defer file.Close()

	out := make(chan *Stock)
	scrapper := NewScrapper(out)

	go scrapper.Scrap(file)
	for stock := range out {
		fmt.Println(stock.CsvOutput())
	}
}
