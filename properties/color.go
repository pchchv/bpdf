package properties

import "fmt"

var (
	WhiteColor = Color{Red: 255, Green: 255, Blue: 255}
	BlackColor = Color{Red: 0, Green: 0, Blue: 0}
	RedColor   = Color{Red: 255, Green: 0, Blue: 0}
	GreenColor = Color{Red: 0, Green: 255, Blue: 0}
	BlueColor  = Color{Red: 0, Green: 0, Blue: 255}
)

// Color represents a color in RGB space,
// mixing of values is possible,
// when all values are 0, the color is black,
// when all values are 255, the color is white.
type Color struct {
	Red   int
	Green int
	Blue  int
}

// ToString returns a string representation of the Color.
func (c *Color) ToString() string {
	if c == nil {
		return ""
	}
	return fmt.Sprintf("RGB(%d, %d, %d)", c.Red, c.Green, c.Blue)
}
