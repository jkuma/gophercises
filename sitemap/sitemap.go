package sitemap

import (
	"errors"
	"fmt"
	"github.com/jkuma/gophercises/linkparser"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type Location struct {
	ParentUrl string
	Url       string
	Depth     int16
}

type Sitemap struct {
	Url      string
	MaxDepth int16
}

func (s Sitemap) Build() map[string][]Location {
	m := make(map[string][]Location)
	var f func(m *map[string][]Location, u string, d int16)

	if s.Url != "" {
		c := make(chan []Location)

		f = func(m *map[string][]Location, u string, d int16) {
			p := *m

			go locations(u, d, c)
			p[u] = <-c

			for _, loc := range p[u] {
				if loc.Depth <= s.MaxDepth {
					if _, ok := p[loc.Url]; !ok {
						f(m, loc.Url, loc.Depth)
					}
				}
			}
		}
		f(&m, s.Url, 0)
	}

	return m
}

func (s Sitemap) PrintXML() {
	var w string
	l := s.Build()

	for u, locs := range l {
		w += fmt.Sprintf("domain: %v\nlocations: %v", u, locs)
	}

	_, _ = io.Copy(os.Stdout, strings.NewReader(w))
}

// Returns a list of locations for given url.
func locations(parent string, depth int16, c chan []Location) {
	var locations []Location
	resp, err := http.Get(parent)

	if err != nil {
		log.Fatalf("Could not resolve %v, \nerror message: %v", parent, err)
	}

	l := linkparser.Parser{}.Parse(resp.Body)

	for _, link := range l {
		if u, e := createUrlFromHref(parent, link.Href); e == nil{
			locations = append(locations, Location{
				ParentUrl: parent,
				Url: u,
				Depth: depth+1,
			})
		}

	}

	c <- locations
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
