// The protection package contains all protection types.
package protection

type Protection byte

const (
	None       Protection = 0
	Print      Protection = 4
	Modify     Protection = 8
	Copy       Protection = 16
	AnnotForms Protection = 32
)
