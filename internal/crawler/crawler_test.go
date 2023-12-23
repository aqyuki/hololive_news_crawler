package crawler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCrawler(t *testing.T) {
	t.Parallel()

	t.Run("should panic if extractor is nil", func(t *testing.T) {
		t.Parallel()

		assert.Panics(t, func() {
			NewCrawler(nil)
		})
	})

	t.Run("should return Crawler", func(t *testing.T) {
		t.Parallel()

		mock := struct{ Extractor }{}
		actual := NewCrawler(&mock)
		assert.NotNil(t, actual)
	})
}
