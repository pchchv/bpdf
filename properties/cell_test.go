package properties_test

import (
	"testing"

	"github.com/pchchv/bpdf/consts/border"
	"github.com/pchchv/bpdf/consts/linestyle"
	"github.com/pchchv/bpdf/internal/fixture"
	"github.com/pchchv/bpdf/properties"
	"github.com/stretchr/testify/assert"
)

func TestCell_ToMap(t *testing.T) {
	t.Run("when cell is nil, should return nil", func(t *testing.T) {
		var sut *properties.Cell
		m := sut.ToMap()

		assert.Nil(t, m)
	})

	t.Run("when cell is filled, should return map filled correctly", func(t *testing.T) {
		sut := fixture.CellProp()
		m := sut.ToMap()

		assert.Equal(t, border.Left, m["prop_border_type"])
		assert.Equal(t, 0.6, m["prop_border_thickness"])
		assert.Equal(t, linestyle.Dashed, m["prop_border_line_style"])
		assert.Equal(t, "RGB(255, 100, 50)", m["prop_background_color"])
		assert.Equal(t, "RGB(200, 80, 60)", m["prop_border_color"])
	})
}
