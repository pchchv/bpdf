package code_test

import (
	"fmt"
	"testing"

	"github.com/pchchv/bpdf/internal/code"
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

func genStringWithLength(length int) (content string) {
	for i := 0; i < length; i++ {
		content += "a"
	}
	return
}
