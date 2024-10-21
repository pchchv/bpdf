package properties_test

import (
	"testing"

	"github.com/pchchv/bpdf/properties"
	"github.com/stretchr/testify/assert"
)

func TestWhiteColor(t *testing.T) {
	sut := properties.WhiteColor

	assert.Equal(t, 255, sut.Red)
	assert.Equal(t, 255, sut.Green)
	assert.Equal(t, 255, sut.Blue)
}

func TestBlackColor(t *testing.T) {
	sut := properties.BlackColor

	assert.Equal(t, 0, sut.Red)
	assert.Equal(t, 0, sut.Green)
	assert.Equal(t, 0, sut.Blue)
}

func TestRedColor(t *testing.T) {
	sut := properties.RedColor

	assert.Equal(t, 255, sut.Red)
	assert.Equal(t, 0, sut.Green)
	assert.Equal(t, 0, sut.Blue)
}

func TestGreenColor(t *testing.T) {
	sut := properties.GreenColor

	assert.Equal(t, 0, sut.Red)
	assert.Equal(t, 255, sut.Green)
	assert.Equal(t, 0, sut.Blue)
}

func TestBlueColor(t *testing.T) {
	blue := properties.BlueColor

	assert.Equal(t, 0, blue.Red)
	assert.Equal(t, 0, blue.Green)
	assert.Equal(t, 255, blue.Blue)
}
