// The align package contains all align types.
package align

// Align is a representation of a column align.
type Align string

const (
	Left    Align = "L" // left horizontal alignment
	Right   Align = "R" // right horizontal alignment
	Center  Align = "C" // center horizontal and vertical alignment
	Top     Align = "T" // top vertical alignment
	Bottom  Align = "B" // bottom vertical alignment
	Middle  Align = "M" // middle alignment
	Justify       = "J" // horizontal alignment that evenly distributes text between the left and right margins
)
