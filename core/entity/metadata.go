package entity

import (
	"fmt"
	"time"
)

// Utf8Text representats a text with a flag to indicate if it's UTF8.
type Utf8Text struct {
	Text string
	UTF8 bool
}

// ToString returns a string representation of the text.
func (u *Utf8Text) ToString() string {
	return fmt.Sprintf("Utf8Text(%s, %v)", u.Text, u.UTF8)
}

// Metadata is the representation of a PDF metadata.
type Metadata struct {
	Author       *Utf8Text
	Creator      *Utf8Text
	Subject      *Utf8Text
	Title        *Utf8Text
	CreationDate *time.Time
	KeywordsStr  *Utf8Text
}
