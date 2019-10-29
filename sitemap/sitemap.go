package sitemap

import (
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/jkuma/gophercises/linkparser"
	"log"
	"net/http"
	"net/url"
	"os"
)

const xmlns = "http://www.sitemaps.org/schemas/sitemap/0.9"

type Location struct {
	ParentUrl string
	Url       string
	Depth     int
}

type Sitemap struct {
	Url      string
	MaxDepth int
}

type loc struct {
	Value string `xml:"loc"`
}

type urlset struct {
	Urls  []loc  `xml:"url"`
	Xmlns string `xml:"xmlns,attr"`
}

func (s Sitemap) Build() map[string][]Location {
	m := make(map[string][]Location)
	var f func(m *map[string][]Location, u string, d int)

	f = func(m *map[string][]Location, u string, d int) {
		p := *m

		p[u] = locations(u, d)

		for _, loc := range p[u] {
			if loc.Depth < s.MaxDepth {
				if _, ok := p[loc.Url]; !ok {
					f(m, loc.Url, loc.Depth)
				}
			}
		}
	}
	f(&m, s.Url, 0)

	return m
}

func (s Sitemap) PrintXML() {
	l := s.Build()

	toXml := urlset{
		Xmlns: xmlns,
	}

	for u, _ := range l {
		toXml.Urls = append(toXml.Urls, loc{u})
	}

	fmt.Print(xml.Header)
	enc := xml.NewEncoder(os.Stdout)
	_ = enc.Encode(toXml)
}

// Returns a list of locations for given url.
func locations(parent string, depth int) []Location {
	var locations []Location
	resp, err := http.Get(parent)

	if err != nil {
		log.Fatalf("Could not resolve %v, \nerror message: %v", parent, err)
	}

	l := linkparser.Parser{}.Parse(resp.Body)

	for _, link := range l {
		if u, e := createUrlFromHref(parent, link.Href); e == nil {
			locations = append(locations, Location{
				ParentUrl: parent,
				Url:       u,
				Depth:     depth + 1,
			})
		}

	}

	return locations
}

func createUrlFromHref(parent string, href string) (string, error) {
	p, err := url.ParseRequestURI(parent)
	checkError(err)

	u, err := url.ParseRequestURI(href)

	if err != nil {
		return href, err
	}

	if u.Scheme == "mailto" {
		return href, errors.New("mailto protocol is not handled")
	}

	u.Host, u.Scheme = p.Hostname(), p.Scheme

	return u.String(), nil
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
