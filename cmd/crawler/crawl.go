package main

import (
	"fmt"
	"os"

	"github.com/aqyuki/hololive_news_crawler/internal/crawler"
	"github.com/aqyuki/hololive_news_crawler/internal/extractor/hololive"
)

const (
	targetURL = "https://hololive.hololivepro.com/news"
)

func main() {

	fmt.Println("Start crawling...")

	extractor := hololive.HololiveNewsExtractor{}
	urlExtractor := hololive.HololiveNewsURLExtractor{}

	c := crawler.NewCrawler(&extractor)
	c.SetURLExtractor(&urlExtractor)

	sites, err := c.Crawl(targetURL)
	if err != nil {
		fmt.Printf("failed to crawl %s: %v\n", targetURL, err)
		os.Exit(1)
	}

	for _, site := range sites {
		fmt.Printf("Title : %s\t URL : %s\t Description : %s\t\n", site.Title, site.URL, site.Description)
	}
}
