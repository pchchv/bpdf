package gofpdf

import (
	"errors"
	"path/filepath"
	"strings"

	"github.com/pchchv/bpdf/consts/extension"
	"github.com/pchchv/bpdf/core"
	"github.com/pchchv/bpdf/core/entity"
	"github.com/pchchv/bpdf/internal/cache"
	"github.com/pchchv/bpdf/internal/providers/gofpdf/cellwriter"
	"github.com/pchchv/bpdf/internal/providers/gofpdf/fpdfwrapper"
	"github.com/pchchv/bpdf/properties"
)

// import (
// 	"bytes"
// 	"errors"
// 	"fmt"
// 	"path/filepath"
// 	"strings"

// 	"github.com/johnfercher/maroto/v2/pkg/consts/barcode"

// 	"github.com/johnfercher/maroto/v2/internal/cache"
// 	"github.com/johnfercher/maroto/v2/internal/merror"
// 	"github.com/johnfercher/maroto/v2/internal/providers/gofpdf/cellwriter"
// 	"github.com/johnfercher/maroto/v2/internal/providers/gofpdf/fpdfwrapper"
// 	"github.com/johnfercher/maroto/v2/pkg/consts/extension"
// 	"github.com/johnfercher/maroto/v2/pkg/core"
// 	"github.com/johnfercher/maroto/v2/pkg/core/entity"
// 	"github.com/johnfercher/maroto/v2/pkg/props"
// )

type provider struct {
	fpdf       fpdfwrapper.Fpdf
	font       core.Font
	text       core.Text
	code       core.Code
	image      core.Image
	line       core.Line
	cache      cache.Cache
	cellWriter cellwriter.CellWriter
	cfg        *entity.Config
}

func (g *provider) GetLinesQuantity(text string, textProp *properties.Text, colWidth float64) int {
	return g.text.GetLinesQuantity(text, textProp, colWidth)
}

func (g *provider) GetFontHeight(prop *properties.Font) float64 {
	return g.font.GetHeight(prop.Family, prop.Style, prop.Size)
}

func (g *provider) AddText(text string, cell *entity.Cell, prop *properties.Text) {
	g.text.Add(text, cell, prop)
}

func (g *provider) AddLine(cell *entity.Cell, prop *properties.Line) {
	g.line.Add(cell, prop)
}

func (g *provider) CreateRow(height float64) {
	g.fpdf.Ln(height)
}

func (g *provider) CreateCol(width, height float64, config *entity.Config, prop *properties.Cell) {
	g.cellWriter.Apply(width, height, config, prop)
}

func (g *provider) SetProtection(protection *entity.Protection) {
	if protection == nil {
		return
	}
	g.fpdf.SetProtection(byte(protection.Type), protection.UserPassword, protection.OwnerPassword)
}

func (g *provider) SetMetadata(metadata *entity.Metadata) {
	if metadata == nil {
		return
	}

	if metadata.Author != nil {
		g.fpdf.SetAuthor(metadata.Author.Text, metadata.Author.UTF8)
	}

	if metadata.Creator != nil {
		g.fpdf.SetCreator(metadata.Creator.Text, metadata.Creator.UTF8)
	}

	if metadata.Subject != nil {
		g.fpdf.SetSubject(metadata.Subject.Text, metadata.Subject.UTF8)
	}

	if metadata.Title != nil {
		g.fpdf.SetTitle(metadata.Title.Text, metadata.Title.UTF8)
	}

	if metadata.CreationDate != nil {
		g.fpdf.SetCreationDate(*metadata.CreationDate)
	}

	if metadata.KeywordsStr != nil {
		g.fpdf.SetKeywords(metadata.KeywordsStr.Text, metadata.KeywordsStr.UTF8)
	}
}

func (g *provider) SetCompression(compression bool) {
	g.fpdf.SetCompression(compression)
}

// GetDimensionsByImage is responsible for obtaining the dimensions of an image.
// If the image cannot be loaded, an error is returned.
func (g *provider) GetDimensionsByImage(file string) (*entity.Dimensions, error) {
	extensionStr := strings.ToLower(strings.TrimPrefix(filepath.Ext(file), "."))
	img, err := g.loadImage(file, extensionStr)
	if err != nil {
		return nil, err
	}

	imgInfo, _ := g.image.GetImageInfo(img, extension.Extension(extensionStr))
	if imgInfo == nil {
		return nil, errors.New("could not read image options, maybe path/name is wrong")
	}

	return &entity.Dimensions{Width: imgInfo.Width(), Height: imgInfo.Height()}, nil
}

// GetDimensionsByMatrixCode is responsible for obtaining the dimensions of an MatrixCode.
// If the image cannot be loaded, an error is returned.
func (g *provider) GetDimensionsByMatrixCode(code string) (*entity.Dimensions, error) {
	img, err := g.loadCode(code, "matrix-code-", g.code.GenDataMatrix)
	if err != nil {
		return nil, err
	}

	imgInfo, _ := g.image.GetImageInfo(img, extension.Png)
	if imgInfo == nil {
		return nil, errors.New("could not read image options, maybe path/name is wrong")
	}

	return &entity.Dimensions{Width: imgInfo.Width(), Height: imgInfo.Height()}, nil
}

// loadImage is responsible for loading an codes
func (g *provider) loadCode(code, codeType string, generate func(code string) (*entity.Image, error)) (*entity.Image, error) {
	image, err := g.cache.GetImage(codeType+code, extension.Png)
	if err != nil {
		if image, err = generate(code); err != nil {
			return nil, err
		}
	} else {
		return image, nil
	}

	g.cache.AddImage(codeType+code, image)
	return image, nil
}

// loadImage is responsible for loading an image
func (g *provider) loadImage(file, extensionStr string) (*entity.Image, error) {
	image, err := g.cache.GetImage(file, extension.Extension(extensionStr))
	if err == nil {
		return image, err
	}

	if err = g.cache.LoadImage(file, extension.Extension(extensionStr)); err != nil {
		return nil, err
	}

	return g.cache.GetImage(file, extension.Extension(extensionStr))
}
