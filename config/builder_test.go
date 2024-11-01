package config_test

import (
	"fmt"
	"testing"

	"github.com/pchchv/bpdf/config"
	"github.com/pchchv/bpdf/consts/extension"
	"github.com/pchchv/bpdf/consts/fontfamily"
	"github.com/pchchv/bpdf/consts/fontstyle"
	"github.com/pchchv/bpdf/consts/generation"
	"github.com/pchchv/bpdf/consts/provider"
	"github.com/stretchr/testify/assert"
)

func TestNewBuilder(t *testing.T) {
	sut := config.NewBuilder()

	assert.NotNil(t, sut)
	assert.Equal(t, "*config.CfgBuilder", fmt.Sprintf("%T", sut))
}

func TestBuilder_Build(t *testing.T) {
	sut := config.NewBuilder()

	cfg := sut.Build()

	assert.Equal(t, provider.Gofpdf, cfg.ProviderType)
	assert.Equal(t, 210.0, cfg.Dimensions.Width)
	assert.Equal(t, 297.0, cfg.Dimensions.Height)
	assert.Equal(t, 10.0, cfg.Margins.Top)
	assert.Equal(t, 10.0, cfg.Margins.Left)
	assert.Equal(t, 10.0, cfg.Margins.Right)
	assert.Equal(t, fontfamily.Arial, cfg.DefaultFont.Family)
	assert.Equal(t, 10.0, cfg.DefaultFont.Size)
	assert.Equal(t, fontstyle.Normal, cfg.DefaultFont.Style)
	assert.Equal(t, 0, cfg.DefaultFont.Color.Red)
	assert.Equal(t, 0, cfg.DefaultFont.Color.Green)
	assert.Equal(t, 0, cfg.DefaultFont.Color.Blue)
	assert.Nil(t, cfg.CustomFonts)
	assert.Equal(t, generation.Sequential, cfg.GenerationMode)
	assert.Equal(t, 1, cfg.ChunkWorkers)
	assert.Equal(t, false, cfg.Debug)
	assert.Equal(t, 12, cfg.MaxGridSize)
	assert.Nil(t, cfg.PageNumber)
	assert.Nil(t, cfg.Protection)
	assert.False(t, cfg.Compression)
	assert.Nil(t, cfg.Metadata)
	assert.Nil(t, cfg.BackgroundImage)
	assert.False(t, cfg.DisableAutoPageBreak)
}

func TestBuilder_WithKeywords(t *testing.T) {
	t.Run("when keywords is empty, should ignore", func(t *testing.T) {
		sut := config.NewBuilder()

		cfg := sut.WithKeywords("", true).Build()

		assert.Nil(t, cfg.Metadata)
	})

	t.Run("when author valid, should apply", func(t *testing.T) {
		sut := config.NewBuilder()

		cfg := sut.WithKeywords("keyword", true).Build()

		assert.Equal(t, "keyword", cfg.Metadata.KeywordsStr.Text)
		assert.Equal(t, true, cfg.Metadata.KeywordsStr.UTF8)
	})
}

func TestCfgBuilder_WithBackgroundImage(t *testing.T) {
	t.Run("when with background, should apply", func(t *testing.T) {
		sut := config.NewBuilder()

		cfg := sut.WithBackgroundImage([]byte{1, 2, 3}, extension.Png).Build()

		assert.Equal(t, []byte{1, 2, 3}, cfg.BackgroundImage.Bytes)
		assert.Equal(t, extension.Png, cfg.BackgroundImage.Extension)
	})
}

func TestBuilder_WithDisableAutoPageBreak(t *testing.T) {
	t.Run("when disable auto page break is false, should not change the default value", func(t *testing.T) {
		sut := config.NewBuilder()

		cfg := sut.WithDisableAutoPageBreak(false).Build()

		assert.Equal(t, false, cfg.DisableAutoPageBreak)
	})

	t.Run("when disable auto page break is true, should change the default value", func(t *testing.T) {
		sut := config.NewBuilder()

		cfg := sut.WithDisableAutoPageBreak(true).Build()

		assert.Equal(t, true, cfg.DisableAutoPageBreak)
	})
}
