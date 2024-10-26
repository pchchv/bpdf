package cache_test

import (
	"fmt"
	"testing"

	"github.com/pchchv/bpdf/internal/cache"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	sut := cache.New()

	assert.NotNil(t, sut)
	assert.Equal(t, "*cache.cache", fmt.Sprintf("%T", sut))
}
