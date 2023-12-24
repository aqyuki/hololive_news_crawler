package crawler

// Site holds the data of the web site
type Site struct {
	// URL is the URL of the web site
	URL string `json:"url" yaml:"url"`
	// TItle is the title of the web site
	Title string `json:"title" yaml:"title"`
	// Description is the description of the web site
	Description string `json:"description" yaml:"description"`
}

// NewSite creates a new Site
func NewSite(url, title, description string) *Site {
	return &Site{
		URL:         url,
		Title:       title,
		Description: description,
	}
}
