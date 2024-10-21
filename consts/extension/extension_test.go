package extension_test

import (
	"testing"

	"github.com/pchchv/bpdf/consts/extension"
	"github.com/stretchr/testify/assert"
)

func TestType_IsValid(t *testing.T) {
	t.Run("when type is empty, should not be valid", func(t *testing.T) {
		extensionType := extension.Extension("")
		assert.False(t, extensionType.IsValid())
	})

	t.Run("when type is jpg, should be valid", func(t *testing.T) {
		extensionType := extension.Jpg
		assert.True(t, extensionType.IsValid())
	})

	t.Run("when type is jpeg, should be valid", func(t *testing.T) {
		extensionType := extension.Jpeg
		assert.True(t, extensionType.IsValid())
	})

	t.Run("when type is png, should be valid", func(t *testing.T) {
		extensionType := extension.Png
		assert.True(t, extensionType.IsValid())
	})
}
