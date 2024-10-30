package cellwriter_test

import (
	"fmt"
	"testing"

	"github.com/pchchv/bpdf/internal/providers/gofpdf/cellwriter"
	"github.com/stretchr/testify/assert"
)

func TestNewBuilder(t *testing.T) {
	sut := cellwriter.NewBuilder()

	assert.NotNil(t, sut)
	assert.Equal(t, "*cellwriter.CellWriterBuilder", fmt.Sprintf("%T", sut))
}
