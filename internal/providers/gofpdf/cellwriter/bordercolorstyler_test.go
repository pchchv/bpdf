package cellwriter_test

import (
	"fmt"
	"testing"

	"github.com/pchchv/bpdf/internal/providers/gofpdf/cellwriter"
	"github.com/stretchr/testify/assert"
)

func TestNewBorderColorStyler(t *testing.T) {
	sut := cellwriter.NewBorderColorStyler(nil)

	assert.NotNil(t, sut)
	assert.Equal(t, "*cellwriter.borderColorStyler", fmt.Sprintf("%T", sut))
}
