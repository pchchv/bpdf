package entity

import "github.com/pchchv/bpdf/consts/protection"

func fixtureProtection() Protection {
	return Protection{
		Type:          protection.Print,
		OwnerPassword: "123456",
		UserPassword:  "654321",
	}
}
