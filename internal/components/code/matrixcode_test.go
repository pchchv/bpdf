// nolint: dupl
package code_test

import (
	"testing"

	"github.com/pchchv/bpdf/internal/components/code"
	"github.com/pchchv/bpdf/internal/fixture"
	"github.com/pchchv/bpdf/mocks"
)

func TestMatrixCode_Render(t *testing.T) {
	t.Run("should call provider correctly", func(t *testing.T) {
		codeValue := "code"
		cell := fixture.CellEntity()
		prop := fixture.RectProp()
		sut := code.NewMatrix(codeValue, prop)
		provider := mocks.NewProvider(t)
		provider.EXPECT().AddMatrixCode(codeValue, &cell, &prop)

		sut.Render(provider, &cell)

		provider.AssertNumberOfCalls(t, "AddMatrixCode", 1)
	})
}

func TestMatrixCode_SetConfig(t *testing.T) {
	t.Run("should call correctly", func(t *testing.T) {
		sut := code.NewMatrix("code")

		sut.SetConfig(nil)
	})
}
