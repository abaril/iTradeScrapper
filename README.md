# iTradeScrapper
A simple go command-line app to scrap the stock values from a dump of an iTrade account's HTML source.

This simple command-line utility is likely of little use to anyone else, but was more developed as an intro to Go.

## Build
You will need to have Go installed (https://golang.org/doc/install)
```
go build -o scrapper
```

## Running
iTradeScrapper requires the HTML page source from iTrade to be available in an input file. It will scrap from the input file provided and output to stdout.
```
./scrapper -i [input.html]
```

