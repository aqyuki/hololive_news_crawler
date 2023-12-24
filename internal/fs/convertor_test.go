package fs

import (
	"encoding/json"
	"testing"

	"github.com/aqyuki/hololive_news_crawler/internal/crawler"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func TestConvert(t *testing.T) {
	t.Parallel()

	t.Run("convertToJSON", func(t *testing.T) {
		t.Parallel()

		data := []crawler.Site{
			{
				URL:         "https://example.com",
				Title:       "example",
				Description: "example",
			},
		}

		b, err := json.Marshal(data)
		if err != nil {
			t.Fatal(err)
		}
		actual, err := convertToJSON(data)
		assert.NoError(t, err)
		assert.EqualValues(t, b, actual)
	})

	t.Run("convertToYAML", func(t *testing.T) {
		t.Parallel()

		data := []crawler.Site{
			{
				URL:         "https://example.com",
				Title:       "example",
				Description: "example",
			},
		}

		b, err := yaml.Marshal(data)
		if err != nil {
			t.Fatal(err)
		}
		actual, err := convertToYAML(data)
		assert.NoError(t, err)
		assert.EqualValues(t, b, actual)
	})

}
