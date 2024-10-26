package cache

import "github.com/pchchv/bpdf/core/entity"

type cache struct {
	images map[string]*entity.Image
	codes  map[string][]byte
}
