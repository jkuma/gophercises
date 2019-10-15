package sitemap

import (
	"fmt"
	"github.com/jkuma/gophercises/linkparser"
	"io"
	"net/http"
	"os"
	"strings"
)

// Add a map that stores every links
// Create a

type Location struct {
	Parent string
	Link linkparser.Link
	Depth int16
}

type Sitemap struct {
	Url string
	Depth int16
}

func (s Sitemap) Build() map[string][]Location {
	var l map[string][]Location
	c := make(chan []linkparser.Link)

	resp, err := http.Get(s.Url)

	return l
}

func (s Sitemap) PrintXML() {
	var w string
	l := s.Build()

	for url, _ := range l {
		w = fmt.Sprintf("%v \n", url)
	}

	_, _ = io.Copy(os.Stdout, strings.NewReader(w))
}
