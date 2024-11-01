package config_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/pchchv/bpdf/config"
	"github.com/pchchv/bpdf/consts/extension"
	"github.com/pchchv/bpdf/consts/fontfamily"
	"github.com/pchchv/bpdf/consts/fontstyle"
	"github.com/pchchv/bpdf/consts/generation"
	"github.com/pchchv/bpdf/consts/pagesize"
	"github.com/pchchv/bpdf/consts/provider"
	"github.com/pchchv/bpdf/core/entity"
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

func TestCfgBuilder_WithCustomFonts(t *testing.T) {
	t.Run("when custom font, should apply", func(t *testing.T) {
		sut := config.NewBuilder()
		customFonts := []*entity.CustomFont{
			{
				Family: "custom",
			},
		}

		cfg := sut.WithCustomFonts(customFonts).Build()

		assert.Equal(t, 1, len(cfg.CustomFonts))
		assert.Equal(t, "custom", cfg.CustomFonts[0].Family)
	})
}

func TestBuilder_WithPageSize(t *testing.T) {
	t.Run("when page size is empty, should not change the default value", func(t *testing.T) {
		sut := config.NewBuilder()

		cfg := sut.WithPageSize("").Build()

		assert.Equal(t, 210.0, cfg.Dimensions.Width)
		assert.Equal(t, 297.0, cfg.Dimensions.Height)
	})

	t.Run("when page size is filled, should change the default value", func(t *testing.T) {
		sut := config.NewBuilder()

		cfg := sut.WithPageSize(pagesize.A2).Build()

		assert.Equal(t, 419.9, cfg.Dimensions.Width)
		assert.Equal(t, 594.0, cfg.Dimensions.Height)
	})
}

func TestBuilder_WithCreationDate(t *testing.T) {
	t.Run("when time is zero, should ignore", func(t *testing.T) {
		sut := config.NewBuilder()

		cfg := sut.WithCreationDate(time.Time{}).Build()

		assert.Nil(t, cfg.Metadata)
	})

	t.Run("when time valid, should apply", func(t *testing.T) {
		sut := config.NewBuilder()
		timeNow := time.Now()

		cfg := sut.WithCreationDate(timeNow).Build()

		assert.Equal(t, &timeNow, cfg.Metadata.CreationDate)
	})
}

func TestBuilder_WithDimensions(t *testing.T) {
	t.Run("when dimensions has invalid width, should not change the default value", func(t *testing.T) {
		sut := config.NewBuilder()

		cfg := sut.WithDimensions(0, 80).Build()

		assert.Equal(t, 210.0, cfg.Dimensions.Width)
		assert.Equal(t, 297.0, cfg.Dimensions.Height)
	})

	t.Run("when dimensions has invalid height, should not change the default value", func(t *testing.T) {
		sut := config.NewBuilder()

		cfg := sut.WithDimensions(80, 0).Build()

		assert.Equal(t, 210.0, cfg.Dimensions.Width)
		assert.Equal(t, 297.0, cfg.Dimensions.Height)
	})

	t.Run("when dimensions has valid values, should change the default value", func(t *testing.T) {
		sut := config.NewBuilder()

		cfg := sut.WithDimensions(80, 80).Build()

		assert.Equal(t, 80.0, cfg.Dimensions.Width)
		assert.Equal(t, 80.0, cfg.Dimensions.Height)
	})

	t.Run("when dimensions are set and page size too, should use dimensions values", func(t *testing.T) {
		sut := config.NewBuilder()

		cfg := sut.WithPageSize(pagesize.A1).WithDimensions(80, 80).Build()

		assert.Equal(t, 80.0, cfg.Dimensions.Width)
		assert.Equal(t, 80.0, cfg.Dimensions.Height)
	})
}

func TestCfgBuilder_WithTopMargin(t *testing.T) {
	t.Run("when top is invalid, should not change the default value", func(t *testing.T) {
		sut := config.NewBuilder()

		cfg := sut.WithTopMargin(-1).Build()

		assert.Equal(t, 10.0, cfg.Margins.Top)
	})

	t.Run("when top is valid, should change the default value", func(t *testing.T) {
		sut := config.NewBuilder()

		cfg := sut.WithTopMargin(5).Build()

		assert.Equal(t, 5.0, cfg.Margins.Top)
	})
}

func TestBuilder_WithTitle(t *testing.T) {
	t.Run("when title is empty, should ignore", func(t *testing.T) {
		sut := config.NewBuilder()

		cfg := sut.WithTitle("", true).Build()

		assert.Nil(t, cfg.Metadata)
	})

	t.Run("when title valid, should apply", func(t *testing.T) {
		sut := config.NewBuilder()

		cfg := sut.WithTitle("title", true).Build()

		assert.Equal(t, "title", cfg.Metadata.Title.Text)
		assert.Equal(t, true, cfg.Metadata.Title.UTF8)
	})
}

func TestCfgBuilder_WithLeftMargin(t *testing.T) {
	t.Run("when left is invalid, should not change the default value", func(t *testing.T) {
		sut := config.NewBuilder()

		cfg := sut.WithLeftMargin(-1).Build()

		assert.Equal(t, 10.0, cfg.Margins.Left)
	})

	t.Run("when left is valid, should change the default value", func(t *testing.T) {
		sut := config.NewBuilder()

		cfg := sut.WithLeftMargin(5).Build()

		assert.Equal(t, 5.0, cfg.Margins.Left)
	})
}

func TestCfgBuilder_WithRightMargin(t *testing.T) {
	t.Run("when right is invalid, should not change the default value", func(t *testing.T) {
		sut := config.NewBuilder()

		cfg := sut.WithRightMargin(-1).Build()

		assert.Equal(t, 10.0, cfg.Margins.Right)
	})

	t.Run("when right is valid, should change the default value", func(t *testing.T) {
		sut := config.NewBuilder()

		cfg := sut.WithRightMargin(5).Build()

		assert.Equal(t, 5.0, cfg.Margins.Right)
	})
}
