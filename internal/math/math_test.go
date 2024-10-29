package math_test

import (
	"fmt"
	"testing"

	"github.com/pchchv/bpdf/internal/math"
	"github.com/stretchr/testify/assert"
)

func TestNewMath(t *testing.T) {
	sut := math.New()

	assert.NotNil(t, sut)
	assert.Equal(t, "*math.math", fmt.Sprintf("%T", sut))
}
