// nolint:dupl
package code

import (
	"github.com/pchchv/bpdf/core"
	"github.com/pchchv/bpdf/core/entity"
	"github.com/pchchv/bpdf/properties"
)

type MatrixCode struct {
	code   string
	prop   properties.Rect
	config *entity.Config
}

// SetConfig sets the configuration of a MatrixCode.
func (m *MatrixCode) SetConfig(config *entity.Config) {
	m.config = config
}

// Render renders a MatrixCode into a PDF context.
func (m *MatrixCode) Render(provider core.Provider, cell *entity.Cell) {
	provider.AddMatrixCode(m.code, cell, &m.prop)
}
