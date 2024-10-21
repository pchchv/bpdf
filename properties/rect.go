package properties

// Rect represents properties from a rectangle (Image, QrCode or Barcode) inside a cell.
type Rect struct {
	// Left is the space between the left cell boundary to the rectangle, if center is false.
	Left float64
	// Top is space between the upper cell limit to the barcode, if center is false.
	Top float64
	// Percent is how much the rectangle will occupy the cell,
	// ex 100%: The rectangle will fulfill the entire cell
	// ex 50%: The greater side from the rectangle will have half the size of the cell.
	Percent float64
	// Indicate whether only the width should be used as a reference to calculate the component size, disregarding the height
	// ex true: The component will be scaled only based on the available width, disregarding the available height
	JustReferenceWidth bool
	// Center define that the barcode will be vertically and horizontally centralized.
	Center bool
}
