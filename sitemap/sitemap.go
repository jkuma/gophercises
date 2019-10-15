package sitemap

import "github.com/jkuma/gophercises/linkparser"

// Add a map that stores every links
// Create a

type Location struct {
	Parent string
	Url string
	Depth int16
}

type Map struct {
	Url string
	Depth int16
}

func (s Map) Build() {
	var l map[string][]Location
	c := make(chan []linkparser.Link)
}