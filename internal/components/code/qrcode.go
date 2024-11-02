// nolint:dupl
package code

import (
	"github.com/pchchv/bpdf/core"
	"github.com/pchchv/bpdf/core/entity"
	"github.com/pchchv/bpdf/properties"
)

type QrCode struct {
	code   string
	prop   properties.Rect
	config *entity.Config
}

// Render renders a QrCode into a PDF context.
func (q *QrCode) Render(provider core.Provider, cell *entity.Cell) {
	provider.AddQrCode(q.code, cell, &q.prop)
}

// SetConfig set the config for the component.
func (q *QrCode) SetConfig(config *entity.Config) {
	q.config = config
}
