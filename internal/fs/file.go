package fs

import (
	"os"

	"github.com/aqyuki/hololive_news_crawler/internal/crawler"
)

// Encode is the type of encoding.
type Encode int

const (
	EncodeJSON Encode = iota // EncodeJSON is the JSON encoding.
	EncodeYAML               // EncodeYAML is the YAML encoding.
)

// EncodeSave encodes the given data and saves it to the given filename.
func EncodeSave(filename string, data []crawler.Site, encode Encode) error {
	var c convertor
	switch encode {
	case EncodeJSON:
		c = convertToJSON
	case EncodeYAML:
		c = convertToYAML
	}

	b, err := c(data)
	if err != nil {
		return err
	}
	return saveToFile(b, filename)
}

// saveToFile saves the given byte slice to the given filename.
func saveToFile(b []byte, filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(b)
	if err != nil {
		return err
	}

	return nil
}
