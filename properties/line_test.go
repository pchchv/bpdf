package properties_test

import (
	"testing"

	"github.com/pchchv/bpdf/consts/linestyle"
	"github.com/pchchv/bpdf/consts/orientation"
	"github.com/pchchv/bpdf/internal/fixture"
	"github.com/pchchv/bpdf/properties"
	"github.com/stretchr/testify/assert"
)

func TestLine_MakeValid(t *testing.T) {
	t.Run("when style is empty, should apply solid", func(t *testing.T) {
		prop := properties.Line{
			Style: "",
		}
		prop.MakeValid()

		assert.Equal(t, linestyle.Solid, prop.Style)
	})

	t.Run("when thickness is 0.0, should apply default", func(t *testing.T) {
		prop := properties.Line{
			Thickness: 0.0,
		}
		prop.MakeValid()

		assert.Equal(t, 0.2, prop.Thickness)
	})

	t.Run("when orientation is empty, should apply horizontal", func(t *testing.T) {
		prop := properties.Line{
			Orientation: "",
		}
		prop.MakeValid()

		assert.Equal(t, orientation.Horizontal, prop.Orientation)
	})

	t.Run("when offset percent is less than 5, should apply 5", func(t *testing.T) {
		prop := properties.Line{
			OffsetPercent: 4,
		}
		prop.MakeValid()

		assert.Equal(t, 5.0, prop.OffsetPercent)
	})

	t.Run("when offset percent is greater than 95, should apply 95", func(t *testing.T) {
		prop := properties.Line{
			OffsetPercent: 96,
		}
		prop.MakeValid()

		assert.Equal(t, 95.0, prop.OffsetPercent)
	})

	t.Run("when size percent is less than 1, should apply 90", func(t *testing.T) {
		prop := properties.Line{
			SizePercent: 0,
		}
		prop.MakeValid()

		assert.Equal(t, 90.0, prop.SizePercent)
	})

	t.Run("when size percent is greater than 100, should apply 100", func(t *testing.T) {
		prop := properties.Line{
			SizePercent: 101,
		}
		prop.MakeValid()

		assert.Equal(t, 100.0, prop.SizePercent)
	})
}

func TestLine_ToMap(t *testing.T) {
	t.Run("when line is nil, should return nil", func(t *testing.T) {
		var prop *properties.Line
		m := prop.ToMap()

		assert.Nil(t, m)
	})

	t.Run("when line is filled, should return map filled", func(t *testing.T) {
		prop := fixture.LineProp()
		m := prop.ToMap()

		assert.Equal(t, "RGB(100, 50, 200)", m["prop_color"])
		assert.Equal(t, linestyle.Dashed, m["prop_style"])
		assert.Equal(t, 1.1, m["prop_thickness"])
		assert.Equal(t, orientation.Vertical, m["prop_orientation"])
		assert.Equal(t, 50.0, m["prop_offset_percent"])
		assert.Equal(t, 20.0, m["prop_size_percent"])
	})
}
