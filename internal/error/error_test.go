package error_test

import (
	"testing"

	"github.com/pchchv/bpdf/consts/fontfamily"
	"github.com/pchchv/bpdf/consts/fontstyle"
	"github.com/pchchv/bpdf/internal/error"
	"github.com/stretchr/testify/assert"
)

func TestDefaultErrorText(t *testing.T) {
	assert.Equal(t, fontfamily.Arial, error.DefaultErrorText.Family)
	assert.Equal(t, fontstyle.Bold, error.DefaultErrorText.Style)
	assert.Equal(t, 10.0, error.DefaultErrorText.Size)
	assert.Equal(t, 255, error.DefaultErrorText.Color.Red)
	assert.Equal(t, 0, error.DefaultErrorText.Color.Green)
	assert.Equal(t, 0, error.DefaultErrorText.Color.Blue)
}
