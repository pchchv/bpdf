// nolint: dupl
package code_test

import (
	"testing"

	"github.com/pchchv/bpdf/internal/components/code"
	"github.com/pchchv/bpdf/internal/fixture"
	"github.com/pchchv/bpdf/mocks"
	"github.com/pchchv/bpdf/properties"
	"github.com/pchchv/bpdf/test"
	"github.com/stretchr/testify/assert"
)

func TestBarcode_Render(t *testing.T) {
	t.Run("should call provider correctly", func(t *testing.T) {
		codeValue := "barcode"
		cell := fixture.CellEntity()
		prop := fixture.BarcodeProp()
		sut := code.NewBar(codeValue, prop)
		provider := mocks.NewProvider(t)
		provider.EXPECT().AddBarCode(codeValue, &cell, &prop)

		sut.Render(provider, &cell)

		provider.AssertNumberOfCalls(t, "AddBarCode", 1)
	})
}

func TestBarcode_SetConfig(t *testing.T) {
	t.Run("should call correctly", func(t *testing.T) {
		sut := code.NewBar("code")

		sut.SetConfig(nil)
	})
}

func TestBarcode_GetHeight(t *testing.T) {
	t.Run("When the barcode height is '20%' of the width, it should return '20%' of the cell width", func(t *testing.T) {
		cell := fixture.CellEntity()
		provider := mocks.NewProvider(t)
		sut := code.NewBar("code", properties.Barcode{Proportion: properties.Proportion{Width: 10.0, Height: 2.0}, Percent: 100.0})

		height := sut.GetHeight(provider, &cell)
		assert.Equal(t, height, cell.Width*0.2)
	})
}

func TestNewBar(t *testing.T) {
	t.Run("when prop is not sent, should use default", func(t *testing.T) {
		sut := code.NewBar("code")

		test.New(t).Assert(sut.GetStructure()).Equals("components/codes/new_bar_default_prop.json")
	})

	t.Run("when prop is sent, should use the provided", func(t *testing.T) {
		sut := code.NewBar("code", fixture.BarcodeProp())

		test.New(t).Assert(sut.GetStructure()).Equals("components/codes/new_bar_custom_prop.json")
	})
}

func TestNewBarCol(t *testing.T) {
	t.Run("when prop is not sent, should use default", func(t *testing.T) {
		sut := code.NewBarCol(12, "code")

		test.New(t).Assert(sut.GetStructure()).Equals("components/codes/new_bar_col_default_prop.json")
	})

	t.Run("when prop is sent, should use the provided", func(t *testing.T) {
		sut := code.NewBarCol(12, "code", fixture.BarcodeProp())

		test.New(t).Assert(sut.GetStructure()).Equals("components/codes/new_bar_col_custom_prop.json")
	})
}

func TestNewBarRow(t *testing.T) {
	t.Run("when prop is not sent, should use default", func(t *testing.T) {
		sut := code.NewBarRow(10, "code")

		test.New(t).Assert(sut.GetStructure()).Equals("components/codes/new_bar_row_default_prop.json")
	})

	t.Run("when prop is sent, should use the provided", func(t *testing.T) {
		sut := code.NewBarRow(10, "code", fixture.BarcodeProp())

		test.New(t).Assert(sut.GetStructure()).Equals("components/codes/new_bar_row_custom_prop.json")
	})
}
