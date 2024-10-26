package cache_test

import (
	"fmt"
	"os"
	"path"
	"strings"
	"testing"

	"github.com/pchchv/bpdf/consts/extension"
	"github.com/pchchv/bpdf/core/entity"
	"github.com/pchchv/bpdf/internal/cache"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	sut := cache.New()

	assert.NotNil(t, sut)
	assert.Equal(t, "*cache.cache", fmt.Sprintf("%T", sut))
}

func TestCache_GetImage(t *testing.T) {
	t.Run("when cannot get image, should return error", func(t *testing.T) {
		sut := cache.New()
		img, err := sut.GetImage("image", extension.Jpg)

		assert.Nil(t, img)
		assert.NotNil(t, err)
	})
	t.Run("when can get image, should return image", func(t *testing.T) {
		sut := cache.New()
		sut.AddImage("image", &entity.Image{
			Extension: extension.Jpg,
		})
		img, err := sut.GetImage("image", extension.Jpg)

		assert.NotNil(t, img)
		assert.Nil(t, err)
	})
}

func TestCache_AddImage(t *testing.T) {
	t.Run("when add image, return works", func(t *testing.T) {
		sut := cache.New()
		sut.AddImage("image", &entity.Image{
			Extension: extension.Jpg,
		})

		img, err := sut.GetImage("image", extension.Jpg)
		assert.NotNil(t, img)
		assert.Nil(t, err)
	})
}

func TestCache_LoadImage(t *testing.T) {
	t.Run("when cannot find image, should return error", func(t *testing.T) {
		sut := cache.New()
		err := sut.LoadImage("image", extension.Jpg)

		assert.NotNil(t, err)
	})

	t.Run("when can find image, should not return error and find image", func(t *testing.T) {
		sut := cache.New()

		err := sut.LoadImage(buildPath("/docs/assets/images/biplane.jpg"), extension.Jpg)

		assert.Nil(t, err)
		img, err := sut.GetImage(buildPath("/docs/assets/images/biplane.jpg"), extension.Jpg)
		assert.Nil(t, err)
		assert.NotNil(t, img)
	})
}

func buildPath(file string) (dir string) {
	var err error
	if dir, err = os.Getwd(); err != nil {
		return ""
	}

	dir = strings.ReplaceAll(dir, "internal/cache", "")
	return path.Join(dir, file)
}
