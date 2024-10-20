// The extension package contains all image extensions.
package extension

// Type is a representation of a Image extension.
type Extension string

const (
	Jpg  Extension = "jpg"
	Jpeg Extension = "jpeg"
	Png  Extension = "png"
)

// IsValid checks if the extension is valid.
func (t Extension) IsValid() bool {
	return t == Jpg || t == Jpeg || t == Png
}
