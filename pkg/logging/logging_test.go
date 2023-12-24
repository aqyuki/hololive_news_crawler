package logging

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewLogger(t *testing.T) {
	t.Parallel()

	actual := NewLogger()
	assert.NotNil(t, actual)
}

func TestContext(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	logger := FromContext(ctx)
	assert.NotNil(t, logger)

	ctx = WithLogger(ctx, logger)
	assert.Equal(t, logger, FromContext(ctx))
}
