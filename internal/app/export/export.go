package export

import (
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/client"
	"github.com/geziyor/geziyor/export"
)

func exporter() {
	page := 1
	for page = 1; page < 50; page++ {
		geziyor.NewGeziyor(&geziyor.Options{
			StartURLs: []string{"https://habr.com/ru/hubs/go/articles/page" + strconv.Itoa(page)},
			ParseFunc: postingParse,
			Exporters: []export.Exporter{&export.JSON{}},
			UserAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/116.0.0.0 Safari/537.36 OPR/102.0.0.0 (Edition Yx GX)",
		}).Start()
	}
}

func postingParse(g *geziyor.Geziyor, r *client.Response) {
	r.HTMLDoc.Find("article.tm-articles-list__item").Each(func(i int, s *goquery.Selection) {
		src, _ := s.Find("img").Last().Attr("src")
		timedate, _ := s.Find("time").Attr("title")
		g.Exports <- map[string]interface{}{
			"Name":  s.Find("h2.tm-title_h2").Text(),
			"Body":  s.Find("p").First().Text(),
			"Image": src,
			"Time":  timedate,
		}
	})

	if href, ok := r.HTMLDoc.Find("li.next > a").Attr("href"); ok {
		g.Get(r.JoinURL(href), postingParse)
	}
}
