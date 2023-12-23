package crawler

// Extractor extract the data from the web site
type Extractor interface {
	// Scrape extract the data from the web site and return the Site
	Scrape(url string) (*Site, error)
}

// URLExtractor extract the url from the web site
type URLExtractor interface {
	// ExtractURLs extract the url from the web site and return the urls
	ExtractURLs(url string) ([]string, error)
}

// DefaultURLExtractor is the default implementation of URLExtractor
type DefaultURLExtractor struct{}

// ExtractURLs extract the url from the web site and return the urls
func (e *DefaultURLExtractor) ExtractURLs(url string) ([]string, error) {
	return []string{url}, nil
}
