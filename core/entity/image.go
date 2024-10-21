package entity

import "github.com/pchchv/bpdf/consts/extension"

// Image is the representation of an image that can be added to the pdf.
type Image struct {
	Bytes      []byte
	Extension  extension.Extension
	Dimensions *Dimensions
}
