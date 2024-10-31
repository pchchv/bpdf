package code_test

import (
	"fmt"
	"testing"

	"github.com/pchchv/bpdf/consts/barcode"
	"github.com/pchchv/bpdf/core/entity"
	"github.com/pchchv/bpdf/internal/code"
	"github.com/pchchv/bpdf/properties"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	t.Run("constructor", func(t *testing.T) {
		sut := code.New()

		assert.NotNil(t, sut)
		assert.Equal(t, "*code.code", fmt.Sprintf("%T", sut))
	})

	t.Run("singleton is applied", func(t *testing.T) {
		sut1 := code.New()
		sut2 := code.New()

		assert.NotNil(t, sut1)
		assert.NotNil(t, sut2)
	})
}

func TestCode_GenQr(t *testing.T) {
	t.Run("When cannot generate qr code, should return error", func(t *testing.T) {
		sut := code.New()
		data := genStringWithLength(5000)

		bytes, err := sut.GenQr(data)

		assert.NotNil(t, err)
		assert.Nil(t, bytes)
	})

	t.Run("When can generate qr code, should return bytes", func(t *testing.T) {
		sut := code.New()
		data := genStringWithLength(50)

		bytes, err := sut.GenQr(data)

		assert.NotNil(t, bytes)
		assert.Nil(t, err)
	})
}

func TestCode_GenBar(t *testing.T) {
	t.Run("When cannot generate bar code, should return error", func(t *testing.T) {
		sut := code.New()
		cell := &entity.Cell{
			X:      10,
			Y:      10,
			Width:  100,
			Height: 100,
		}
		prop := &properties.Barcode{}
		prop.MakeValid()
		data := genStringWithLength(5000)

		bytes, err := sut.GenBar(data, cell, prop)

		assert.NotNil(t, err)
		assert.Nil(t, bytes)
	})

	t.Run("When can generate bar code, should return bytes", func(t *testing.T) {
		sut := code.New()
		cell := &entity.Cell{
			X:      10,
			Y:      10,
			Width:  100,
			Height: 100,
		}
		prop := &properties.Barcode{}
		prop.MakeValid()
		data := genStringWithLength(60)

		bytes, err := sut.GenBar(data, cell, prop)

		assert.NotNil(t, bytes)
		assert.Nil(t, err)
	})

	t.Run("When is ean and can generate bar code, should return bytes", func(t *testing.T) {
		sut := code.New()
		cell := &entity.Cell{
			X:      10,
			Y:      10,
			Width:  100,
			Height: 100,
		}
		prop := &properties.Barcode{
			Type: barcode.EAN,
		}
		prop.MakeValid()

		bytes, err := sut.GenBar("123456789123", cell, prop)

		assert.NotNil(t, bytes)
		assert.Nil(t, err)
	})
}

func genStringWithLength(length int) (content string) {
	for i := 0; i < length; i++ {
		content += "a"
	}
	return
}
