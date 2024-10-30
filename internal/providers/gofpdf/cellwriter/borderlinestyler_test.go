package cellwriter_test

import (
	"fmt"
	"testing"

	"github.com/pchchv/bpdf/internal/providers/gofpdf/cellwriter"
	"github.com/stretchr/testify/assert"
)

func TestNewBorderLineStyler(t *testing.T) {
	sut := cellwriter.NewBorderLineStyler(nil)

	assert.NotNil(t, sut)
	assert.Equal(t, "*cellwriter.borderLineStyler", fmt.Sprintf("%T", sut))
}
