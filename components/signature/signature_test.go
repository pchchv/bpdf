package signature_test

import (
	"testing"

	"github.com/pchchv/bpdf/components/signature"
	"github.com/pchchv/bpdf/internal/fixture"
	"github.com/pchchv/bpdf/test"
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
