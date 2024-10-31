package gofpdf_test

import (
	"fmt"
	"testing"

	"github.com/pchchv/bpdf/internal/providers/gofpdf"
	"github.com/stretchr/testify/assert"
)

func TestNewBuilder(t *testing.T) {
	sut := gofpdf.NewBuilder()

	assert.NotNil(t, sut)
	assert.Equal(t, "*gofpdf.builder", fmt.Sprintf("%T", sut))
}
