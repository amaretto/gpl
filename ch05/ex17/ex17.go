package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

var depth int

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}
	elms := ElementByTagName(doc, "script", "h2")
	for _, elm := range elms {
		fmt.Println(elm)
	}
}

var first bool = true

var nodeArry []*html.Node

func ElementByTagName(doc *html.Node, name ...string) []*html.Node {

	for c := doc.FirstChild; c != nil; c = c.NextSibling {
		ElementByTagName(c, name...)
	}

	for _, elm := range name {
		if doc.Type == html.ElementNode && doc.Data == elm {
			nodeArry = append(nodeArry, doc)
		}
	}

	return nodeArry
}
