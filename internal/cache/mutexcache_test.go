package cache_test

import (
	"fmt"
	"testing"

	"github.com/pchchv/bpdf/internal/cache"
	"github.com/stretchr/testify/assert"
)

func TestNewMutexDecorator(t *testing.T) {
	sut := cache.NewMutexDecorator(nil)

	assert.NotNil(t, sut)
	assert.Equal(t, "*cache.mutexCache", fmt.Sprintf("%T", sut))
}
