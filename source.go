package fullfeed

import (
	"github.com/gorilla/feeds"
	"github.com/mmcdole/gofeed"
)

// LoadSourceFeed return source feed
func LoadSourceFeed(config Config) (feed *feeds.Feed, err error) {
	fp := gofeed.NewParser()
	sourceFeed, err := fp.ParseURL(config.URL)
	if err != nil {
		return feed, err
	}

	feed = &feeds.Feed{
		Title:       sourceFeed.Title,
		Link:        &feeds.Link{Href: config.URL},
		Description: sourceFeed.Description,
	}

	if feed.Description == "" {
		feed.Description = config.Description
	}

	if sourceFeed.PublishedParsed != nil {
		feed.Created = *sourceFeed.PublishedParsed
	}

	if sourceFeed.Author != nil {
		feed.Author = &feeds.Author{
			Name:  sourceFeed.Author.Name,
			Email: sourceFeed.Author.Email,
		}
	}

	if sourceFeed.Image != nil {
		feed.Image = &feeds.Image{
			Link:  sourceFeed.Link,
			Url:   sourceFeed.Image.URL,
			Title: sourceFeed.Image.Title,
		}
	}

	for _, item := range sourceFeed.Items {
		if !stringIsFiltered(item.Title, config.Filters.Titles) &&
			!stringIsFiltered(item.Description, config.Filters.Descriptions) {

			var author *feeds.Author
			if item.Author != nil {
				author = &feeds.Author{
					Name:  item.Author.Name,
					Email: item.Author.Email,
				}
			}

			var enclosure *feeds.Enclosure
			if item.Enclosures != nil && len(item.Enclosures) > 0 {
				enclosure = &feeds.Enclosure{
					Length: item.Enclosures[0].Length,
					Type:   item.Enclosures[0].Type,
					Url:    item.Enclosures[0].URL,
				}
			}

			feed.Add(&feeds.Item{
				Title:       item.Title,
				Link:        &feeds.Link{Href: item.Link},
				Id:          item.Link,
				Author:      author,
				Created:     *item.PublishedParsed,
				Description: item.Description,
				Enclosure:   enclosure,
			})
		}
	}

	return
}
