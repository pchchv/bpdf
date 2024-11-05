// nolint: dupl
package code_test

import (
	"errors"
	"testing"

	"github.com/pchchv/bpdf/components/code"
	"github.com/pchchv/bpdf/core/entity"
	"github.com/pchchv/bpdf/internal/fixture"
	"github.com/pchchv/bpdf/mocks"
	"github.com/pchchv/bpdf/test"
	"github.com/stretchr/testify/assert"
)

func TestQrCode_Render(t *testing.T) {
	t.Run("should call provider correctly", func(t *testing.T) {
		codeValue := "code"
		cell := fixture.CellEntity()
		prop := fixture.RectProp()
		sut := code.NewQr(codeValue, prop)
		provider := mocks.NewProvider(t)
		provider.EXPECT().AddQrCode(codeValue, &cell, &prop)

		sut.Render(provider, &cell)

		provider.AssertNumberOfCalls(t, "AddQrCode", 1)
	})
}

func TestQrCode_SetConfig(t *testing.T) {
	t.Run("should call correctly", func(t *testing.T) {
		sut := code.NewQr("code")

		sut.SetConfig(nil)
	})
}

func TestQrCode_GetHeight(t *testing.T) {
	t.Run("When it is not possible to know the dimensions of the qrcode, should return height 0", func(t *testing.T) {
		cell := fixture.CellEntity()
		provider := mocks.NewProvider(t)
		provider.EXPECT().GetDimensionsByQrCode("code").Return(nil, errors.New("anyError2"))
		sut := code.NewQr("code")

		height := sut.GetHeight(provider, &cell)
		assert.Equal(t, height, 0.0)
	})

	t.Run("When the height of the qr code is half the width, should return half the width of the cell", func(t *testing.T) {
		cell := fixture.CellEntity()
		provider := mocks.NewProvider(t)
		provider.EXPECT().GetDimensionsByQrCode("code").Return(&entity.Dimensions{Width: 10, Height: 5}, nil)
		sut := code.NewQr("code")

		height := sut.GetHeight(provider, &cell)
		assert.Equal(t, height, cell.Width/2)
	})
}

func TestNewQr(t *testing.T) {
	t.Run("when prop is not sent, should use default", func(t *testing.T) {
		sut := code.NewQr("code")

		test.New(t).Assert(sut.GetStructure()).Equals("components/codes/new_qr_default_prop.json")
	})

	t.Run("when prop is sent, should use the provided", func(t *testing.T) {
		sut := code.NewQr("code", fixture.RectProp())

		test.New(t).Assert(sut.GetStructure()).Equals("components/codes/new_qr_custom_prop.json")
	})
}

func TestNewAutoQrRow(t *testing.T) {
	t.Run("when prop is not sent, should use default", func(t *testing.T) {
		sut := code.NewAutoQrRow("code")

		test.New(t).Assert(sut.GetStructure()).Equals("components/codes/new_auto_qr_row_default_prop.json")
	})

	t.Run("when prop is sent, should use the provided", func(t *testing.T) {
		sut := code.NewAutoQrRow("code", fixture.RectProp())

		test.New(t).Assert(sut.GetStructure()).Equals("components/codes/new_auto_qr_row_custom_prop.json")
	})
}

func TestNewQrCol(t *testing.T) {
	t.Run("when prop is not sent, should use default", func(t *testing.T) {
		sut := code.NewQrCol(12, "code")

		test.New(t).Assert(sut.GetStructure()).Equals("components/codes/new_qr_col_default_prop.json")
	})

	t.Run("when prop is sent, should use the provided", func(t *testing.T) {
		sut := code.NewQrCol(12, "code", fixture.RectProp())

		test.New(t).Assert(sut.GetStructure()).Equals("components/codes/new_qr_col_custom_prop.json")
	})
}

func TestNewQrRow(t *testing.T) {
	t.Run("when prop is not sent, should use default", func(t *testing.T) {
		sut := code.NewQrRow(10, "code")

		test.New(t).Assert(sut.GetStructure()).Equals("components/codes/new_qr_row_default_prop.json")
	})

	t.Run("when prop is sent, should use the provided", func(t *testing.T) {
		sut := code.NewQrRow(10, "code", fixture.RectProp())

		test.New(t).Assert(sut.GetStructure()).Equals("components/codes/new_qr_row_custom_prop.json")
	})
}
