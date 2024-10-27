package fpdfwrapper_test

import (
	"fmt"
	"testing"

	"github.com/jung-kurt/gofpdf"
	"github.com/pchchv/bpdf/internal/providers/fpdfwrapper"
	"github.com/stretchr/testify/assert"
)

func TestNewCustom(t *testing.T) {
	sut := fpdfwrapper.NewCustom(&gofpdf.InitType{})

	assert.NotNil(t, "", fmt.Sprintf("%T", sut))
}
