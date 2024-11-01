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
	"github.com/pchchv/bpdf/consts/orientation"
	"github.com/pchchv/bpdf/consts/pagesize"
	"github.com/pchchv/bpdf/consts/protection"
	"github.com/pchchv/bpdf/consts/provider"
	"github.com/pchchv/bpdf/core/entity"
	"github.com/pchchv/bpdf/properties"
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

func TestBuilder_WithAuthor(t *testing.T) {
	t.Run("when author is empty, should ignore", func(t *testing.T) {
		sut := config.NewBuilder()

		cfg := sut.WithAuthor("", true).Build()

		assert.Nil(t, cfg.Metadata)
	})

	t.Run("when author valid, should apply", func(t *testing.T) {
		sut := config.NewBuilder()

		cfg := sut.WithAuthor("author", true).Build()

		assert.Equal(t, "author", cfg.Metadata.Author.Text)
		assert.Equal(t, true, cfg.Metadata.Author.UTF8)
	})
}

func TestCfgBuilder_WithBottomMargin(t *testing.T) {
	t.Run("when bottom is invalid, should not change the default value", func(t *testing.T) {
		sut := config.NewBuilder()

		cfg := sut.WithBottomMargin(-1).Build()

		assert.Equal(t, 20.0025, cfg.Margins.Bottom)
	})

	t.Run("when bottom is valid, should change the default value", func(t *testing.T) {
		sut := config.NewBuilder()

		cfg := sut.WithBottomMargin(5).Build()

		assert.Equal(t, 5.0, cfg.Margins.Bottom)
	})
}

func TestBuilder_WithSubject(t *testing.T) {
	t.Run("when subject is empty, should ignore", func(t *testing.T) {
		sut := config.NewBuilder()

		cfg := sut.WithSubject("", true).Build()

		assert.Nil(t, cfg.Metadata)
	})

	t.Run("when subject valid, should apply", func(t *testing.T) {
		sut := config.NewBuilder()

		cfg := sut.WithSubject("subject", true).Build()

		assert.Equal(t, "subject", cfg.Metadata.Subject.Text)
		assert.Equal(t, true, cfg.Metadata.Subject.UTF8)
	})
}

func TestBuilder_WithOrientation(t *testing.T) {
	t.Run("when using default page size and orientation is not set, should use vertical", func(t *testing.T) {
		sut := config.NewBuilder()

		cfg := sut.Build()

		assert.Equal(t, 210.0, cfg.Dimensions.Width)
		assert.Equal(t, 297.0, cfg.Dimensions.Height)
		assert.True(t, cfg.Dimensions.Height > cfg.Dimensions.Width)
	})

	t.Run("when using default page size and orientation is set to horizontal, should use horizontal", func(t *testing.T) {
		sut := config.NewBuilder()

		cfg := sut.WithOrientation(orientation.Horizontal).Build()

		assert.Equal(t, 297.0, cfg.Dimensions.Width)
		assert.Equal(t, 210.0, cfg.Dimensions.Height)
		assert.True(t, cfg.Dimensions.Width > cfg.Dimensions.Height)
	})

	t.Run("when using default page size and orientation is not set, should use vertical", func(t *testing.T) {
		sut := config.NewBuilder()

		cfg := sut.WithPageSize(pagesize.A5).Build()

		assert.Equal(t, 148.4, cfg.Dimensions.Width)
		assert.Equal(t, 210.0, cfg.Dimensions.Height)
		assert.True(t, cfg.Dimensions.Height > cfg.Dimensions.Width)
	})

	t.Run("when using default page size and orientation is set to horizontal, should use horizontal", func(t *testing.T) {
		sut := config.NewBuilder()

		cfg := sut.WithPageSize(pagesize.A5).WithOrientation(orientation.Horizontal).Build()

		assert.Equal(t, 210.0, cfg.Dimensions.Width)
		assert.Equal(t, 148.4, cfg.Dimensions.Height)
		assert.True(t, cfg.Dimensions.Width > cfg.Dimensions.Height)
	})
}

