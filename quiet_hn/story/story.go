package story

import (
	"errors"
	"github.com/jkuma/gophercises/quiet_hn/hn"
	"net/url"
	"strings"
	"time"
)

// Story is the same as the hn.Item, but adds the Host field
type Story struct {
	hn.Item
	Host string
}

type result struct {
	Index int
	Story Story
	Error error
}

var (
	cache           []Story
	cacheExpiration time.Time
)

func GetTopStories(numStories int) ([]Story, error) {
	var client hn.Client
	stories := make([]Story, numStories)
	ids, err := client.TopItems()
	if err != nil {
		return nil, errors.New("failed to load top stories")
	}

	c := make(chan result, numStories)

	for k, id := range ids[:numStories] {
		go func(id int, index int) {
			var client hn.Client
			hnItem, err := client.GetItem(id)
			if err == nil {
				Story := parseHNItem(hnItem)
				if isStoryLink(Story) {
					c <- result{Index: index, Story: Story}
					return
				}
			}

			c <- result{Error: err, Index: index}
		}(id, k)
	}

	for i := 0; i < numStories; i++ {
		result := <-c
		if result.Error == nil {
			stories[result.Index] = result.Story
		}
	}

	return stories, nil
}

func isStoryLink(story Story) bool {
	return story.Type == "story" && story.URL != ""
}

func parseHNItem(hnItem hn.Item) Story {
	ret := Story{Item: hnItem}
	url2, err := url.Parse(ret.URL)
	if err == nil {
		ret.Host = strings.TrimPrefix(url2.Hostname(), "www.")
	}
	return ret
}
