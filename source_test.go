package fullfeed

import (
	"testing"
	"time"

	"github.com/n0madic/fullfeed/test"
)

func TestLoadSourceFeed(t *testing.T) {
	addr := "localhost:8093"
	go test.RunTestHTTPServer(addr)

	atomURL := "http://" + addr + "/atom"
	rssURL := "http://" + addr + "/rss"

	testConfig := Config{
		URL:         rssURL,
		Description: "test",
	}

	_, err := LoadSourceFeed(Config{URL: rssURL + "notfound"})
	if err == nil {
		t.Error("Feed loading error expected")
	}

	feed, err := LoadSourceFeed(testConfig)
	if err != nil {
		t.Fatal("No error expected:", err)
	}

	if feed.Title != test.Feed.Title {
		t.Errorf("Expect '%s' in feed Title, got '%s'", test.Feed.Title, test.Feed.Title)
	}
	if feed.Link.Href != rssURL {
		t.Errorf("Expect feed Link %s in feed Link.Href, got %s", rssURL, feed.Link.Href)
	}
	if feed.Description != testConfig.Description {
		t.Errorf("Expect feed description '%s', got '%s'", testConfig.Description, feed.Description)
	}
	if feed.Created.IsZero() {
		t.Errorf("Expect created time '%s', got '%s'", time.Now().UTC().Truncate(time.Minute), feed.Created)
	}

	if diff := test.Diff(test.Feed.Author, feed.Author); diff != "" {
		t.Error(diff)
	}
	if diff := test.Diff(test.Feed.Image, feed.Image); diff != "" {
		t.Error(diff)
	}

	feed, err = LoadSourceFeed(Config{URL: atomURL})
	if err != nil {
		t.Fatal("No error expected:", err)
	}

	if feed.Updated.IsZero() {
		t.Errorf("Expect updated time '%s', got '%s'", time.Now().UTC().Truncate(time.Minute), feed.Updated)
	}
	if diff := test.Diff(test.Feed.Items, feed.Items); diff != "" {
		t.Error(diff)
	}

}
