package crawler

import (
	"errors"
	"fmt"
)

var (
	// ErrURLExtractorIsNil is an error that URLExtractor is nil
	ErrURLExtractorIsNil = errors.New("url extractor is nil")
)

// Crawler crawl the web site
type Crawler struct {
	// extractor is a object that implements the Extractor interface
	extractor Extractor
	// urlExtractor is a object that implements the URLExtractor interface
	urlExtractor URLExtractor
}

// SetURLExtractor set the URLExtractor
func (c *Crawler) SetURLExtractor(ext URLExtractor) error {
	if ext == nil {
		return ErrURLExtractorIsNil
	}
	c.urlExtractor = ext
	return nil
}

// Crawl crawl the web site
func (c *Crawler) Crawl(url string) ([]Site, error) {

	// targetURLs is a slice of URLs which are target of crawling
	var targetURLs []string
	// collected is a slice of collected web sites
	var collected = make([]Site, 0)

	targetURLs, err := c.urlExtractor.ExtractURLs(url)
	if err != nil {
		return nil, fmt.Errorf("failed to create url list : %w", err)
	}

	for _, target := range targetURLs {
		site, err := c.extractor.Scrape(target)
		if err != nil {
			return nil, fmt.Errorf("failed to scrape : %w", err)
		}
		collected = append(collected, *site)
	}

	return collected, nil
}

// NewCrawler creates a new Crawler
func NewCrawler(ext Extractor) *Crawler {
	if ext == nil {
		panic("ext is nil")
	}
	return &Crawler{
		extractor:    ext,
		urlExtractor: &DefaultURLExtractor{},
	}
}
