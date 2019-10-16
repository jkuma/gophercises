package sitemap

import (
	"fmt"
	"github.com/jkuma/gophercises/linkparser"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

type Location struct {
	Parent string
	Link   linkparser.Link
	Depth  int16
}

type Sitemap struct {
	Url   string
	Depth int16
}

func (s Sitemap) Build() map[string][]Location {
	m := make(map[string][]Location)

	if s.Url != "" {
		c := make(chan []Location)

		go locations(s.Url, s.Depth, c)
		m[s.Url] = <-c
	}

	return m
}

func (s Sitemap) PrintXML() {
	var w string
	l := s.Build()

	for url, locs := range l {
		w += fmt.Sprintf("domain: %v\nlocations: %v", url, locs)
	}

	_, _ = io.Copy(os.Stdout, strings.NewReader(w))
}

// Returns a list of locations for the given parent url.
func locations(parent string, depth int16, c chan []Location) {
	var locs []Location
	resp, err := http.Get(parent)

	if err != nil {
		log.Fatalf("Could not resolve %v, \nerror message: %v", parent, err)
	}

	l := linkparser.Parser{}.Parse(resp.Body)

	for _, link := range l {
		locs = append(locs, Location{
			Parent: parent,
			Link: linkparser.Link{
				Href: link.Href,
				Text: link.Text,
			},
			Depth: depth,
		})
	}

	c <- locs
}
