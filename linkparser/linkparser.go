package linkparser

import (
	"golang.org/x/net/html"
	"io"
	"log"
)

type Link struct {
	Href string
	Text string
}

type Parser struct {}

func checkError(err error)  {
	if err != nil {
		log.Fatal(err)
	}
}

func (p Parser) Parse(r io.Reader) []Link {
	doc, err := html.Parse(r)
	checkError(err)

	var f func(*html.Node)
	var l []Link

	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			l = append(l, Link{
				Href: n.Attr[0].Val,
				Text: n.FirstChild.Data,
			})
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

	return l
}
