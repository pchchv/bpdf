package gofpdf

import (
	"errors"

	"github.com/pchchv/bpdf/consts/extension"
	"github.com/pchchv/bpdf/core/entity"
)

func FromBytes(bytes []byte, ext extension.Extension) (*entity.Image, error) {
	if !ext.IsValid() {
		return nil, errors.New("invalid image format")
	}

	return &entity.Image{
		Bytes:     bytes,
		Extension: ext,
	}, nil
}
