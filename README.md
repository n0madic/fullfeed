# fullfeed

[![Go Reference](https://pkg.go.dev/badge/github.com/n0madic/fullfeed.svg)](https://pkg.go.dev/github.com/n0madic/fullfeed)

Convert partial feed to full-text feed with golang

## Usage

```go
import "github.com/n0madic/fullfeed"

feed, errors := fullfeed.GetFullFeed(fullfeed.Config{
    Method:        fullfeed.QueryMethod,
    MethodRequest: ".article",
    URL:           "https://blog.golang.org/feed.atom",
})

if len(errors) > 0 {
    fmt.Println(errors)
    return
}

fmt.Println(feed.ToRss())
```
