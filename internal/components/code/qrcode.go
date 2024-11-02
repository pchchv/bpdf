// nolint:dupl
package code

import (
	"github.com/pchchv/bpdf/core/entity"
	"github.com/pchchv/bpdf/properties"
)

type QrCode struct {
	code   string
	prop   properties.Rect
	config *entity.Config
}
