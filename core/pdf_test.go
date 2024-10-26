package core_test

import (
	"fmt"
	"testing"

	"github.com/pchchv/bpdf/core"
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
