package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/dbaltas/nomes/parser"
)

func main() {
	var inFile string
	// Load HTML file.
	flag.StringVar(&inFile, "in", "", "the input file in html format")
	flag.Parse()
	if inFile == "" {
		log.Print("no in file defined")
		flag.Usage()
		os.Exit(1)
	}

	f, err := os.Open(inFile)
	if err != nil {
		log.Printf("failed to open file %s", inFile)
		panic(err)
	}

	fmt.Printf("\n\n")

	var configMap = make(map[string]string)

	configMap["main.url"] = "//head/meta[@property='og:url']/@content"
	configMap["main.results"] = "//div[@id='mainbar']//h2"
	configMap["main.maxpage"] = "//div[@id='mainbar']//span[@class='page-numbers dots']/following::*/span"

	results := parser.ProcessPatterns(f, configMap)
	for k, v := range results {
		fmt.Printf("%s: %s\n", k, v)
	}
}
