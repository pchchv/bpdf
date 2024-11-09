package bpdf_test

import (
	"fmt"
	"testing"

	"github.com/pchchv/bpdf"
	"github.com/stretchr/testify/assert"
)

func TestNewMetricsDecorator(t *testing.T) {
	sut := bpdf.NewMetricsDecorator(nil)

	assert.NotNil(t, sut)
	assert.Equal(t, "*maroto.MetricsDecorator", fmt.Sprintf("%T", sut))
}
