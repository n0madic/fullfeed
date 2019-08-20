// Package fullfeed converts partial feed to full-text feed
package fullfeed

import (
	"bytes"
	"fmt"
	"net/url"
	"strings"
	"sync"
	"unicode/utf8"

	"github.com/PuerkitoBio/goquery"
	"github.com/antchfx/htmlquery"
	readability "github.com/go-shiori/go-readability"
	"github.com/gorilla/feeds"
	"golang.org/x/net/html"
)

// GetFullFeed with full text content
func GetFullFeed(config Config) (feed *feeds.Feed, errors []error) {
	feed, err := LoadSourceFeed(config)
	if err != nil {
		return nil, []error{err}
	}

	if config.MaxWorkers == 0 {
		config.MaxWorkers = 10
	}
	semaphore := make(chan struct{}, config.MaxWorkers)

	var mutex = &sync.Mutex{}
	var wg sync.WaitGroup

	for idx := range feed.Items {
		semaphore <- struct{}{}
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			fullContent, err := GetFullContent(config, feed.Items[idx].Link.Href)
			if err != nil {
				mutex.Lock()
				errors = append(errors, err)
				mutex.Unlock()
			} else {
				if fullContent != "" && utf8.RuneCountInString(fullContent) > utf8.RuneCountInString(feed.Items[idx].Description) {
					feed.Items[idx].Description = fullContent
				}
			}
			<-semaphore
		}(idx)
	}
	wg.Wait()

	return feed, errors
}

// GetFullContent for the specified link
func GetFullContent(config Config, link string) (fullContent string, err error) {
	content, err := GetURL(link)
	if err != nil {
		return "", err
	}

	if config.BaseHref == "" {
		config.BaseHref = config.URL
	}
	baseURL, err := url.Parse(config.BaseHref)
	if err != nil {
		return "", err
	}

	switch strings.ToLower(config.Method) {
	case "query":
		doc, err := goquery.NewDocumentFromReader(content)
		if err == nil {
			html, err := doc.Find(config.MethodQuery).Html()
			if err == nil {
				fullContent = html
			}
		}
	case "xpath":
		doc, err := htmlquery.Parse(content)
		if err == nil {
			list := htmlquery.Find(doc, config.MethodQuery)
			if len(list) > 0 {
				var b bytes.Buffer
				err = html.Render(&b, list[0])
				if err == nil {
					fullContent = b.String()
				}
			}
		}
	default:
		doc, err := readability.FromReader(content, baseURL.String())
		if err == nil {
			fullContent = doc.Content
		}
	}

	if fullContent != "" {
		doc, err := goquery.NewDocumentFromReader(strings.NewReader(fullContent))
		if err != nil {
			return "", err
		}

		doc.Find("script").Remove()

		if len(config.Filters.Selectors) > 0 {
			doc.Find(strings.Join(config.Filters.Selectors, ", ")).Remove()
		}

		if len(config.Filters.Text) > 0 {
			var searchText []string
			for _, text := range config.Filters.Text {
				searchText = append(searchText, fmt.Sprintf("div:contains('%s')", text))
			}
			doc.Find(strings.Join(searchText, ", ")).Remove()
		}

		makeAllLinksAbsolute(baseURL, doc)

		fullContent, err = doc.Html()
		if err != nil {
			return "", err
		}
	}

	return fullContent, nil
}
