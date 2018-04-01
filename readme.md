[![Build Status](https://travis-ci.org/dbaltas/nomes.svg?branch=master)](https://travis-ci.org/dbaltas/nomes)
# Nomes
html to json using xpath

## TODO
* input from kafka
* handle xpath expressions returning lists [parser]
* handle xpath expressions returning unwanted embedded html ex.`121,283 <span class="results-label">results</span>` [parser]

## Run
```
$ go run main.go -in examples/so.htm -config config.example.yaml
```
returns
```json
{
  "main.maxpage": "8086",
  "main.results": "121,283 \u003cspan class=\"results-label\"\u003eresults\u003c/span\u003e",
  "main.second.url": "/questions/3577641/how-do-you-parse-and-process-html-xml-in-php/3577662#3577662",
  "main.second.votes": "642",
  "main.url": "https://stackoverflow.com/search?tab=votes\u0026q=xpath"
}
```
## Test
```
go test
```