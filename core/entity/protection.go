package entity

import "github.com/pchchv/bpdf/consts/protection"

// Protection is the representation of a pdf protection.
type Protection struct {
	Type          protection.Protection
	UserPassword  string
	OwnerPassword string
}