func TestBuilder_WithCreator(t *testing.T) {
	t.Run("when creator is empty, should ignore", func(t *testing.T) {
		sut := config.NewBuilder()

		cfg := sut.WithCreator("", true).Build()

		assert.Nil(t, cfg.Metadata)
	})

	t.Run("when creator valid, should apply", func(t *testing.T) {
		sut := config.NewBuilder()

		cfg := sut.WithCreator("creator", true).Build()

		assert.Equal(t, "creator", cfg.Metadata.Creator.Text)
		assert.Equal(t, true, cfg.Metadata.Creator.UTF8)
	})
}

// nolint:dupl
// dupl is good here
func TestBuilder_WithConcurrentMode(t *testing.T) {
	t.Run("when chunk size is invalid, should not change the default value", func(t *testing.T) {
		sut := config.NewBuilder()

		cfg := sut.WithConcurrentMode(-1).Build()

		assert.Equal(t, generation.Sequential, cfg.GenerationMode)
		assert.Equal(t, 1, cfg.ChunkWorkers)
	})

	t.Run("when chunk size is valid, should change the default value", func(t *testing.T) {
		sut := config.NewBuilder()

		cfg := sut.WithConcurrentMode(7).Build()

		assert.Equal(t, generation.Concurrent, cfg.GenerationMode)
		assert.Equal(t, 7, cfg.ChunkWorkers)
	})

	t.Run("when chunk size is valid, should override sequential", func(t *testing.T) {
		sut := config.NewBuilder()
		sut.WithSequentialMode()

		cfg := sut.WithConcurrentMode(7).Build()

		assert.Equal(t, generation.Concurrent, cfg.GenerationMode)
		assert.Equal(t, 7, cfg.ChunkWorkers)
	})

	t.Run("when chunk size is valid, should override sequential low memory", func(t *testing.T) {
		sut := config.NewBuilder()
		sut.WithSequentialLowMemoryMode(5)

		cfg := sut.WithConcurrentMode(7).Build()

		assert.Equal(t, generation.Concurrent, cfg.GenerationMode)
		assert.Equal(t, 7, cfg.ChunkWorkers)
	})
}

func TestCfgBuilder_WithSequentialMode(t *testing.T) {
	t.Run("when sequential, should apply sequential", func(t *testing.T) {
		sut := config.NewBuilder()

		cfg := sut.WithSequentialMode().Build()

		assert.Equal(t, generation.Sequential, cfg.GenerationMode)
		assert.Equal(t, 1, cfg.ChunkWorkers)
	})

	t.Run("when sequential, should override sequential low memory", func(t *testing.T) {
		sut := config.NewBuilder()
		sut.WithSequentialLowMemoryMode(10)

		cfg := sut.WithSequentialMode().Build()

		assert.Equal(t, generation.Sequential, cfg.GenerationMode)
		assert.Equal(t, 1, cfg.ChunkWorkers)
	})

	t.Run("when sequential, should override concurrent", func(t *testing.T) {
		sut := config.NewBuilder()
		sut.WithConcurrentMode(10)

		cfg := sut.WithSequentialMode().Build()

		assert.Equal(t, generation.Sequential, cfg.GenerationMode)
		assert.Equal(t, 1, cfg.ChunkWorkers)
	})
}

func TestCfgBuilder_WithProtection(t *testing.T) {
	t.Run("when with protection, should apply correct", func(t *testing.T) {
		sut := config.NewBuilder()

		cfg := sut.WithProtection(protection.Copy, "password", "owner-password").Build()

		assert.Equal(t, protection.Copy, cfg.Protection.Type)
		assert.Equal(t, "password", cfg.Protection.UserPassword)
		assert.Equal(t, "owner-password", cfg.Protection.OwnerPassword)
	})
}

func TestCfgBuilder_WithCompression(t *testing.T) {
	t.Run("when with compression, should apply correct", func(t *testing.T) {
		sut := config.NewBuilder()

		cfg := sut.WithCompression(true).Build()

		assert.True(t, cfg.Compression)
	})
}

