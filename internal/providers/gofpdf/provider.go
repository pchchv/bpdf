package gofpdf

import (
	"bytes"
	"errors"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/pchchv/bpdf/consts/barcode"
	"github.com/pchchv/bpdf/consts/extension"
	"github.com/pchchv/bpdf/core"
	"github.com/pchchv/bpdf/core/entity"
	"github.com/pchchv/bpdf/internal/cache"
	merror "github.com/pchchv/bpdf/internal/error"
	"github.com/pchchv/bpdf/internal/providers/gofpdf/cellwriter"
	"github.com/pchchv/bpdf/internal/providers/gofpdf/fpdfwrapper"
	"github.com/pchchv/bpdf/properties"
)

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

func (g *provider) AddMatrixCode(code string, cell *entity.Cell, prop *properties.Rect) {
	img, err := g.loadCode(code, "matrix-code-", g.code.GenDataMatrix)
	if err != nil {
		g.text.Add("could not generate matrixcode", cell, merror.DefaultErrorText)
		return
	}

	if err = g.image.Add(img, cell, g.cfg.Margins, prop, extension.Png, false); err != nil {
		g.fpdf.ClearError()
		g.text.Add("could not add matrixcode to document", cell, merror.DefaultErrorText)
	}
}

func (g *provider) AddQrCode(code string, cell *entity.Cell, prop *properties.Rect) {
	img, err := g.loadCode(code, "qr-code-", g.code.GenQr)
	if err != nil {
		g.text.Add("could not generate qrcode", cell, merror.DefaultErrorText)
		return
	}

	if err = g.image.Add(img, cell, g.cfg.Margins, prop, extension.Png, false); err != nil {
		g.fpdf.ClearError()
		g.text.Add("could not add qrcode to document", cell, merror.DefaultErrorText)
	}
}

func (g *provider) AddBarCode(code string, cell *entity.Cell, prop *properties.Barcode) {
	image, err := g.cache.GetImage(g.getBarcodeImageName(fmt.Sprintf("bar-code-%s", code), prop), extension.Png)
	if err != nil {
		if image, err = g.code.GenBar(code, cell, prop); err != nil {
			g.text.Add("could not generate barcode", cell, merror.DefaultErrorText)
			return
		}
	}

	g.cache.AddImage(g.getBarcodeImageName(fmt.Sprintf("bar-code-%s", code), prop), image)
	if err = g.image.Add(image, cell, g.cfg.Margins, prop.ToRectProp(), extension.Png, false); err != nil {
		g.fpdf.ClearError()
		g.text.Add("could not add barcode to document", cell, merror.DefaultErrorText)
	}
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

// GetDimensionsByQrCode is responsible for obtaining the dimensions of an QrCode
// If the image cannot be loaded, an error is returned
func (g *provider) GetDimensionsByQrCode(code string) (*entity.Dimensions, error) {
	img, err := g.loadCode(code, "qr-code-", g.code.GenQr)
	if err != nil {
		return nil, err
	}

	imgInfo, _ := g.image.GetImageInfo(img, extension.Png)
	if imgInfo == nil {
		return nil, errors.New("could not read image options, maybe path/name is wrong")
	}

	return &entity.Dimensions{Width: imgInfo.Width(), Height: imgInfo.Height()}, nil
}

func (g *provider) GenerateBytes() ([]byte, error) {
	var buffer bytes.Buffer
	return buffer.Bytes(), g.fpdf.Output(&buffer)
}

// GetDimensionsByImageByte is responsible for obtaining the dimensions of an image
// If the image cannot be loaded, an error is returned
func (g *provider) GetDimensionsByImageByte(bytes []byte, extension extension.Extension) (*entity.Dimensions, error) {
	img, err := FromBytes(bytes, extension)
	if err != nil {
		return nil, err
	}

	imgInfo, _ := g.image.GetImageInfo(img, extension)
	if imgInfo == nil {
		return nil, errors.New("could not read image options, maybe path/name is wrong")
	}

	return &entity.Dimensions{Width: imgInfo.Width(), Height: imgInfo.Height()}, nil
}

func (g *provider) getBarcodeImageName(code string, prop *properties.Barcode) string {
	if prop == nil {
		return code + string(barcode.Code128)
	}

	return code + string(prop.Type)
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
