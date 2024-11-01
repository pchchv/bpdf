package cache_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/pchchv/bpdf/consts/extension"
	"github.com/pchchv/bpdf/core/entity"
	"github.com/pchchv/bpdf/internal/cache"
	"github.com/pchchv/bpdf/mocks"
	"github.com/stretchr/testify/assert"
)

func TestNewMutexDecorator(t *testing.T) {
	sut := cache.NewMutexDecorator(nil)

	assert.NotNil(t, sut)
	assert.Equal(t, "*cache.mutexCache", fmt.Sprintf("%T", sut))
}

func TestMutexCache_AddImage(t *testing.T) {
	value := "value1"
	img := &entity.Image{}
	innerMock := mocks.NewCache(t)
	innerMock.EXPECT().AddImage(value, img)
	sut := cache.NewMutexDecorator(innerMock)

	sut.AddImage(value, img)

	innerMock.AssertNumberOfCalls(t, "AddImage", 1)
}

func TestMutexCache_GetImage(t *testing.T) {
	value := "value2"
	ext := extension.Jpg
	imgToReturn := &entity.Image{}
	errToReturn := errors.New("any error")
	innerMock := mocks.NewCache(t)
	innerMock.EXPECT().GetImage(value, ext).Return(imgToReturn, errToReturn)
	sut := cache.NewMutexDecorator(innerMock)

	img, err := sut.GetImage(value, ext)

	assert.Equal(t, imgToReturn, img)
	assert.Equal(t, errToReturn, err)
	innerMock.AssertNumberOfCalls(t, "GetImage", 1)
}

func TestMutexCache_LoadImage(t *testing.T) {
	value := "value3"
	ext := extension.Jpg
	errToReturn := errors.New("any error")
	innerMock := mocks.NewCache(t)
	innerMock.EXPECT().LoadImage(value, ext).Return(errToReturn)
	sut := cache.NewMutexDecorator(innerMock)

	err := sut.LoadImage(value, ext)

	assert.Equal(t, errToReturn, err)
	innerMock.AssertNumberOfCalls(t, "LoadImage", 1)
}
