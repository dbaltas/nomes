[![Build Status](https://travis-ci.org/dbaltas/nomes.svg?branch=master)](https://travis-ci.org/dbaltas/nomes)
# Nomes
html to machine friendly format using xpath

## TODO
* Convert results to json
* input from kafka
* handle xpath expressions returning lists [parser]
* handle xpath expressions returning unwanted embedded html ex.`121,283 <span class="results-label">results</span>` [parser]

## Run
```
$ cp config.yaml.example config.yaml
$ go run main.go -in examples/so.htm
```
## Test
```
go test
```