func TestCfgBuilder_WithPageNumber(t *testing.T) {
	t.Run("when using empty, should apply default", func(t *testing.T) {
		sut := config.NewBuilder()

		cfg := sut.WithPageNumber().Build()

		assert.Equal(t, "{current} / {total}", cfg.PageNumber.Pattern)
		assert.Equal(t, properties.Bottom, cfg.PageNumber.Place)
		assert.Equal(t, fontstyle.Normal, cfg.PageNumber.Style)
		assert.Equal(t, fontfamily.Arial, cfg.PageNumber.Family)
		assert.Equal(t, 10.0, cfg.PageNumber.Size)
		assert.Equal(t, 0, cfg.PageNumber.Color.Red)
		assert.Equal(t, 0, cfg.PageNumber.Color.Green)
		assert.Equal(t, 0, cfg.PageNumber.Color.Blue)
	})

	t.Run("when string pattern doesn´t have current, should apply default pattern", func(t *testing.T) {
		sut := config.NewBuilder()
		pageNumber := properties.PageNumber{
			Pattern: "{total}",
		}

		cfg := sut.WithPageNumber(pageNumber).Build()

		assert.Equal(t, "{current} / {total}", cfg.PageNumber.Pattern)
	})

	t.Run("when string pattern doesn´t have total, should apply default pattern", func(t *testing.T) {
		sut := config.NewBuilder()
		pageNumber := properties.PageNumber{
			Pattern: "{current}",
		}

		cfg := sut.WithPageNumber(pageNumber).Build()

		assert.Equal(t, "{current} / {total}", cfg.PageNumber.Pattern)
	})

	t.Run("when string pattern is correct, should apply pattern", func(t *testing.T) {
		sut := config.NewBuilder()
		pageNumber := properties.PageNumber{
			Pattern: "Page {current} of {total}",
		}

		cfg := sut.WithPageNumber(pageNumber).Build()

		assert.Equal(t, "Page {current} of {total}", cfg.PageNumber.Pattern)
	})

	t.Run("when place is not valid, should apply default", func(t *testing.T) {
		sut := config.NewBuilder()
		pageNumber := properties.PageNumber{
			Place: "invalid",
		}

		cfg := sut.WithPageNumber(pageNumber).Build()

		assert.Equal(t, properties.Bottom, cfg.PageNumber.Place)
	})

	t.Run("when place is valid, should apply config", func(t *testing.T) {
		sut := config.NewBuilder()
		pageNumber := properties.PageNumber{
			Place: properties.Top,
		}

		cfg := sut.WithPageNumber(pageNumber).Build()

		assert.Equal(t, properties.Top, cfg.PageNumber.Place)
	})
}

// nolint:dupl // dupl is good here
func TestCfgBuilder_WithSequentialLowMemoryMode(t *testing.T) {
	t.Run("when chunk size is invalid, should not change the default value", func(t *testing.T) {

		sut := config.NewBuilder()

		cfg := sut.WithSequentialLowMemoryMode(-1).Build()

		assert.Equal(t, generation.Sequential, cfg.GenerationMode)
		assert.Equal(t, 1, cfg.ChunkWorkers)
	})

	t.Run("when chunk size is valid, should change the default value", func(t *testing.T) {

		sut := config.NewBuilder()

		cfg := sut.WithSequentialLowMemoryMode(7).Build()

		assert.Equal(t, generation.SequentialLowMemory, cfg.GenerationMode)
		assert.Equal(t, 7, cfg.ChunkWorkers)
	})

	t.Run("when chunk size is valid, should override sequential low memory", func(t *testing.T) {

		sut := config.NewBuilder()
		sut.WithSequentialMode()

		cfg := sut.WithSequentialLowMemoryMode(7).Build()

		assert.Equal(t, generation.SequentialLowMemory, cfg.GenerationMode)
		assert.Equal(t, 7, cfg.ChunkWorkers)
	})

	t.Run("when chunk size is valid, should override concurrent", func(t *testing.T) {

		sut := config.NewBuilder()
		sut.WithConcurrentMode(5)

		cfg := sut.WithSequentialLowMemoryMode(7).Build()

		assert.Equal(t, generation.SequentialLowMemory, cfg.GenerationMode)
		assert.Equal(t, 7, cfg.ChunkWorkers)
	})
}

func TestBuilder_WithDebug(t *testing.T) {
	sut := config.NewBuilder()

	cfg := sut.WithDebug(true).Build()

	assert.Equal(t, true, cfg.Debug)
}

