package crawler

// Extractor extract the data from the web site
type Extractor interface {
	// Scrape extract the data from the web site and return the Site
	Scrape(url string) (*Site, error)
}
