package cellwriter

import "github.com/pchchv/bpdf/internal/providers/gofpdf/fpdfwrapper"

type borderLineStyler struct {
	stylerTemplate
}

func NewBorderLineStyler(fpdf fpdfwrapper.Fpdf) *borderLineStyler {
	return &borderLineStyler{
		stylerTemplate: stylerTemplate{
			fpdf: fpdf,
			name: "borderLineStyler",
		},
	}
}
