package fs

import (
	"encoding/json"

	"github.com/aqyuki/hololive_news_crawler/internal/crawler"
	"gopkg.in/yaml.v3"
)

type convertor func([]crawler.Site) ([]byte, error)

// convertToJSON converts the given data to JSON format binary.
func convertToJSON(data []crawler.Site) ([]byte, error) {
	b, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// convertToYAML converts the given data to YAML format binary.
func convertToYAML(data []crawler.Site) ([]byte, error) {
	b, err := yaml.Marshal(data)
	if err != nil {
		return nil, err
	}
	return b, nil
}
