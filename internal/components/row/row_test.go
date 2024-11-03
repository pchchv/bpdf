package row_test

import (
	"testing"

	"github.com/pchchv/bpdf/core"
	"github.com/pchchv/bpdf/internal/components/col"
	"github.com/pchchv/bpdf/internal/components/row"
	"github.com/pchchv/bpdf/internal/fixture"
	"github.com/pchchv/bpdf/mocks"
	"github.com/stretchr/testify/assert"
)

func TestRow_GetHeight(t *testing.T) {
	t.Run("When a row has a column with height 5, should return 5", func(t *testing.T) {
		cell := fixture.CellEntity()
		provider := mocks.NewProvider(t)
		columns := mocks.NewCol(t)
		columns.EXPECT().GetHeight(provider, &cell).Return(5)

		r := row.New().Add(columns)

		assert.Equal(t, 5.0, r.GetHeight(provider, &cell))
	})
}

func TestRow_GetColumns(t *testing.T) {
	t.Run("when GetColumns is called, should return the number of registered columns", func(t *testing.T) {
		newCol := []core.Col{col.New(12)}
		r := row.New(10).Add(newCol[0])

		assert.Equal(t, newCol, r.GetColumns())
	})
}
