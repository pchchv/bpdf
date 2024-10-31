package cellwriter_test

import (
	"fmt"
	"testing"

	"github.com/pchchv/bpdf/internal/providers/gofpdf/cellwriter"
	"github.com/stretchr/testify/assert"
)

func TestNewCellCreator(t *testing.T) {
	sut := cellwriter.NewCellWriter(nil)

	assert.NotNil(t, sut)
	assert.Equal(t, "*cellwriter.cellWriter", fmt.Sprintf("%T", sut))
}
