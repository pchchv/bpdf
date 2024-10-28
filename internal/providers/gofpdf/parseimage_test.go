package gofpdf_test

import (
	"testing"

	"github.com/pchchv/bpdf/consts/extension"
	"github.com/pchchv/bpdf/internal/providers/gofpdf"
	"github.com/stretchr/testify/assert"
)

func TestFromBytes(t *testing.T) {
	t.Run("when extension is not valid, should return error", func(t *testing.T) {
		img, err := gofpdf.FromBytes([]byte{1, 2, 3}, "invalid")

		assert.Nil(t, img)
		assert.NotNil(t, err)
	})

	t.Run("when extension is not valid, should return error", func(t *testing.T) {
		img, err := gofpdf.FromBytes([]byte{1, 2, 3}, extension.Jpg)

		assert.NotNil(t, img)
		assert.Nil(t, err)
	})
}
