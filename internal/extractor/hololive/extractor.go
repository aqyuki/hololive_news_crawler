package hololive

import (
	"github.com/aqyuki/hololive_news_crawler/internal/crawler"
	"github.com/gocolly/colly"
)

// HololiveExtractor extracts the information from official web site of hololive
// target url is https://hololive.hololivepro.com/news
type HololiveNewsExtractor struct{}

// Scrape extracts the information from official web site of hololive
func (e *HololiveNewsExtractor) Scrape(url string) (*crawler.Site, error) {
	// site holds the information of the scraped site
	var site *crawler.Site = new(crawler.Site)

	// collector is a colly collector
	var collector *colly.Collector = colly.NewCollector()
	collector.AllowedDomains = allowedDomains

	collector.OnHTML("title", func(h *colly.HTMLElement) {
		site.Title = h.Text
	})

	collector.OnHTML("meta[name]", func(h *colly.HTMLElement) {
		switch h.Attr("name") {
		case "description":
			site.Description = h.Attr("content")
		}
	})

	collector.OnResponse(func(r *colly.Response) {
		site.URL = r.Request.URL.String()
	})

	if err := collector.Visit(url); err != nil {
		return nil, err
	}

	return site, nil
}
