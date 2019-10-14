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

type Parser struct{}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getLink(n *html.Node) Link {
	var l Link

	for _, attr := range n.Attr {
		if attr.Key == "href" {
			l.Href, l.Text = attr.Val, n.FirstChild.Data
		}
	}

	return l
}

func (p Parser) Parse(r io.Reader) []Link {
	doc, err := html.Parse(r)
	checkError(err)

	var f func(*html.Node)
	var l []Link

	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			l = append(l, getLink(n))
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

	return l
}
