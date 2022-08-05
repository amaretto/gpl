package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		doc, err := html.Parse(resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "outline: %v\n", err)
			os.Exit(1)
		}
		n := ElementByID(doc, "h1")
		fmt.Println(n)
	}
}

func ElementByID(doc *html.Node, id string) *html.Node {
	return forEachNode(doc, id, startElement, nil)
}

func forEachNode(n *html.Node, id string, pre, post func(n *html.Node, id string) bool) *html.Node {
	var isStop bool
	var h *html.Node
	if pre != nil {
		isStop = pre(n, id)
	}
	if isStop {
		return n
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		h = forEachNode(c, id, pre, post)
		if h != nil {
			return h
		}
	}

	//	if post != nil {
	//		isStop = post(n, id)
	//	}
	return h
}

var depth int

func startElement(n *html.Node, id string) bool {
	if n.Type == html.ElementNode && n.Data == id {
		return true
	}
	return false
}

//func endElement(n *html.Node, id string) bool {
//	if n.Type == html.ElementNode && n.Data == id {
//		return true
//	}
//	return false
//}
