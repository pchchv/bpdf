package border_test

import (
	"testing"

	"github.com/pchchv/bpdf/consts/border"
	"github.com/stretchr/testify/assert"
)

func TestType_IsValid(t *testing.T) {
	t.Run("When type is empty, should not be valid", func(t *testing.T) {
		borderType := border.Border("")

		assert.False(t, borderType.IsValid())
	})

	t.Run("When type is full, should be valid", func(t *testing.T) {
		borderType := border.Full

		assert.True(t, borderType.IsValid())
	})

	t.Run("When type is left, should be valid", func(t *testing.T) {
		borderType := border.Left

		assert.True(t, borderType.IsValid())
	})

	t.Run("When type is top, should be valid", func(t *testing.T) {
		borderType := border.Top

		assert.True(t, borderType.IsValid())
	})

	t.Run("When type is right, should be valid", func(t *testing.T) {
		borderType := border.Right

		assert.True(t, borderType.IsValid())
	})

	t.Run("When type is bottom, should be valid", func(t *testing.T) {
		borderType := border.Bottom

		assert.True(t, borderType.IsValid())
	})
}
