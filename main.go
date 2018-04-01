package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

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

	// hack to trim spaces
	// due to inability to use normalize-space()
	for k, v := range results {
		results[k] = strings.TrimSpace(v)
	}

	jsonString, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		fmt.Printf("error json marshaling: %s\n", err)
		return
	}
	fmt.Println(string(jsonString))
}

func configMap() map[string]string {
	viper.AddConfigPath(".")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	return viper.GetStringMapString("mappings")
}
