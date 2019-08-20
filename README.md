# fullfeed

Convert partial feed to full-text feed with golang

## Usage

```go
import "github.com/n0madic/fullfeed"

feed, errors := fullfeed.GetFullFeed(fullfeed.Config{
    Method:      "query",
    MethodQuery: ".article",
    URL:         "https://blog.golang.org/feed.atom",
})

if len(errors) > 0 {
    fmt.Println(errors)
    return
}

fmt.Println(feed.ToRss())
```
