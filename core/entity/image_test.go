package entity

import "github.com/pchchv/bpdf/consts/extension"

func fixtureImage() Image {
	dimensions := fixtureDimensions()

	return Image{
		Bytes:      []byte{1, 2, 3},
		Extension:  extension.Png,
		Dimensions: &dimensions,
	}
}
