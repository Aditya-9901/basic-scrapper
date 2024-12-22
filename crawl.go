package main

import (
	"log"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func crawl(url string, ch chan string, chFininshed chan bool) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		chFininshed <- true
	}()

	b := resp.Body
	defer b.Close()

	z := html.NewTokenizer(b)
	for {
		tt := z.Next()
		switch tt {
		case html.ErrorToken:
			return
		case html.StartTagToken:
			t := z.Token()
			if strings.TrimSpace(t.Data) == "a" {
				atrs := t.Attr
				for _, k := range atrs {
					if k.Key == "href" && (strings.Contains(k.Val, "http") || (strings.Contains(k.Val, "https"))) {
						ch <- k.Val
					}
				}
			}
		}

	}
}
