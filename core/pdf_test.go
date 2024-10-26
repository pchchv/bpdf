package core_test

import (
	"fmt"
	"os"
	"path"
	"strings"
	"testing"

	"github.com/pchchv/bpdf/core"
	"github.com/pchchv/bpdf/metrics"
	"github.com/stretchr/testify/assert"
)

func TestNewPDF(t *testing.T) {
	sut := core.NewPDF(nil, nil)

	assert.NotNil(t, sut)
	assert.Equal(t, "*core.Pdf", fmt.Sprintf("%T", sut))
}

func TestPdf_GetBase64(t *testing.T) {
	sut := core.NewPDF([]byte{1, 2, 3}, nil)
	b64 := sut.GetBase64()

	assert.Equal(t, "AQID", b64)
}

func TestPdf_GetBytes(t *testing.T) {
	sut := core.NewPDF([]byte{1, 2, 3}, nil)
	bytes := sut.GetBytes()

	assert.Equal(t, []byte{1, 2, 3}, bytes)
}

func TestPdf_GetReport(t *testing.T) {
	sut := core.NewPDF(nil, &metrics.Report{SizeMetric: metrics.SizeMetric{
		Key: "key",
		Size: metrics.Size{
			Value: 10.0,
			Scale: metrics.Byte,
		},
	}})
	report := sut.GetReport()

	assert.Equal(t, "key", report.SizeMetric.Key)
}

func buildPath(file string) (dir string) {
	var err error
	if dir, err = os.Getwd(); err != nil {
		return ""
	}

	dir = strings.ReplaceAll(dir, "pkg/core/entity", "")
	return path.Join(dir, file)
}
