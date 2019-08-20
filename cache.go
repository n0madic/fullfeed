package fullfeed

import lru "github.com/hashicorp/golang-lru"

var contentCache *lru.Cache

// InitContentCache setup optional download cache
func InitContentCache(n int) (err error) {
	contentCache, err = lru.New(n)
	return err
}
