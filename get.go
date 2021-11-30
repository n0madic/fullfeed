package fullfeed

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"golang.org/x/net/html/charset"
)

// UserAgent header
var UserAgent string = "fullfeed/1.0"

// GetURL return content from internet or cache
func GetURL(url string) (io.Reader, error) {
	if contentCache != nil {
		if content, ok := contentCache.Get(url); ok {
			return strings.NewReader(content.(string)), nil
		}
	}

	client := http.Client{
		Timeout: time.Duration(30 * time.Second),
	}

	req, err := http.NewRequest("GET", url, nil)
	if err == nil {
		req.Header.Add("User-Agent", UserAgent)

		res, err := client.Do(req)
		if err == nil {
			defer res.Body.Close()

			if res.StatusCode != http.StatusOK {
				return nil, fmt.Errorf("status code for url %s : %s", url, res.Status)
			}

			utf8, err := charset.NewReader(res.Body, res.Header.Get("Content-Type"))
			if err == nil {
				bodyBytes, err := ioutil.ReadAll(utf8)
				if err == nil {
					bodyString := string(bodyBytes)
					if contentCache != nil {
						contentCache.Add(url, bodyString)
					}

					return strings.NewReader(bodyString), nil
				}
			}
		}
	}
	return nil, err
}
