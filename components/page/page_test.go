package page_test

import (
	"testing"

	"github.com/pchchv/bpdf/components/image"
	"github.com/pchchv/bpdf/components/page"
	"github.com/pchchv/bpdf/consts/extension"
	"github.com/pchchv/bpdf/core"
	"github.com/pchchv/bpdf/core/entity"
	"github.com/pchchv/bpdf/internal/fixture"
	"github.com/pchchv/bpdf/mocks"
	"github.com/pchchv/bpdf/properties"
	"github.com/pchchv/bpdf/test"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	t.Run("when prop is not sent, should use default", func(t *testing.T) {
		sut := page.New()

		test.New(t).Assert(sut.GetStructure()).Equals("components/lines/new_page_default_prop.json")
	})

	t.Run("when prop is sent, should use the provided", func(t *testing.T) {
		sut := page.New(fixture.PageProp())

		test.New(t).Assert(sut.GetStructure()).Equals("components/lines/new_page_custom_prop.json")
	})

	t.Run("when prop is sent and there is rows, should use the provided", func(t *testing.T) {
		sut := page.New(fixture.PageProp())
		row := image.NewFromFileRow(10, "path")
		sut.Add(row)

		test.New(t).Assert(sut.GetStructure()).Equals("components/lines/new_page_custom_prop_and_with_rows.json")
	})
}

func TestPage_GetRows(t *testing.T) {
	t.Run("when called get rows, should return rows correctly", func(t *testing.T) {
		row := mocks.NewRow(t)
		sut := page.New()
		sut.Add(row)

		rows := sut.GetRows()

		assert.Equal(t, []core.Row{row}, rows)
	})
}

func TestPage_SetNumber(t *testing.T) {
	t.Run("when called set number, should set correctly", func(t *testing.T) {
		sut := page.New()

		sut.SetNumber(1, 2)

		assert.Equal(t, 1, sut.GetNumber())
	})
}

func TestPage_Render(t *testing.T) {
	t.Run("when there is no background image and there is no page pattern, should call row render correctly", func(t *testing.T) {
		cell := fixture.CellEntity()
		prop := fixture.PageProp()
		prop.Pattern = ""
		cfg := &entity.Config{}
		provider := mocks.NewProvider(t)
		row := mocks.NewRow(t)
		row.EXPECT().Render(provider, cell)
		row.EXPECT().GetHeight(provider, &cell).Return(10.0)
		row.EXPECT().SetConfig(cfg)
		sut := page.New(prop)
		sut.Add(row)
		sut.SetConfig(cfg)

		sut.Render(provider, cell)

		row.AssertNumberOfCalls(t, "Render", 1)
		row.AssertNumberOfCalls(t, "GetHeight", 1)
	})

	t.Run("when there is background image and there is no page pattern, should call row render and provider correctly", func(t *testing.T) {
		cell := fixture.CellEntity()
		prop := fixture.PageProp()
		prop.Pattern = ""
		cfg := &entity.Config{
			BackgroundImage: &entity.Image{
				Bytes:     []byte{1, 2, 3},
				Extension: extension.Jpg,
			},
		}
		rectProp := &properties.Rect{}
		rectProp.MakeValid()
		provider := mocks.NewProvider(t)
		provider.EXPECT().AddBackgroundImageFromBytes(cfg.BackgroundImage.Bytes, &cell, rectProp, cfg.BackgroundImage.Extension)
		row := mocks.NewRow(t)
		row.EXPECT().Render(provider, cell)
		row.EXPECT().GetHeight(provider, &cell).Return(10.0)
		row.EXPECT().SetConfig(cfg)
		sut := page.New(prop)
		sut.Add(row)
		sut.SetConfig(cfg)

		sut.Render(provider, cell)

		provider.AssertNumberOfCalls(t, "AddBackgroundImageFromBytes", 1)
		row.AssertNumberOfCalls(t, "Render", 1)
		row.AssertNumberOfCalls(t, "GetHeight", 1)
	})

	t.Run("when there is background image and there is page pattern, should call row render and provider correctly", func(t *testing.T) {
		cell := fixture.CellEntity()
		prop := fixture.PageProp()
		cfg := &entity.Config{
			BackgroundImage: &entity.Image{
				Bytes:     []byte{1, 2, 3},
				Extension: extension.Jpg,
			},
		}
		rectProp := &properties.Rect{}
		rectProp.MakeValid()
		provider := mocks.NewProvider(t)
		provider.EXPECT().AddBackgroundImageFromBytes(cfg.BackgroundImage.Bytes, &cell, rectProp, cfg.BackgroundImage.Extension)
		provider.EXPECT().AddText("0 / 0", &cell, prop.GetNumberTextProp(cell.Height))
		row := mocks.NewRow(t)
		row.EXPECT().Render(provider, cell)
		row.EXPECT().GetHeight(provider, &cell).Return(10.0)
		row.EXPECT().SetConfig(cfg)
		sut := page.New(prop)
		sut.Add(row)
		sut.SetConfig(cfg)

		sut.Render(provider, cell)

		provider.AssertNumberOfCalls(t, "AddBackgroundImageFromBytes", 1)
		provider.AssertNumberOfCalls(t, "AddText", 1)
		row.AssertNumberOfCalls(t, "Render", 1)
		row.AssertNumberOfCalls(t, "GetHeight", 1)
	})
}
