package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/aqyuki/hololive_news_crawler/internal/crawler"
	"github.com/aqyuki/hololive_news_crawler/internal/extractor/hololive"
	"github.com/aqyuki/hololive_news_crawler/pkg/logging"
)

const (
	targetURL = "https://hololive.hololivepro.com/news"
)

func main() {

	logger := logging.NewLogger()
	logger.Info("start crawling")

	extractor := hololive.HololiveNewsExtractor{}
	urlExtractor := hololive.HololiveNewsURLExtractor{}
	ctx := logging.WithLogger(context.Background(), logger)

	c := crawler.NewCrawler(&extractor)
	c.SetURLExtractor(&urlExtractor)

	sites, err := c.Crawl(ctx, targetURL)
	if err != nil {
		logger.Error("failed to crawl", slog.Any("error", err))
		os.Exit(1)
	}

	for _, site := range sites {
		fmt.Printf("Title : %s\t URL : %s\t Description : %s\t\n", site.Title, site.URL, site.Description)
	}
}
