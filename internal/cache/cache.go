package cache

import (
	"errors"
	"os"

	"github.com/pchchv/bpdf/consts/extension"
	"github.com/pchchv/bpdf/core/entity"
)

type cache struct {
	images map[string]*entity.Image
	codes  map[string][]byte
}

// LoadImage loads an image from a file.
func (c *cache) LoadImage(file string, extension extension.Extension) error {
	imageBytes, err := os.ReadFile(file)
	if err != nil {
		return err
	}

	img := &entity.Image{Bytes: imageBytes, Extension: extension}
	c.images[file+string(extension)] = img
	return nil
}

// AddImage adds an image to the cache.
func (c *cache) AddImage(value string, image *entity.Image) {
	c.images[value+string(image.Extension)] = image
}

// GetImage returns an image from the cache.
func (c *cache) GetImage(file string, extension extension.Extension) (*entity.Image, error) {
	image, ok := c.images[file+string(extension)]
	if ok {
		return image, nil
	}

	return nil, errors.New("image not found")
}
