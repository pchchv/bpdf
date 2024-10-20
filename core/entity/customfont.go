package entity

import "github.com/pchchv/bpdf/consts/fontstyle"

// CustomFont representats a font that can be added to the pdf.
type CustomFont struct {
	Family string
	Style  fontstyle.Fontstyle
	File   string
	Bytes  []byte
}
