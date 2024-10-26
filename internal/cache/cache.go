package cache

import "github.com/pchchv/bpdf/core/entity"

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
