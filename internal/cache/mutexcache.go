package cache

import (
	"sync"

	"github.com/pchchv/bpdf/consts/extension"
	"github.com/pchchv/bpdf/core/entity"
)

type mutexCache struct {
	inner      Cache
	imageMutex sync.RWMutex
}

// AddImage adds a behavior to lock/unlock cache write.
func (c *mutexCache) AddImage(value string, image *entity.Image) {
	c.imageMutex.Lock()
	defer c.imageMutex.Unlock()

	c.inner.AddImage(value, image)
}

// GetImage adds a behavior to lock/unlock cache read.
func (c *mutexCache) GetImage(file string, extension extension.Extension) (*entity.Image, error) {
	return c.inner.GetImage(file, extension)
}
