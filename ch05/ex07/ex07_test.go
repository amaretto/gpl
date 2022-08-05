package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

var testCases = []struct {
	Input  io.Reader
	Output string
}{
	// minimal
	{strings.NewReader(`
<html>
<head>
<title>
head title
</title>
</head>
<body>
body
</body>
</html>
`),
		`
<html>
	<head>
		<title>
			head title
		</title>
	</head>
	<body>
		body
	</body>
</html>
	`,
	},
	//img
	{strings.NewReader(`
<html>
<img border='0' width='150' src='image.png'>
</img>
</html>`),
		`
<html>
	<img border='0' width='150' src='image.png' />
</html>
`,
	},
	//link
	{strings.NewReader(`
<html>
<head>
<title>
title hoge
</title>
</head>
<body>
<a href="https://google.com">
<code>
eg
</code>
</a>
</img>
</body>
</html>`),
		`
<html>
	<head>
		<title>
			title hoge
		</title>
	</head>
	<body>
		<a href="https://google.com">
			<code>
				eg
			</code>
		</a>
	</body>
</html>
`,
	},
}

func TestForEachNode(t *testing.T) {
	for _, testCase := range testCases {
		buf := new(bytes.Buffer)
		doc, err := html.Parse(testCase.Input)
		if err != nil {
			fmt.Fprintf(os.Stderr, "outline: %v\n", err)
			os.Exit(1)
		}
		forEachNode(doc, startElement, endElement, buf)
		fmt.Println(string(buf.Bytes()))

		doc, err = html.Parse(strings.NewReader(string(buf.Bytes())))
		if err != nil {
			t.Errorf("invalid result(cannot parse). testCaes:%v, actual:%v", testCase, string(buf.Bytes()))
			fmt.Fprintf(os.Stderr, "outline: %v\n", err)
			os.Exit(1)
		}
	}
}
