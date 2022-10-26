package main

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/ericchiang/css"
	"golang.org/x/net/html"
)

var data = `
<p>
  <h2 id="foo">a header</h2>
  <h2 id="bar">another header</h2>
</p>`

func main() {
	sel, err := css.Parse("h2#foo")
	if err != nil {
		panic(err)
	}
	node, err := html.Parse(strings.NewReader(data))
	if err != nil {
		panic(err)
	}

	for _, ele := range sel.Select(node) {
		b := &bytes.Buffer{}
		html.Render(b, ele)
		fmt.Printf("node: %s\n", b.String())
	}
	fmt.Println()
}
