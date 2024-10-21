package entity

import (
	"testing"

	"github.com/pchchv/bpdf/consts/protection"
	"github.com/stretchr/testify/assert"
)

func TestProtection_AppendMap(t *testing.T) {
	sut := fixtureProtection()
	m := make(map[string]interface{})
	m = sut.AppendMap(m)

	assert.Equal(t, sut.Type, m["config_protection_type"])
	assert.Equal(t, sut.UserPassword, m["config_user_password"])
	assert.Equal(t, sut.OwnerPassword, m["config_owner_password"])
}

func fixtureProtection() Protection {
	return Protection{
		Type:          protection.Print,
		OwnerPassword: "123456",
		UserPassword:  "654321",
	}
}
