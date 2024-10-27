package cellwriter

import (
	"github.com/pchchv/bpdf/core/entity"
	"github.com/pchchv/bpdf/internal/providers/gofpdf/fpdfwrapper"
	"github.com/pchchv/bpdf/properties"
)

type stylerTemplate struct {
	next CellWriter
	fpdf fpdfwrapper.Fpdf
	name string
}

func (s *stylerTemplate) GetName() string {
	return s.name
}

func (s *stylerTemplate) GetNext() CellWriter {
	return s.next
}

func (s *stylerTemplate) SetNext(next CellWriter) {
	s.next = next
}

func (s *stylerTemplate) GoToNext(width, height float64, config *entity.Config, prop *properties.Cell) {
	if s.next == nil {
		return
	}
	s.next.Apply(width, height, config, prop)
}
