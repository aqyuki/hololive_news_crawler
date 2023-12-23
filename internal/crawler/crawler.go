package crawler

// Crawler crawl the web site
type Crawler struct {
	// extractor is a object that implements the Extractor interface
	extractor Extractor
}

// NewCrawler creates a new Crawler
func NewCrawler(ext Extractor) *Crawler {
	if ext == nil {
		panic("ext is nil")
	}

	return &Crawler{
		extractor: ext,
	}
}
