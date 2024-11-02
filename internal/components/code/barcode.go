// Package code implements creation of Barcode, MatrixCode and QrCode.
// nolint:dupl
// It's similar to Barcode.go and it's hard to extract common code.
package code

import (
	"github.com/pchchv/bpdf/core/entity"
	"github.com/pchchv/bpdf/properties"
)

type Barcode struct {
	code   string
	prop   properties.Barcode
	config *entity.Config
}
