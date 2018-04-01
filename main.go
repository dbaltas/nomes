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
	var configFile string

	flag.StringVar(&inFile, "in", "", "the input file in html format")
	flag.StringVar(&configFile, "config", "config.*", "filename for configuration")
	flag.Parse()

	if inFile == "" {
		log.Println("no in file defined")
		flag.Usage()
		return
	}

	f, err := os.Open(inFile)
	if err != nil {
		log.Printf("failed to open file %s\n", inFile)
		panic(err)
	}

	results := parser.ProcessPatterns(f, configMap(configFile))

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

func configMap(configFile string) map[string]string {
	viper.AddConfigPath(".")
	if configFile != "" {
		viper.SetConfigFile(configFile)
	}
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	return viper.GetStringMapString("mappings")
}
