package properties

import "github.com/pchchv/bpdf/consts/barcode"

// Barcode represents properties from a barcode inside a cell.
type Barcode struct {
	// Left is the space between the left cell boundary to the barcode, if center is false.
	Left float64
	// Top is space between the upper cell limit to the barcode, if center is false.
	Top float64
	// Percent is how much the barcode will occupy the cell,
	// ex 100%: The barcode will fulfill the entire cell
	// ex 50%: The greater side from the barcode will have half the size of the cell.
	Percent float64
	// Proportion is the proportion between size of the barcode.
	// Ex: 16x9, 4x3...
	Proportion Proportion
	// Center define that the barcode will be vertically and horizontally centralized.
	Center bool
	// Type represents the barcode type. Default: code128
	Type barcode.Code
}

// ToMap from Barcode will return a map representation from Barcode.
func (b *Barcode) ToMap() map[string]interface{} {
	if b == nil {
		return nil
	}

	m := make(map[string]interface{})
	if b.Left != 0 {
		m["prop_left"] = b.Left
	}

	if b.Top != 0 {
		m["prop_top"] = b.Top
	}

	if b.Percent != 0 {
		m["prop_percent"] = b.Percent
	}

	if b.Proportion.Width != 0 {
		m["prop_proportion_width"] = b.Proportion.Width
	}

	if b.Proportion.Height != 0 {
		m["prop_proportion_height"] = b.Proportion.Height
	}

	if b.Center {
		m["prop_center"] = b.Center
	}

	return m
}

// ToRectProp from Barcode will return a Rect representation from Barcode.
func (b *Barcode) ToRectProp() *Rect {
	return &Rect{
		Left:    b.Left,
		Top:     b.Top,
		Percent: b.Percent,
		Center:  b.Center,
	}
}
