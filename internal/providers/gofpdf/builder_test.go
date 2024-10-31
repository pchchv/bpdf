package gofpdf_test

import (
	"fmt"
	"testing"

	"github.com/pchchv/bpdf/consts/fontfamily"
	"github.com/pchchv/bpdf/core/entity"
	"github.com/pchchv/bpdf/internal/fixture"
	"github.com/pchchv/bpdf/internal/providers/gofpdf"
	"github.com/stretchr/testify/assert"
)

func TestNewBuilder(t *testing.T) {
	sut := gofpdf.NewBuilder()

	assert.NotNil(t, sut)
	assert.Equal(t, "*gofpdf.builder", fmt.Sprintf("%T", sut))
}

func TestBuilder_Build(t *testing.T) {
	t.Run("when DisableAutoPageBreak true, should build correctly", func(t *testing.T) {
		sut := gofpdf.NewBuilder()
		font := fixture.FontProp()
		cfg := &entity.Config{
			Dimensions: &entity.Dimensions{
				Width:  100,
				Height: 200,
			},
			Margins: &entity.Margins{
				Left:   10,
				Top:    10,
				Right:  10,
				Bottom: 10,
			},
			DefaultFont: &font,
			CustomFonts: []*entity.CustomFont{
				{
					Family: fontfamily.Arial,
				},
			},
			DisableAutoPageBreak: true,
		}

		dep := sut.Build(cfg, nil)

		assert.NotNil(t, dep)
	})

	t.Run("when DisableAutoPageBreak false, should build correctly", func(t *testing.T) {
		sut := gofpdf.NewBuilder()
		font := fixture.FontProp()
		cfg := &entity.Config{
			Dimensions: &entity.Dimensions{
				Width:  100,
				Height: 200,
			},
			Margins: &entity.Margins{
				Left:   10,
				Top:    10,
				Right:  10,
				Bottom: 10,
			},
			DefaultFont: &font,
			CustomFonts: []*entity.CustomFont{
				{
					Family: fontfamily.Arial,
				},
			},
			DisableAutoPageBreak: false,
		}

		dep := sut.Build(cfg, nil)

		assert.NotNil(t, dep)
	})
}
