// nolint: dupl
package code_test

import (
	"testing"

	"github.com/pchchv/bpdf/internal/components/code"
	"github.com/pchchv/bpdf/internal/fixture"
	"github.com/pchchv/bpdf/mocks"
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
