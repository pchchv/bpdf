package cache

import "sync"

type mutexCache struct {
	inner      Cache
	imageMutex sync.RWMutex
}
