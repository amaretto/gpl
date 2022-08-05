package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
	}
	m := make(map[string]int)
	for k, v := range countElements(m, doc) {
		fmt.Printf("%s : %d\n", k, v)
	}
}

// memo: no need to return map object
func countElements(em map[string]int, n *html.Node) map[string]int {
	if n.Type == html.ElementNode {
		em[n.Data]++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		em = countElements(em, c)
	}
	return em
}
