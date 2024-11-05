// nolint: dupl
package code_test

import (
	"errors"
	"testing"

	"github.com/pchchv/bpdf/core/entity"
	"github.com/pchchv/bpdf/internal/components/code"
	"github.com/pchchv/bpdf/internal/fixture"
	"github.com/pchchv/bpdf/mocks"
	"github.com/pchchv/bpdf/test"
	"github.com/stretchr/testify/assert"
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

func TestMatrixCode_GetHeight(t *testing.T) {
	t.Run("When it is not possible to know the dimensions of the matrix code, should return height 0", func(t *testing.T) {
		cell := fixture.CellEntity()
		provider := mocks.NewProvider(t)
		provider.EXPECT().GetDimensionsByMatrixCode("code").Return(nil, errors.New("anyError2"))
		sut := code.NewMatrix("code")

		height := sut.GetHeight(provider, &cell)
		assert.Equal(t, height, 0.0)
	})

	t.Run("When the height of the matrix code is half the width, should return half the width of the cell", func(t *testing.T) {
		cell := fixture.CellEntity()
		provider := mocks.NewProvider(t)
		provider.EXPECT().GetDimensionsByMatrixCode("code").Return(&entity.Dimensions{Width: 10, Height: 5}, nil)
		sut := code.NewMatrix("code")

		height := sut.GetHeight(provider, &cell)
		assert.Equal(t, height, cell.Width/2)
	})
}

func TestNewMatrix(t *testing.T) {
	t.Run("when prop is not sent, should use default", func(t *testing.T) {
		sut := code.NewMatrix("code")

		test.New(t).Assert(sut.GetStructure()).Equals("components/codes/new_matrix_default_prop.json")
	})

	t.Run("when prop is sent, should use the provided", func(t *testing.T) {
		sut := code.NewMatrix("code", fixture.RectProp())

		test.New(t).Assert(sut.GetStructure()).Equals("components/codes/new_matrix_custom_prop.json")
	})
}

func TestAutoNewMatrixRow(t *testing.T) {
	t.Run("when prop is not sent, should use default", func(t *testing.T) {
		sut := code.NewAutoMatrixRow("code")

		test.New(t).Assert(sut.GetStructure()).Equals("components/codes/new_auto_matrix_row_default_prop.json")
	})

	t.Run("when prop is sent, should use the provided", func(t *testing.T) {
		sut := code.NewAutoMatrixRow("code", fixture.RectProp())

		test.New(t).Assert(sut.GetStructure()).Equals("components/codes/new_auto_matrix_row_custom_prop.json")
	})
}
