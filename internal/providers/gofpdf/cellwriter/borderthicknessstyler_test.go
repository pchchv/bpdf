package cellwriter_test

import (
	"fmt"
	"testing"
	"github.com/pchchv/bpdf/internal/providers/gofpdf/cellwriter"
	"github.com/stretchr/testify/assert"
)

func TestNewBorderThicknessStyler(t *testing.T) {
	sut := cellwriter.NewBorderThicknessStyler(nil)

	assert.NotNil(t, sut)
	assert.Equal(t, "*cellwriter.borderThicknessStyler", fmt.Sprintf("%T", sut))
}
