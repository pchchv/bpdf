package bpdf_test

import (
	"fmt"
	"testing"

	"github.com/pchchv/bpdf"
	"github.com/pchchv/bpdf/config"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	t.Run("new default", func(t *testing.T) {
		sut := bpdf.New()

		assert.NotNil(t, sut)
		assert.Equal(t, "*bpdf.Bpdf", fmt.Sprintf("%T", sut))
	})

	t.Run("new with config", func(t *testing.T) {
		cfg := config.NewBuilder().
			Build()

		sut := bpdf.New(cfg)

		assert.NotNil(t, sut)
		assert.Equal(t, "*bpdf.Bpdf", fmt.Sprintf("%T", sut))
	})

	t.Run("new with config an concurrent mode on", func(t *testing.T) {
		cfg := config.NewBuilder().
			WithConcurrentMode(7).
			Build()

		sut := bpdf.New(cfg)

		assert.NotNil(t, sut)
		assert.Equal(t, "*bpdf.Bpdf", fmt.Sprintf("%T", sut))
	})

	t.Run("new with config an low memory mode on", func(t *testing.T) {
		cfg := config.NewBuilder().
			WithSequentialLowMemoryMode(10).
			Build()

		sut := bpdf.New(cfg)

		assert.NotNil(t, sut)
		assert.Equal(t, "*bpdf.Bpdf", fmt.Sprintf("%T", sut))
	})
}
