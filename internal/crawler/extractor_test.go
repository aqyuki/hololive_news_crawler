package crawler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultExtractor_ExtractURLs(t *testing.T) {
	t.Parallel()

	const target = "https://example.com"
	expected := []string{target}
	actual, err := (&DefaultURLExtractor{}).ExtractURLs(target)
	assert.NoError(t, err)
	if assert.NotNil(t, actual) {
		assert.Equal(t, expected, actual)
	}
}
