package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMargins_AppendMap(t *testing.T) {
	sut := fixtureMargins()
	m := make(map[string]interface{})
	m = sut.AppendMap(m)

	assert.Equal(t, 20.0, m["config_margin_left"])
	assert.Equal(t, 30.0, m["config_margin_top"])
	assert.Equal(t, 40.0, m["config_margin_right"])
	assert.Equal(t, 50.0, m["config_margin_bottom"])
}

func fixtureMargins() Margins {
	return Margins{
		Left:   20,
		Top:    30,
		Right:  40,
		Bottom: 50,
	}
}
