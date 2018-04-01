package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/dbaltas/nomes/parser"
	"github.com/spf13/viper"
)

func main() {
	var inFile string

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

	results := parser.ProcessPatterns(f, configMap())
	for k, v := range results {
		fmt.Printf("%s: %s\n", k, v)
	}
}

func configMap() map[string]string {
	viper.AddConfigPath(".")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	return viper.GetStringMapString("mappings")
}
