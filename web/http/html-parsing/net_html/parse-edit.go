// parse partial html code and add attributes
// input:  <iframe src="https://docs.google.com/document/d/12345"></iframe>
// output: <iframe src="https://docs.google.com/document/d/12345" width="800" height="1200"></iframe>

package main

import (
	"bytes"
	"fmt"
	"log"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

var data = `<iframe src="https://docs.google.com/document/d/12345"></iframe>`

func main() {

	h, err := updateIFrameAttr(data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(h)

}

// updateIFrameAttr updates iframe html code with the size attributes
func updateIFrameAttr(iFrameString string) (string, error) {
	ctx := &html.Node{
		Type:     html.ElementNode,
		DataAtom: atom.Div,
		Data:     "div",
	}
	doc, err := html.ParseFragment(strings.NewReader(iFrameString), ctx)
	if err != nil {
		return "", err
	}

	for _, n := range doc {
		if !(n.Type == html.ElementNode && n.Data == "iframe") {
			continue
		}
		n.Attr = append(n.Attr, html.Attribute{Key: "width", Val: "800"}, html.Attribute{Key: "height", Val: "1200"})

		var b bytes.Buffer
		if err := html.Render(&b, n); err != nil {
			return "", err
		}

		return b.String(), nil
	}

	return "", fmt.Errorf("iframe not found")
}
