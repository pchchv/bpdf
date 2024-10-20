package entity

import "fmt"

// Utf8Text representats a text with a flag to indicate if it's UTF8.
type Utf8Text struct {
	Text string
	UTF8 bool
}

// ToString returns a string representation of the text.
func (u *Utf8Text) ToString() string {
	return fmt.Sprintf("Utf8Text(%s, %v)", u.Text, u.UTF8)
}
