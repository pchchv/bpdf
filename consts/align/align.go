// The align package contains all align types.
package align

// Type is a representation of a column align.
type Type string

const (
	Left    Type = "L" // left horizontal alignment
	Right   Type = "R" // right horizontal alignment
	Center  Type = "C" // center horizontal and vertical alignment
	Top     Type = "T" // top vertical alignment
	Bottom  Type = "B" // bottom vertical alignment
	Middle  Type = "M" // middle alignment
	Justify      = "J" // horizontal alignment that evenly distributes text between the left and right margins
)
