package row_test

import (
	"testing"

	"github.com/pchchv/bpdf/components/col"
	"github.com/pchchv/bpdf/components/row"
	"github.com/pchchv/bpdf/core"
	"github.com/pchchv/bpdf/core/entity"
	"github.com/pchchv/bpdf/internal/fixture"
	"github.com/pchchv/bpdf/mocks"
	"github.com/pchchv/bpdf/test"
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

func TestRow_GetStructure(t *testing.T) {
	t.Run("when there is no style, should call provider correctly", func(t *testing.T) {
		cfg := &entity.Config{
			MaxGridSize: 12,
		}
		cell := fixture.CellEntity()
		provider := mocks.NewProvider(t)
		provider.EXPECT().CreateRow(cell.Height)
		col := mocks.NewCol(t)
		col.EXPECT().Render(provider, cell, true)
		col.EXPECT().SetConfig(cfg)
		col.EXPECT().GetSize().Return(12)
		sut := row.New(cell.Height).Add(col)
		sut.SetConfig(cfg)

		sut.Render(provider, cell)

		provider.AssertNumberOfCalls(t, "CreateRow", 1)
		col.AssertNumberOfCalls(t, "Render", 1)
		col.AssertNumberOfCalls(t, "SetConfig", 1)
	})

	t.Run("when there is style, should call provider correctly", func(t *testing.T) {
		cfg := &entity.Config{
			MaxGridSize: 12,
		}
		cell := fixture.CellEntity()
		prop := fixture.CellProp()
		provider := mocks.NewProvider(t)
		provider.EXPECT().CreateRow(cell.Height)
		provider.EXPECT().CreateCol(cell.Width, cell.Height, cfg, &prop)
		col := mocks.NewCol(t)
		col.EXPECT().Render(provider, cell, false)
		col.EXPECT().SetConfig(cfg)
		col.EXPECT().GetSize().Return(12)
		sut := row.New(cell.Height).Add(col).WithStyle(&prop)
		sut.SetConfig(cfg)

		sut.Render(provider, cell)

		provider.AssertNumberOfCalls(t, "CreateCol", 1)
		provider.AssertNumberOfCalls(t, "CreateRow", 1)
		col.AssertNumberOfCalls(t, "Render", 1)
		col.AssertNumberOfCalls(t, "SetConfig", 1)
	})
}

func TestRow_SetConfig(t *testing.T) {
	t.Run("should call correctly", func(t *testing.T) {
		sut := row.New(10)

		sut.SetConfig(nil)
	})
}

func TestNew(t *testing.T) {
	t.Run("when there is no cols", func(t *testing.T) {
		r := row.New(10)

		test.New(t).Assert(r.GetStructure()).Equals("components/rows/new_empty_col.json")
	})

	t.Run("when has component, should retrieve components", func(t *testing.T) {
		r := row.New(12).Add(col.New(12))

		test.New(t).Assert(r.GetStructure()).Equals("components/rows/new_filled_col.json")
	})

	t.Run("when has prop, should apply correctly", func(t *testing.T) {
		prop := fixture.CellProp()
		r := row.New(12).WithStyle(&prop)

		test.New(t).Assert(r.GetStructure()).Equals("components/rows/new_col_with_prop.json")
	})
}
