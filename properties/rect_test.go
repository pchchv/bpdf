package properties_test

import (
	"testing"

	"github.com/pchchv/bpdf/internal/fixture"
	"github.com/pchchv/bpdf/properties"
	"github.com/stretchr/testify/assert"
)

func TestRect_MakeValid(t *testing.T) {
	t.Run("when percent is less than zero, should become 100", func(t *testing.T) {
		prop := properties.Rect{Percent: -2}
		prop.MakeValid()

		assert.Equal(t, prop.Percent, 100.0)
	})

	t.Run("when percent is greater than 100, should become 100", func(t *testing.T) {
		prop := properties.Rect{Percent: 102}
		prop.MakeValid()

		assert.Equal(t, prop.Percent, 100.0)
	})

	t.Run("when is center, top and left should become 0", func(t *testing.T) {
		prop := properties.Rect{Center: true, Top: 5, Left: 5}
		prop.MakeValid()

		assert.Equal(t, prop.Top, 0.0)
		assert.Equal(t, prop.Left, 0.0)
	})

	t.Run("when left is less than 0, should become 0", func(t *testing.T) {
		prop := properties.Rect{Left: -5}
		prop.MakeValid()

		assert.Equal(t, prop.Left, 0.0)
	})

	t.Run("when top is less than 0, should become 0", func(t *testing.T) {
		prop := properties.Rect{Top: -5}
		prop.MakeValid()

		assert.Equal(t, prop.Top, 0.0)
	})
}

func TestRect_ToMap(t *testing.T) {
	sut := fixture.RectProp()
	sut.Center = true
	m := sut.ToMap()

	assert.Equal(t, 10.0, m["prop_left"])
	assert.Equal(t, 10.0, m["prop_top"])
	assert.Equal(t, 98.0, m["prop_percent"])
	assert.Equal(t, true, m["prop_center"])
}
