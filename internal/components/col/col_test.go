package col_test

import (
	"testing"

	"github.com/pchchv/bpdf/core/entity"
	"github.com/pchchv/bpdf/internal/components/col"
	"github.com/pchchv/bpdf/internal/fixture"
	"github.com/pchchv/bpdf/mocks"
	"github.com/pchchv/bpdf/properties"
)

func TestCol_Render(t *testing.T) {
	t.Run("when not createCell, should call provider correctly", func(t *testing.T) {
		cfg := &entity.Config{}
		cell := fixture.CellEntity()
		style := &properties.Cell{}
		provider := mocks.NewProvider(t)
		component := mocks.NewComponent(t)
		component.EXPECT().Render(provider, &cell)
		component.EXPECT().SetConfig(cfg)
		sut := col.New(12).Add(component)
		sut.WithStyle(style)
		sut.SetConfig(cfg)

		sut.Render(provider, cell, false)

		component.AssertNumberOfCalls(t, "Render", 1)
		component.AssertNumberOfCalls(t, "SetConfig", 1)
	})

	t.Run("when createCell, should call provider correctly", func(t *testing.T) {
		cfg := &entity.Config{}
		cell := fixture.CellEntity()
		style := &properties.Cell{}
		provider := mocks.NewProvider(t)
		provider.EXPECT().CreateCol(cell.Width, cell.Height, cfg, style)
		component := mocks.NewComponent(t)
		component.EXPECT().Render(provider, &cell)
		component.EXPECT().SetConfig(cfg)
		sut := col.New(12).Add(component)
		sut.WithStyle(style)
		sut.SetConfig(cfg)

		sut.Render(provider, cell, true)

		provider.AssertNumberOfCalls(t, "CreateCol", 1)
		component.AssertNumberOfCalls(t, "Render", 1)
		component.AssertNumberOfCalls(t, "SetConfig", 1)
	})
}
