package fullfeed

import (
	"net/url"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func TestFixLazyImage(t *testing.T) {
	html := `<img src="data:image/svg+xml;base64,FooBar" data-src="http://www.example.com/image.png"/>`
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(html))
	fixLazyImage(doc)
	doc.Find("img").Each(func(i int, sel *goquery.Selection) {
		src := sel.AttrOr("src", "")
		if src != "http://www.example.com/image.png" {
			t.Errorf("Lazy image not fixed, got \"%s\"", src)
		}
	})
}

func TestMakeAllLinksAbsolute(t *testing.T) {
	html := `<a href="index.html"><img src="image.png" data-src="image.png"/></a>`
	baseURL, _ := url.Parse("http://www.example.com")
	attributes := []string{"src", "data-src", "href"}

	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(html))
	makeAllLinksAbsolute(baseURL, doc)
	doc.Find("a,img").Each(func(i int, sel *goquery.Selection) {
		for _, attr := range attributes {
			if val, exists := sel.Attr(attr); exists {
				if !strings.HasPrefix(val, "http://www.example.com") {
					t.Errorf("Expected find absolute URL in %s=\"%s\"", attr, val)
				}
			}
		}
	})
}

func TestStringIsFiltered(t *testing.T) {
	if !stringIsFiltered("test", []string{"test"}) {
		t.Error("Expected find 'test' in ['test']")
	}

	if stringIsFiltered("test", []string{"test2"}) {
		t.Error("Not expected to find 'test' in ['test2']")
	}
}
