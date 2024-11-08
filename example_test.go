package bpdf_test

import (
	"github.com/pchchv/bpdf"
	"github.com/pchchv/bpdf/config"
)

// ExampleNew demonstrates how to create a bpdf instance.
func ExampleNew() {
	// optional
	b := config.NewBuilder()
	cfg := b.Build()

	m := bpdf.New(cfg) // cfg is an optional

	// Do things and generate
	_, _ = m.Generate()
}
