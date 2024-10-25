package core

import "github.com/pchchv/bpdf/metrics"

type Pdf struct {
	bytes  []byte
	report *metrics.Report
}
