package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	w := os.Stdout
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
		forEachNode(doc, startElement, endElement, w)
	}
}

func forEachNode(n *html.Node, pre func(n *html.Node, w io.Writer) bool, post func(n *html.Node, w io.Writer), w io.Writer) {
	isSkip := false
	if pre != nil {
		isSkip = pre(n, w)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post, w)
	}

	if post != nil && !isSkip {
		post(n, w)
	}
}

var depth int

func startElement(n *html.Node, w io.Writer) bool {
	var suffix string
	isSkip := false
	if n.FirstChild == nil {
		isSkip = true
		suffix = "/"
	}
	if n.Type == html.ElementNode {
		if n.Attr != nil {
			var attrs string
			for _, attr := range n.Attr {
				attrs = attrs + fmt.Sprintf(" %s='%s'", attr.Key, attr.Val)
			}
			fmt.Fprintf(w, "%*s<%s%s%s>\n", depth*2, "", n.Data, attrs, suffix)
		} else {
			fmt.Fprintf(w, "%*s<%s%s>\n", depth*2, "", n.Data, suffix)
		}

		if !isSkip {
			depth++
		}
	} else {
		s := strings.Split(n.Data, "\n")
		for _, row := range s {
			if row == "" {
				continue
			}
			fmt.Fprintf(w, "%*s%s\n", depth*2, "", row)
		}
	}

	return isSkip
}

func endElement(n *html.Node, w io.Writer) {
	if n.Type == html.ElementNode {
		depth--
		fmt.Fprintf(w, "%*s</%s>\n", depth*2, "", n.Data)
	}
}
