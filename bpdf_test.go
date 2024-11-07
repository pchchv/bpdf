package bpdf_test

import (
	"fmt"
	"testing"

	"github.com/pchchv/bpdf"
	"github.com/pchchv/bpdf/components/code"
	"github.com/pchchv/bpdf/components/col"
	"github.com/pchchv/bpdf/components/page"
	"github.com/pchchv/bpdf/components/row"
	"github.com/pchchv/bpdf/components/text"
	"github.com/pchchv/bpdf/config"
	"github.com/pchchv/bpdf/core"
	"github.com/pchchv/bpdf/test"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	t.Run("new default", func(t *testing.T) {
		sut := bpdf.New()

		assert.NotNil(t, sut)
		assert.Equal(t, "*bpdf.Bpdf", fmt.Sprintf("%T", sut))
	})

	t.Run("new with config", func(t *testing.T) {
		cfg := config.NewBuilder().
			Build()

		sut := bpdf.New(cfg)

		assert.NotNil(t, sut)
		assert.Equal(t, "*bpdf.Bpdf", fmt.Sprintf("%T", sut))
	})

	t.Run("new with config an concurrent mode on", func(t *testing.T) {
		cfg := config.NewBuilder().
			WithConcurrentMode(7).
			Build()

		sut := bpdf.New(cfg)

		assert.NotNil(t, sut)
		assert.Equal(t, "*bpdf.Bpdf", fmt.Sprintf("%T", sut))
	})

	t.Run("new with config an low memory mode on", func(t *testing.T) {
		cfg := config.NewBuilder().
			WithSequentialLowMemoryMode(10).
			Build()

		sut := bpdf.New(cfg)

		assert.NotNil(t, sut)
		assert.Equal(t, "*bpdf.Bpdf", fmt.Sprintf("%T", sut))
	})
}

func TestBPDF_AddRow(t *testing.T) {
	t.Run("when col is not sent, should empty col is set", func(t *testing.T) {
		sut := bpdf.New()
		sut.AddRow(10)

		test.New(t).Assert(sut.GetStructure()).Equals("bpdf_add_row_4.json")
	})

	t.Run("add one row", func(t *testing.T) {
		sut := bpdf.New()

		sut.AddRow(10, col.New(12))

		test.New(t).Assert(sut.GetStructure()).Equals("bpdf_add_row_1.json")
	})

	t.Run("add one row", func(t *testing.T) {
		sut := bpdf.New()

		sut.AddRow(10, col.New(12))

		test.New(t).Assert(sut.GetStructure()).Equals("bpdf_add_row_1.json")
	})

	t.Run("add two rows", func(t *testing.T) {
		sut := bpdf.New()

		sut.AddRow(10, col.New(12))
		sut.AddRow(10, col.New(12))

		test.New(t).Assert(sut.GetStructure()).Equals("bpdf_add_row_2.json")
	})

	t.Run("add rows until add new page", func(t *testing.T) {
		sut := bpdf.New()

		for i := 0; i < 30; i++ {
			sut.AddRow(10, col.New(12))
		}

		test.New(t).Assert(sut.GetStructure()).Equals("bpdf_add_row_3.json")
	})
}

func TestBPDF_AddPages(t *testing.T) {
	t.Run("add one page", func(t *testing.T) {
		sut := bpdf.New()

		sut.AddPages(
			page.New().Add(
				row.New(20).Add(col.New(12)),
			),
		)

		test.New(t).Assert(sut.GetStructure()).Equals("bpdf_add_pages_1.json")
	})
	t.Run("add two pages", func(t *testing.T) {
		sut := bpdf.New()

		sut.AddPages(
			page.New().Add(
				row.New(20).Add(col.New(12)),
			),
			page.New().Add(
				row.New(20).Add(col.New(12)),
			),
		)

		test.New(t).Assert(sut.GetStructure()).Equals("bpdf_add_pages_2.json")
	})

	t.Run("add page greater than one page", func(t *testing.T) {
		var rows []core.Row
		sut := bpdf.New()
		for i := 0; i < 15; i++ {
			rows = append(rows, row.New(20).Add(col.New(12)))
		}

		sut.AddPages(page.New().Add(rows...))

		test.New(t).Assert(sut.GetStructure()).Equals("bpdf_add_pages_3.json")
	})
}

