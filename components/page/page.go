// Page package implements creation of pages.
package page

import (
	"github.com/pchchv/bpdf/core"
	"github.com/pchchv/bpdf/core/entity"
	"github.com/pchchv/bpdf/properties"
)

type Page struct {
	number int
	total  int
	rows   []core.Row
	config *entity.Config
	prop   properties.PageNumber
}
