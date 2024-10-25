package core

import (
	"encoding/base64"
	"os"

	"github.com/pchchv/bpdf/internal/time"
	"github.com/pchchv/bpdf/merge"
	"github.com/pchchv/bpdf/metrics"
)

type Pdf struct {
	bytes  []byte
	report *metrics.Report
}

// GetBytes returns the PDF bytes.
func (p *Pdf) GetBytes() []byte {
	return p.bytes
}

// GetBase64 returns the PDF bytes in base64.
func (p *Pdf) GetBase64() string {
	return base64.StdEncoding.EncodeToString(p.bytes)
}

// GetReport returns the metrics.Report.
func (p *Pdf) GetReport() *metrics.Report {
	return p.report
}

// Save saves the PDF in a file.
func (p *Pdf) Save(file string) error {
	return os.WriteFile(file, p.bytes, os.ModePerm)
}

// Merge merges the PDF with another PDF.
func (p *Pdf) Merge(bytes []byte) (err error) {
	var mergedBytes []byte
	timeSpent := time.GetTimeSpent(func() {
		mergedBytes, err = merge.Bytes(p.bytes, bytes)
	})
	if err != nil {
		return err
	}

	p.bytes = mergedBytes
	if p.report != nil {
		p.appendMetric(timeSpent)
	}

	return nil
}

func (p *Pdf) appendMetric(timeSpent *metrics.Time) {
	timeMetric := metrics.TimeMetric{
		Key:   "merge_pdf",
		Times: []*metrics.Time{timeSpent},
		Avg:   timeSpent,
	}
	timeMetric.Normalize()
	p.report.TimeMetrics = append(p.report.TimeMetrics, timeMetric)
	p.report.SizeMetric = metrics.SizeMetric{
		Key: "file_size",
		Size: metrics.Size{
			Value: float64(len(p.bytes)),
			Scale: metrics.Byte,
		},
	}
	p.report.Normalize()
}
