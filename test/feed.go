package test

import (
	"time"

	"github.com/gorilla/feeds"
)

// Feed with test data
var Feed = &feeds.Feed{
	Title:   "Test feed",
	Link:    &feeds.Link{Href: "localhost"},
	Created: time.Now().UTC().Truncate(time.Minute),
	Author: &feeds.Author{
		Name:  "test",
		Email: "test@localhost",
	},
	Image: &feeds.Image{
		Link:  "localhost",
		Url:   "test",
		Title: "test",
	},
	Items: []*feeds.Item{&feeds.Item{
		Title:       "test",
		Link:        &feeds.Link{Href: "http://localhost/test-page"},
		Id:          "http://localhost/test-page",
		Description: "test",
		Author: &feeds.Author{
			Name:  "test",
			Email: "test@localhost",
		},
		Updated: time.Now().UTC().Truncate(time.Minute),
		Enclosure: &feeds.Enclosure{
			Url:    "test",
			Length: "test",
			Type:   "test",
		},
	}},
}
