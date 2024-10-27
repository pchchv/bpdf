package cellwriter

import "github.com/pchchv/bpdf/internal/providers/gofpdf/fpdfwrapper"

type stylerTemplate struct {
	next CellWriter
	fpdf fpdfwrapper.Fpdf
	name string
}
