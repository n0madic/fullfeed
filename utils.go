package fullfeed

import (
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func absoluteAttr(base *url.URL, sel *goquery.Selection, attr string) {
	if link, ok := sel.Attr(attr); link != "" && ok {
		u, err := url.Parse(link)
		if err == nil && !u.IsAbs() {
			sel.SetAttr(attr, base.ResolveReference(u).String())
		}
	}
}

func fixLazyImage(doc *goquery.Document) {
	doc.Find("img").Each(func(i int, sel *goquery.Selection) {
		if strings.HasPrefix(sel.AttrOr("src", ""), "data:") {
			if val, exists := sel.Attr("data-src"); exists && val != "" {
				sel.SetAttr("src", val)
			}
		}
	})
}

func makeAllLinksAbsolute(base *url.URL, doc *goquery.Document) {
	doc.Find("a,img").Each(func(i int, sel *goquery.Selection) {
		absoluteAttr(base, sel, "src")
		absoluteAttr(base, sel, "data-src")
		absoluteAttr(base, sel, "href")
	})
}

func stringIsFiltered(str string, filters []string) bool {
	for _, filter := range filters {
		if strings.Contains(strings.ToLower(str), strings.ToLower(filter)) {
			return true
		}
	}
	return false
}
