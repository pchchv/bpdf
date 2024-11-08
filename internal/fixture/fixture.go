package fixture

import (
	"github.com/pchchv/bpdf/consts/align"
	"github.com/pchchv/bpdf/consts/border"
	"github.com/pchchv/bpdf/consts/breakline"
	"github.com/pchchv/bpdf/consts/extension"
	"github.com/pchchv/bpdf/consts/fontfamily"
	"github.com/pchchv/bpdf/consts/fontstyle"
	"github.com/pchchv/bpdf/consts/linestyle"
	"github.com/pchchv/bpdf/consts/orientation"
	"github.com/pchchv/bpdf/core"
	"github.com/pchchv/bpdf/core/entity"
	"github.com/pchchv/bpdf/node"
	"github.com/pchchv/bpdf/properties"
)

// ColorProp is responsible to give a valid properties.Color.
func ColorProp() properties.Color {
	return properties.Color{
		Red:   100,
		Green: 50,
		Blue:  200,
	}
}

// FontProp is responsible to give a valid properties.Font.
func FontProp() properties.Font {
	colorProp := ColorProp()
	prop := properties.Font{
		Family: fontfamily.Helvetica,
		Style:  fontstyle.Bold,
		Size:   14,
		Color:  &colorProp,
	}
	prop.MakeValid(fontfamily.Arial)
	return prop
}

// RectProp is responsible to give a valid properties.Rect.
func RectProp() properties.Rect {
	prop := properties.Rect{
		Top:     10,
		Left:    10,
		Percent: 98,
		Center:  false,
	}
	prop.MakeValid()
	return prop
}

// TextProp is responsible to give a valid properties.Text.
func TextProp() properties.Text {
	fontProp := FontProp()
	google := "https://www.google.com"
	prop := properties.Text{
		Top:               12,
		Bottom:            13,
		Left:              3,
		Family:            fontProp.Family,
		Style:             fontProp.Style,
		Size:              fontProp.Size,
		Align:             align.Right,
		BreakLineStrategy: breakline.DashStrategy,
		VerticalPadding:   20,
		Color:             fontProp.Color,
		Hyperlink:         &google,
	}
	prop.MakeValid(&fontProp)
	return prop
}

// CellProp is responsible to give a valid properties.Cell.
func CellProp() properties.Cell {
	prop := properties.Cell{
		BackgroundColor: &properties.Color{
			Red:   255,
			Green: 100,
			Blue:  50,
		},
		BorderColor: &properties.Color{
			Red:   200,
			Green: 80,
			Blue:  60,
		},
		BorderType:      border.Left,
		BorderThickness: 0.6,
		LineStyle:       linestyle.Dashed,
	}
	return prop
}

// LineProp is responsible to give a valid properties.Line.
func LineProp() properties.Line {
	colorProp := ColorProp()
	prop := properties.Line{
		Color:         &colorProp,
		Style:         linestyle.Dashed,
		Thickness:     1.1,
		Orientation:   orientation.Vertical,
		OffsetPercent: 50,
		SizePercent:   20,
	}
	prop.MakeValid()
	return prop
}

// BarcodeProp is responsible to give a valid properties.Barcode.
func BarcodeProp() properties.Barcode {
	prop := properties.Barcode{
		Top:     10,
		Left:    10,
		Percent: 98,
		Proportion: properties.Proportion{
			Width:  16,
			Height: 9,
		},
		Center: false,
	}
	prop.MakeValid()
	return prop
}

// PageProp is responsible to give a valid properties.PageNumber.
func PageProp() properties.PageNumber {
	fontProp := FontProp()
	prop := properties.PageNumber{
		Pattern: "{current} / {total}",
		Place:   properties.LeftBottom,
		Family:  fontProp.Family,
		Style:   fontProp.Style,
		Size:    fontProp.Size,
		Color:   fontProp.Color,
	}
	return prop
}

// SignatureProp is responsible to give a valid properties.Signature.
func SignatureProp() properties.Signature {
	textProp := TextProp()
	lineProp := LineProp()
	prop := properties.Signature{
		FontFamily:    textProp.Family,
		FontStyle:     textProp.Style,
		FontSize:      textProp.Size,
		FontColor:     textProp.Color,
		LineColor:     lineProp.Color,
		LineStyle:     lineProp.Style,
		LineThickness: lineProp.Thickness,
	}
	prop.MakeValid(textProp.Family)
	return prop
}

// ConfigEntity is responsible to give a valid entity.Config.
func ConfigEntity() entity.Config {
	return entity.Config{
		Margins: &entity.Margins{
			Left:   10,
			Top:    10,
			Right:  10,
			Bottom: 10,
		},
	}
}

// CellEntity is responsible to give a valid entity.Cell.
func CellEntity() entity.Cell {
	return entity.Cell{
		X:      10,
		Y:      15,
		Width:  100,
		Height: 150,
	}
}

// MarginsEntity is responsible to give a valid entity.Margins.
func MarginsEntity() entity.Margins {
	return entity.Margins{
		Left:   10,
		Top:    10,
		Right:  10,
		Bottom: 10,
	}
}

// ImageEntity is responsible to give a valid entity.Image.
func ImageEntity() entity.Image {
	return entity.Image{
		Bytes:     []byte{1, 2, 3},
		Extension: extension.Png,
	}
}

// Node is responsible to give a valid node.Node.
func Node(rootType string) *node.Node[core.Structure] {
	bpdfNode := node.New[core.Structure](core.Structure{
		Type: rootType,
	})
	pageNode := node.New[core.Structure](core.Structure{
		Type: "page",
	})

	bpdfNode.AddNext(pageNode)
	return bpdfNode
}
