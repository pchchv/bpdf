package gofpdf_test

import (
	"fmt"
	"testing"

	"github.com/pchchv/bpdf/internal/providers/gofpdf"
	"github.com/stretchr/testify/assert"
)

func TestNewLine(t *testing.T) {
	sut := gofpdf.NewLine(nil)

	assert.NotNil(t, sut)
	assert.Equal(t, "*gofpdf.line", fmt.Sprintf("%T", sut))
}
