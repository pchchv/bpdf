package gofpdf_test

import (
	"fmt"
	"testing"

	"github.com/pchchv/bpdf/internal/providers/gofpdf"
	"github.com/pchchv/bpdf/mocks"
	"github.com/stretchr/testify/assert"
)

func TestNewText(t *testing.T) {
	text := gofpdf.NewText(mocks.NewFpdf(t), mocks.NewMath(t), mocks.NewFont(t))

	assert.NotNil(t, text)
	assert.Equal(t, fmt.Sprintf("%T", text), "*gofpdf.text")
}
