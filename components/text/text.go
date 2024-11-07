// Package text implements creation of texts.
package text

import (
	"github.com/pchchv/bpdf/core/entity"
	"github.com/pchchv/bpdf/properties"
)

type Text struct {
	value  string
	prop   properties.Text
	config *entity.Config
}
