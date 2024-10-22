// The border package contains all border types.
package border

// Type represents a border type.
type Border string

const (
	None   Border = ""  // default border type
	Full   Border = "1" // borders all sides
	Left   Border = "L"
	Top    Border = "T"
	Right  Border = "R"
	Bottom Border = "B"
)

// IsValid checks if the border type is valid.
func (t Border) IsValid() bool {
	return t == Full || t == Left || t == Top || t == Right || t == Bottom
}
