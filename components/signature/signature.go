// Package signature implements creation of signatures.
package signature

import (
	"github.com/pchchv/bpdf/core/entity"
	"github.com/pchchv/bpdf/properties"
)

type Signature struct {
	value  string
	prop   properties.Signature
	config *entity.Config
}
