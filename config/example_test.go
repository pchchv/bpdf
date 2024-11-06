package config_test

import (
	"github.com/pchchv/bpdf"
	"github.com/pchchv/bpdf/config"
)

// ExampleNewBuilder demonstrates how to use builder.
func ExampleNewBuilder() {
	cfg := config.NewBuilder().Build()

	_ = bpdf.New(cfg)

	// generate document
}
