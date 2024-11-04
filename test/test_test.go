package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	t.Run("when called first, should setup singleton and set t", func(t *testing.T) {
		sut := New(t)

		assert.Equal(t, t, sut.t)
	})

	t.Run("when called not first, should use singleton and set t", func(t *testing.T) {
		_ = New(t)

		sut := New(t)

		assert.Equal(t, t, sut.t)
	})
}
