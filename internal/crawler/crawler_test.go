package crawler

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

// successURLExtractorMock is a mock of URLExtractor
type successURLExtractorMock struct{}

// ExtractURLs is a mock of ExtractURLs
func (e *successURLExtractorMock) ExtractURLs(url string) ([]string, error) {
	return []string{url}, nil
}

// errorURLExtractorMock is a mock of URLExtractor
type errorURLExtractorMock struct{}

// ExtractURLs is a mock of ExtractURLs
func (e *errorURLExtractorMock) ExtractURLs(url string) ([]string, error) {
	return nil, errors.New("mock")
}

type successExtractorMock struct{}

func (e *successExtractorMock) Scrape(url string) (*Site, error) {
	return &Site{
		URL:         url,
		Title:       "title",
		Description: "description",
	}, nil
}

type errorExtractorMock struct{}

func (e *errorExtractorMock) Scrape(url string) (*Site, error) {
	return nil, errors.New("mock")
}

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

func TestCrawler_SetURLExtractor(t *testing.T) {
	t.Parallel()

	t.Run("should return error if url extractor is nil", func(t *testing.T) {
		t.Parallel()

		c := NewCrawler(&successExtractorMock{})
		err := c.SetURLExtractor(nil)
		assert.ErrorIs(t, err, ErrURLExtractorIsNil)
	})

	t.Run("should be success if url extractor is not nil", func(t *testing.T) {
		t.Parallel()

		c := NewCrawler(&successExtractorMock{})
		err := c.SetURLExtractor(&successURLExtractorMock{})
		assert.NoError(t, err)
		assert.EqualValues(t, &successURLExtractorMock{}, c.urlExtractor)
	})
}

func TestCrawler_Crawl(t *testing.T) {
	t.Parallel()

	t.Run("should be success", func(t *testing.T) {
		t.Parallel()

		c := NewCrawler(&successExtractorMock{})
		c.SetURLExtractor(&successURLExtractorMock{})
		actual, err := c.Crawl(context.Background(), "https://example.com")
		assert.NoError(t, err)
		assert.EqualValues(t, []Site{
			{
				URL:         "https://example.com",
				Title:       "title",
				Description: "description",
			},
		}, actual)
	})

	t.Run("should return error if url extractor is nil", func(t *testing.T) {
		t.Parallel()

		c := NewCrawler(&errorExtractorMock{})
		actual, err := c.Crawl(context.Background(), "https://example.com")
		assert.Error(t, err)
		assert.Nil(t, actual)
	})

	t.Run("should return error if url extractor returns error", func(t *testing.T) {
		t.Parallel()

		c := NewCrawler(&successExtractorMock{})
		c.SetURLExtractor(&errorURLExtractorMock{})
		actual, err := c.Crawl(context.Background(), "https://example.com")
		assert.Error(t, err)
		assert.Nil(t, actual)
	})
}
