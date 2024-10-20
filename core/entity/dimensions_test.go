package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDimensions_AppendMap(t *testing.T) {
	sut := fixtureDimensions()
	m := make(map[string]interface{})
	m = sut.AppendMap("label", m)

	assert.Equal(t, 100.0, m["label_dimension_width"])
	assert.Equal(t, 200.0, m["label_dimension_height"])
}

func fixtureDimensions() Dimensions {
	return Dimensions{
		Width:  100,
		Height: 200,
	}
}
