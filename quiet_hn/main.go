package main

import (
	"flag"
	"fmt"
	"github.com/jkuma/gophercises/quiet_hn/hn"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func main() {
	// parse flags
	var port, numStories int
	flag.IntVar(&port, "port", 3000, "the port to start the web server on")
	flag.IntVar(&numStories, "num_stories", 30, "the number of top stories to display")
	flag.Parse()

	tpl := template.Must(template.ParseFiles("./index.gohtml"))

	http.HandleFunc("/", handler(numStories, tpl))

	// Start the server
	fmt.Println("Server listening on port", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

func handler(numStories int, tpl *template.Template) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		var client hn.Client
		ids, err := client.TopItems()
		if err != nil {
			http.Error(w, "Failed to load top stories", http.StatusInternalServerError)
			return
		}
		var stories []item
		c := make(chan item, numStories)
		ids = ids[:numStories]

		for _, id := range ids {
			go getStoryItem(client, id, c)
		}

		for i := 0; i < numStories; i++ {
			stories = append(stories, <-c)
		}

		data := templateData{
			Stories: sort(stories, ids),
			Time:    time.Now().Sub(start),
		}
		err = tpl.Execute(w, data)
		if err != nil {
			http.Error(w, "Failed to process the template", http.StatusInternalServerError)
			return
		}
	})
}

func sort(items []item, ids []int) []item {
	sorted := make([]item, len(items))

	for k, id := range ids {
		for _, item := range items {
			if id == item.ID {
				sorted[k] = item
			}
		}
	}

	return sorted
}

func getStoryItem(client hn.Client, id int, c chan item) {
	hnItem, err := client.GetItem(id)
	if err == nil {
		item := parseHNItem(hnItem)
		if isStoryLink(item) {
			c <- item
		}
	}
}

func isStoryLink(item item) bool {
	return item.Type == "story" && item.URL != ""
}

func parseHNItem(hnItem hn.Item) item {
	ret := item{Item: hnItem}
	url, err := url.Parse(ret.URL)
	if err == nil {
		ret.Host = strings.TrimPrefix(url.Hostname(), "www.")
	}
	return ret
}

// item is the same as the hn.Item, but adds the Host field
type item struct {
	hn.Item
	Host string
}

type templateData struct {
	Stories []item
	Time    time.Duration
}
