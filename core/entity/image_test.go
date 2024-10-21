package entity

import (
	"testing"

	"github.com/pchchv/bpdf/consts/extension"
	"github.com/stretchr/testify/assert"
)

func TestImage_AppendMap(t *testing.T) {
	sut := fixtureImage()
	m := make(map[string]interface{})
	m = sut.AppendMap(m)

	assert.Equal(t, "[1 2 3]", m["entity_image_bytes"])
	assert.Equal(t, extension.Png, m["entity_extension"])
	assert.Equal(t, 100.0, m["background_dimension_width"])
	assert.Equal(t, 200.0, m["background_dimension_height"])
}

func fixtureImage() Image {
	dimensions := fixtureDimensions()

	return Image{
		Bytes:      []byte{1, 2, 3},
		Extension:  extension.Png,
		Dimensions: &dimensions,
	}
}
