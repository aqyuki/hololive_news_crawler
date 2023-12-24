package hololive

import (
	"errors"
	"fmt"

	"github.com/gocolly/colly"
)

var (
	ErrEmptyTargetURL = errors.New("empty target url")
	ErrNotFoundNext   = errors.New("not found next page")
)

// HololiveExtractor extracts the urls from official web site of hololive
// target url is https://hololive.hololivepro.com/news
type HololiveNewsURLExtractor struct{}

// ExtractURLs extracts the urls from official web site of hololive
func (e *HololiveNewsURLExtractor) ExtractURLs(url string) ([]string, error) {
	if url == "" {
		return nil, ErrEmptyTargetURL
	}

	// collectedURLs is a slice of collected URLs
	var collectedURLs []string

	// targetURL is a target url
	var targetURL string = url

	for {

		urls, err := e.extractURLs(targetURL)
		if err != nil {
			return nil, fmt.Errorf("failed to extract urls: %w", err)
		}
		collectedURLs = append(collectedURLs, urls...)

		nextPageURL, err := e.extractNextPage(url)
		if errors.Is(err, ErrNotFoundNext) {
			break
		} else if err != nil {
			return nil, fmt.Errorf("failed to extract next page: %w", err)
		}
		targetURL = nextPageURL
	}
	return collectedURLs, nil
}

// extractURLs extracts the urls from the target url
func (e *HololiveNewsURLExtractor) extractURLs(url string) ([]string, error) {
	// collectedURLs is a slice of collected URLs
	var collectedURLs []string

	// collector is a colly collector
	var collector *colly.Collector = colly.NewCollector()
	collector.AllowedDomains = allowedDomains

	collector.OnHTML("#container > main > div.in_news.fade_b.isPlay_b > div > ul > li", func(h *colly.HTMLElement) {
		href, exist := h.DOM.Find("a[href]").Attr("href")
		if !exist {
			return
		}
		collectedURLs = append(collectedURLs, href)
	})

	if err := collector.Visit(url); err != nil {
		return nil, fmt.Errorf("failed to visit: %w", err)
	}

	return collectedURLs, nil
}

// extractNextPage extracts the next page url
func (e *HololiveNewsURLExtractor) extractNextPage(url string) (string, error) {

	// next is a next page url
	var next string

	// collector is a colly collector
	var collector *colly.Collector = colly.NewCollector()
	collector.AllowedDomains = allowedDomains

	collector.OnHTML("#pagination > a.next.page-numbers", func(h *colly.HTMLElement) {
		next = h.Attr("href")
	})

	if err := collector.Visit(url); err != nil {
		return "", fmt.Errorf("failed to visit: %w", err)
	}
	return next, nil
}
