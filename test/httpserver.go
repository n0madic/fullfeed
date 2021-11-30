package test

import (
	"fmt"
	"net/http"
)

// RunTestHTTPServer with test data
func RunTestHTTPServer(addr string) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "FullFeed Test HTTP server")
	})

	http.HandleFunc("/atom", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/xml")
		atom, err := Feed.ToAtom()
		if err != nil {
			panic(err)
		}
		fmt.Fprint(w, atom)
	})

	http.HandleFunc("/rss", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/xml")
		rss, err := Feed.ToRss()
		if err != nil {
			panic(err)
		}
		fmt.Fprint(w, rss)
	})

	panic(http.ListenAndServe(addr, nil))
}
