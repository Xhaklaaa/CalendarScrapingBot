package main

import (
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/client"
	"github.com/geziyor/geziyor/export"
)

func parsingSites() {
	geziyor.NewGeziyor(&geziyor.Options{
		StartURLs: []string{"https://habr.com/ru/hubs/go/articles/" + strconv.Itoa()},
		ParseFunc: postingParse,
		Exporters: []export.Exporter{&export.JSON{}},
	}).Start()
}

func postingParse(g *geziyor.Geziyor, r *client.Response) {
	r.HTMLDoc.Find("div.quote").Each(func(i int, s *goquery.Selection) {
		g.Exports <- map[string]interface{}{
			"text":   s.Find("span.text").Text(),
			"author": s.Find("small.author").Text(),
		}
	})
	if href, ok := r.HTMLDoc.Find("li.next > a").Attr("href"); ok {
		g.Get(r.JoinURL(href), postingParse)
	}
}
