package properties_test

import (
	"testing"

	"github.com/pchchv/bpdf/internal/fixture"
	"github.com/stretchr/testify/assert"
)

func TestRect_ToMap(t *testing.T) {
	sut := fixture.RectProp()
	sut.Center = true
	m := sut.ToMap()

	assert.Equal(t, 10.0, m["prop_left"])
	assert.Equal(t, 10.0, m["prop_top"])
	assert.Equal(t, 98.0, m["prop_percent"])
	assert.Equal(t, true, m["prop_center"])
}
