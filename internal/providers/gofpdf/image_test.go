package gofpdf_test

import (
	"fmt"
	"testing"

	gofpdf2 "github.com/pchchv/bpdf/internal/providers/gofpdf"
	"github.com/pchchv/bpdf/mocks"
	"github.com/stretchr/testify/assert"
)

func TestNewImage(t *testing.T) {
	image := gofpdf2.NewImage(mocks.NewFpdf(t), mocks.NewMath(t))

	assert.NotNil(t, image)
	assert.Equal(t, fmt.Sprintf("%T", image), "*gofpdf.image")
}
