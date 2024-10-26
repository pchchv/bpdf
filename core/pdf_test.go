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
