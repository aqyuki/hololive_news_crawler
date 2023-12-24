package crawler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSite(t *testing.T) {
	t.Parallel()

	url := "http://www.example.com"
	title := "Example Title"
	description := "Example Description"

	actual := NewSite(url, title, description)

	if assert.NotNil(t, actual) {
		assert.Equal(t, url, actual.URL)
		assert.Equal(t, title, actual.Title)
		assert.Equal(t, description, actual.Description)
	} else {
		t.Fatal()
	}
}
