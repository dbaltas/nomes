package parser

import (
	"log"
	"os"
	"strings"
	"testing"

	"github.com/antchfx/htmlquery"
)

const htmlSample = `<!DOCTYPE html><html lang="en-US">
<head>
<title>Hello,World!</title>
</head>
<body>
<div contain="container">
<header>
	 <!-- Logo -->
   <h1>City Gallery</h1>
</header>  
<nav>
  <ul>
    <li><a href="#">London</a></li>
    <li><a href="#">Paris</a></li>
    <li><a href="#">Tokyo</a></li>
  </ul>
</nav>
<article>
  <h1>London</h1>
	<ul>
		<li><a href="#">London</a></li>
		<li><a href="#">Paris</a></li>
		<li><a href="#">Tokyo</a></li>
	</ul>
	<img src="pic_mountain.jpg" alt="Mountain View" style="width:304px;height:228px;">
  <p>London is the capital city of England. It is the most populous city in the  United Kingdom, with a metropolitan area of over 13 million inhabitants.</p>
  <p>Standing on the River Thames, London has been a major settlement for two millennia, its history going back to its founding by the Romans, who named it Londinium.</p>
</article>
<footer>Copyright &copy; W3Schools.com</footer>
</div>
</body>
</html>
`

func TestEvaluateString(t *testing.T) {
	doc, err := htmlquery.Parse(strings.NewReader(htmlSample))
	if err != nil {
		t.Fatal(err)
	}
	pattern := "//head/title"
	expected := "Hello,World!"

	s, err := evaluateString(doc, pattern)
	if err != nil {
		t.Fatalf("Error processing pattern %s:, %s", pattern, err)
	}

	if s != expected {
		t.Fatalf("'%s' does not match expected '%s'", s, expected)
	}

	pattern = "//article/img[@src='pic_mountain.jpg']/@alt"
	expected = "Mountain View"

	s, err = evaluateString(doc, pattern)
	if err != nil {
		t.Fatalf("Error processing pattern %s:, %s", pattern, err)
	}

	if s != expected {
		t.Fatalf("'%s' does not match expected '%s'", s, expected)
	}
}

func TestStackOverflow(t *testing.T) {
	inFile := "./../examples/so.htm"
	f, err := os.Open(inFile)
	if err != nil {
		log.Printf("failed to open file %s", inFile)
		panic(err)
	}

	var configMap = make(map[string]string)
	var expected = make(map[string]string)
	configMap["main.url"] = "//head/meta[@property='og:url']/@content"
	configMap["main.results"] = "//div[@id='mainbar']//h2"
	configMap["main.maxpage"] = "//div[@id='mainbar']//span[@class='page-numbers dots']/following::*/span"
	//TODO: calculate 2nd element by position, and not data-position
	configMap["main.second.votes"] = "//div[@class='question-summary search-result']//div[@data-position='2']//div[@class='votes answered']//strong"
	configMap["main.second.url"] = "//div[@class='question-summary search-result']//a/@href"

	expected["main.url"] = "https://stackoverflow.com/search?tab=votes&q=xpath"
	expected["main.results"] = "121,283 <span class=\"results-label\">results</span>"
	expected["main.maxpage"] = "8086"
	expected["main.second.votes"] = "642"
	expected["main.second.url"] = "/questions/3577641/how-do-you-parse-and-process-html-xml-in-php/3577662#3577662"

	results := ProcessPatterns(f, configMap)

	if len(results) != 5 {
		t.Fatalf("Expected results %d. Found %d", len(results), 5)
	}

	for k, v := range expected {
		if results[k] != v {
			t.Fatalf("Expected:\n%s\nFound\n%s\nFor: %s %s", v, results[k], k, configMap[k])
		}
	}
}