func TestBPDF_Generate(t *testing.T) {
	t.Run("add one row", func(t *testing.T) {
		sut := bpdf.New()

		sut.AddRow(10, col.New(12))

		doc, err := sut.Generate()
		assert.Nil(t, err)
		assert.NotNil(t, doc)
	})

	t.Run("add two rows", func(t *testing.T) {
		sut := bpdf.New()

		sut.AddRow(10, col.New(12))
		sut.AddRow(10, col.New(12))

		doc, err := sut.Generate()
		assert.Nil(t, err)
		assert.NotNil(t, doc)
	})

	t.Run("add rows until add new page", func(t *testing.T) {
		sut := bpdf.New()

		for i := 0; i < 30; i++ {
			sut.AddRow(10, col.New(12))
		}

		doc, err := sut.Generate()
		assert.Nil(t, err)
		assert.NotNil(t, doc)
	})

	t.Run("add rows until add new page, execute in parallel", func(t *testing.T) {
		cfg := config.NewBuilder().
			WithConcurrentMode(7).
			Build()
		sut := bpdf.New(cfg)

		for i := 0; i < 30; i++ {
			sut.AddRow(10, col.New(12))
		}

		doc, err := sut.Generate()
		assert.Nil(t, err)
		assert.NotNil(t, doc)
	})

	t.Run("add rows until add new page, execute in low memory mode", func(t *testing.T) {
		cfg := config.NewBuilder().
			WithSequentialLowMemoryMode(10).
			Build()
		sut := bpdf.New(cfg)

		for i := 0; i < 30; i++ {
			sut.AddRow(10, col.New(12))
		}

		doc, err := sut.Generate()
		assert.Nil(t, err)
		assert.NotNil(t, doc)
	})

	t.Run("sequential generation", func(t *testing.T) {
		cfg := config.NewBuilder().
			WithSequentialMode().
			Build()
		sut := bpdf.New(cfg)

		for i := 0; i < 30; i++ {
			sut.AddRow(10, col.New(12))
		}

		test.New(t).Assert(sut.GetStructure()).Equals("bpdf_sequential.json")
	})

	t.Run("sequential low memory generation", func(t *testing.T) {
		cfg := config.NewBuilder().
			WithSequentialLowMemoryMode(10).
			Build()
		sut := bpdf.New(cfg)

		for i := 0; i < 30; i++ {
			sut.AddRow(10, col.New(12))
		}

		test.New(t).Assert(sut.GetStructure()).Equals("bpdf_sequential_low_memory.json")
	})

	t.Run("sequential low memory generation", func(t *testing.T) {
		cfg := config.NewBuilder().
			WithConcurrentMode(10).
			Build()
		sut := bpdf.New(cfg)

		for i := 0; i < 30; i++ {
			sut.AddRow(10, col.New(12))
		}

		test.New(t).Assert(sut.GetStructure()).Equals("bpdf_concurrent.json")
	})

	t.Run("page number", func(t *testing.T) {
		cfg := config.NewBuilder().
			WithPageNumber().
			Build()
		sut := bpdf.New(cfg)

		for i := 0; i < 30; i++ {
			sut.AddRow(10, col.New(12))
		}

		test.New(t).Assert(sut.GetStructure()).Equals("bpdf_page_number.json")
	})
}

func TestBPDF_GetCurrentConfig(t *testing.T) {
	t.Run("When GetCurrentConfig is called then current settings are returned", func(t *testing.T) {
		sut := bpdf.New(config.NewBuilder().
			WithMaxGridSize(20).
			Build())

		assert.Equal(t, sut.GetCurrentConfig().MaxGridSize, 20)
	})
}

// nolint:dupl
func TestBPDF_RegisterFooter(t *testing.T) {
	t.Run("when footer size is greater than useful area, should return error", func(t *testing.T) {
		sut := bpdf.New()
		err := sut.RegisterFooter(row.New(1000))
		assert.NotNil(t, err)
		assert.Equal(t, "footer height is greater than page useful area", err.Error())
	})

	t.Run("when header size is correct, should not return error and apply header", func(t *testing.T) {
		var rows []core.Row
		sut := bpdf.New()
		err := sut.RegisterFooter(code.NewBarRow(10, "footer"))
		for i := 0; i < 5; i++ {
			rows = append(rows, row.New(100).Add(col.New(12)))
		}

		sut.AddRows(rows...)

		assert.Nil(t, err)
		test.New(t).Assert(sut.GetStructure()).Equals("footer.json")
	})

	t.Run("when autoRow is sent, should set autoRow", func(t *testing.T) {
		var rows []core.Row
		sut := bpdf.New()
		err := sut.RegisterFooter(text.NewAutoRow("header"))
		for i := 0; i < 5; i++ {
			rows = append(rows, row.New(100).Add(col.New(12)))
		}

		sut.AddRows(rows...)

		assert.Nil(t, err)
		test.New(t).Assert(sut.GetStructure()).Equals("footer_auto_row.json")
	})
}

func TestBPDF_AddAutoRow(t *testing.T) {
	t.Run("When 100 automatic rows are sent, it should create 2 pages", func(t *testing.T) {
		sut := bpdf.New()

		for i := 0; i < 150; i++ {
			sut.AddAutoRow(text.NewCol(12, "teste"))
		}

		test.New(t).Assert(sut.GetStructure()).Equals("bpdf_add_auto_row_1.json")
	})
}
