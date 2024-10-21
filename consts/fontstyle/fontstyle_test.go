package fontstyle_test

import (
	"testing"

	"github.com/pchchv/bpdf/consts/fontstyle"
	"github.com/stretchr/testify/assert"
)

func TestType_IsValid(t *testing.T) {
	t.Run("when style is invalid, should be invalid", func(t *testing.T) {
		fontStyle := fontstyle.Fontstyle("invalid")
		assert.False(t, fontStyle.IsValid())
	})

	t.Run("when style is normal, should be valid", func(t *testing.T) {
		fontStyle := fontstyle.Normal
		assert.True(t, fontStyle.IsValid())
	})

	t.Run("when style is bold, should be valid", func(t *testing.T) {
		fontStyle := fontstyle.Bold
		assert.True(t, fontStyle.IsValid())
	})

	t.Run("when style is bold italic, should be valid", func(t *testing.T) {
		fontStyle := fontstyle.BoldItalic
		assert.True(t, fontStyle.IsValid())
	})
}
