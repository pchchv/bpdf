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

func TestCellWriterBuilder_Build(t *testing.T) {
	sut := cellwriter.NewBuilder()

	chain := sut.Build(nil)

	assert.Equal(t, "borderThicknessStyler", chain.GetName())
	chain = chain.GetNext()
	assert.Equal(t, "borderLineStyler", chain.GetName())
	chain = chain.GetNext()
	assert.Equal(t, "borderColorStyler", chain.GetName())
	chain = chain.GetNext()
	assert.Equal(t, "fillColorStyler", chain.GetName())
	chain = chain.GetNext()
	assert.Equal(t, "cellWriter", chain.GetName())
	chain = chain.GetNext()
	assert.Nil(t, chain)
}
