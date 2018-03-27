package parser

import (
	"errors"
	"io"
	"strings"

	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
)

// ProcessPatterns For a set of declarations and an input html stream, returns the mapped data
func ProcessPatterns(reader io.Reader, declarations map[string]string) map[string](string) {
	results := make(map[string]string)
	// Parse HTML document
	doc, err := htmlquery.Parse(reader)
	if err != nil {
		panic(err)
	}
	for k, v := range declarations {
		s, err := evaluateString(doc, v)
		if err == nil {
			results[k] = s
		}
	}
	return results
}

func evaluateString(parentNode *html.Node, pattern string) (string, error) {
	n := htmlquery.FindOne(parentNode, pattern)
	if n != nil {
		wantAttribute, attribute := lastPatternPartIfAttribute(pattern)

		if wantAttribute {
			return htmlquery.SelectAttr(n, attribute), nil
		}
		// fmt.Printf("From: %s\n", htmlquery.OutputHTML(n, false))
		return htmlquery.OutputHTML(n, false), nil
	}

	return "", errors.New("Pattern not matched")
}

// Return true and the last part of an XPATH pattern, if it is an attribute. Otherwise: false, ""
func lastPatternPartIfAttribute(pattern string) (bool, string) {
	sl := strings.Split(pattern, "/")
	if len(sl) < 2 {
		return false, ""
	}

	lastPart := sl[len(sl)-1]
	if strings.HasPrefix(lastPart, "@") {
		return true, strings.TrimLeft(lastPart, "@")
	}

	return false, ""
}
