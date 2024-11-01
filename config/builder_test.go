package config_test

import (
	"fmt"
	"testing"

	"github.com/pchchv/bpdf/config"
	"github.com/stretchr/testify/assert"
)

func TestNewBuilder(t *testing.T) {
	sut := config.NewBuilder()

	assert.NotNil(t, sut)
	assert.Equal(t, "*config.CfgBuilder", fmt.Sprintf("%T", sut))
}