func TestBuilder_WithMaxGridSize(t *testing.T) {
	t.Run("when max grid size is invalid, should not change the default value", func(t *testing.T) {
		sut := config.NewBuilder()

		cfg := sut.WithMaxGridSize(-1).Build()

		assert.Equal(t, 12, cfg.MaxGridSize)
	})

	t.Run("when max grid size is valid, should change the default value", func(t *testing.T) {
		sut := config.NewBuilder()

		cfg := sut.WithMaxGridSize(8).Build()

		assert.Equal(t, 8, cfg.MaxGridSize)
	})
}

func TestBuilder_WithDefaultFont(t *testing.T) {
	t.Run("when fontstyle is nil, should not change the default value", func(t *testing.T) {
		sut := config.NewBuilder()

		cfg := sut.WithDefaultFont(nil).Build()

		assert.Equal(t, fontfamily.Arial, cfg.DefaultFont.Family)
		assert.Equal(t, 10.0, cfg.DefaultFont.Size)
		assert.Equal(t, fontstyle.Normal, cfg.DefaultFont.Style)
		assert.Equal(t, 0, cfg.DefaultFont.Color.Red)
		assert.Equal(t, 0, cfg.DefaultFont.Color.Green)
		assert.Equal(t, 0, cfg.DefaultFont.Color.Blue)
	})

	t.Run("when family is filled, should change the default value", func(t *testing.T) {
		sut := config.NewBuilder()

		cfg := sut.WithDefaultFont(&properties.Font{
			Family: "new family",
		}).Build()

		assert.Equal(t, "new family", cfg.DefaultFont.Family)
		assert.Equal(t, 10.0, cfg.DefaultFont.Size)
		assert.Equal(t, fontstyle.Normal, cfg.DefaultFont.Style)
		assert.Equal(t, 0, cfg.DefaultFont.Color.Red)
		assert.Equal(t, 0, cfg.DefaultFont.Color.Green)
		assert.Equal(t, 0, cfg.DefaultFont.Color.Blue)
	})

	t.Run("when style is filled, should change the default value", func(t *testing.T) {
		sut := config.NewBuilder()

		cfg := sut.WithDefaultFont(&properties.Font{
			Style: fontstyle.Bold,
		}).Build()

		assert.Equal(t, fontfamily.Arial, cfg.DefaultFont.Family)
		assert.Equal(t, 10.0, cfg.DefaultFont.Size)
		assert.Equal(t, fontstyle.Bold, cfg.DefaultFont.Style)
		assert.Equal(t, 0, cfg.DefaultFont.Color.Red)
		assert.Equal(t, 0, cfg.DefaultFont.Color.Green)
		assert.Equal(t, 0, cfg.DefaultFont.Color.Blue)
	})

	t.Run("when size is filled, should change the default value", func(t *testing.T) {
		sut := config.NewBuilder()

		cfg := sut.WithDefaultFont(&properties.Font{
			Size: 13,
		}).Build()

		assert.Equal(t, fontfamily.Arial, cfg.DefaultFont.Family)
		assert.Equal(t, 13.0, cfg.DefaultFont.Size)
		assert.Equal(t, fontstyle.Normal, cfg.DefaultFont.Style)
		assert.Equal(t, 0, cfg.DefaultFont.Color.Red)
		assert.Equal(t, 0, cfg.DefaultFont.Color.Green)
		assert.Equal(t, 0, cfg.DefaultFont.Color.Blue)
	})

	t.Run("when color is filled, should change the default value", func(t *testing.T) {
		sut := config.NewBuilder()

		cfg := sut.WithDefaultFont(&properties.Font{
			Color: &properties.Color{Red: 10, Green: 10, Blue: 10},
		}).Build()

		assert.Equal(t, fontfamily.Arial, cfg.DefaultFont.Family)
		assert.Equal(t, 10.0, cfg.DefaultFont.Size)
		assert.Equal(t, fontstyle.Normal, cfg.DefaultFont.Style)
		assert.Equal(t, 10, cfg.DefaultFont.Color.Red)
		assert.Equal(t, 10, cfg.DefaultFont.Color.Green)
		assert.Equal(t, 10, cfg.DefaultFont.Color.Blue)
	})
}
