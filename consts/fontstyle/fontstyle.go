// The fontstyle package contains all default font styles.
package fontstyle

// Fontstyle is a representation of a style DefaultFont.
type Fontstyle string

const (
	Bold       Fontstyle = "B"
	Normal     Fontstyle = ""
	Italic     Fontstyle = "I"
	BoldItalic Fontstyle = "BI"
)

// IsValid checks if the style is valid.
func (s Fontstyle) IsValid() bool {
	return s == Normal || s == Italic || s == BoldItalic || s == Bold
}
