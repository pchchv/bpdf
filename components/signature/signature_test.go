package signature_test

import (
	"testing"

	"github.com/pchchv/bpdf/components/signature"
	"github.com/pchchv/bpdf/internal/fixture"
	"github.com/pchchv/bpdf/mocks"
	"github.com/pchchv/bpdf/properties"
	"github.com/pchchv/bpdf/test"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	t.Run("when prop is not sent, should use default", func(t *testing.T) {
		sut := signature.New("signature")

		test.New(t).Assert(sut.GetStructure()).Equals("components/signatures/new_signature_default_prop.json")
	})

	t.Run("when prop is sent, should use the provided", func(t *testing.T) {
		sut := signature.New("signature", fixture.SignatureProp())

		test.New(t).Assert(sut.GetStructure()).Equals("components/signatures/new_signature_custom_prop.json")
	})
}

func TestNewCol(t *testing.T) {
	t.Run("when prop is not sent, should use default", func(t *testing.T) {
		sut := signature.NewCol(12, "signature")

		test.New(t).Assert(sut.GetStructure()).Equals("components/signatures/new_signature_col_default_prop.json")
	})

	t.Run("when prop is sent, should use the provided", func(t *testing.T) {
		sut := signature.NewCol(12, "signature", fixture.SignatureProp())

		test.New(t).Assert(sut.GetStructure()).Equals("components/signatures/new_signature_col_custom_prop.json")
	})
}

func TestNewRow(t *testing.T) {
	t.Run("when prop is not sent, should use default", func(t *testing.T) {
		sut := signature.NewRow(10, "signature")

		test.New(t).Assert(sut.GetStructure()).Equals("components/signatures/new_signature_row_default_prop.json")
	})

	t.Run("when prop is sent, should use the provided", func(t *testing.T) {
		sut := signature.NewRow(10, "signature", fixture.SignatureProp())

		test.New(t).Assert(sut.GetStructure()).Equals("components/signatures/new_signature_row_custom_prop.json")
	})
}

func TestNewAutoRow(t *testing.T) {
	t.Run("when prop is not sent, should use default", func(t *testing.T) {
		sut := signature.NewAutoRow("signature")

		test.New(t).Assert(sut.GetStructure()).Equals("components/signatures/new_signature_auto_row_default_prop.json")
	})

	t.Run("when prop is sent, should use the provided", func(t *testing.T) {
		sut := signature.NewAutoRow("signature", fixture.SignatureProp())

		test.New(t).Assert(sut.GetStructure()).Equals("components/signatures/new_signature_auto_row_custom_prop.json")
	})
}

func TestSignature_GetHeight(t *testing.T) {
	t.Run("When signature has a height of 10, should return 10", func(t *testing.T) {
		cell := fixture.CellEntity()
		font := fixture.FontProp()
		sut := signature.New("signature",
			properties.Signature{
				SafePadding: 1,
				FontFamily:  font.Family,
				FontStyle:   font.Style,
				FontSize:    font.Size, FontColor: font.Color,
				LineThickness: 2,
			})
		provider := mocks.NewProvider(t)
		provider.EXPECT().GetFontHeight(&font).Return(5.0)

		height := sut.GetHeight(provider, &cell)
		assert.Equal(t, 7.0, height)
	})
}
