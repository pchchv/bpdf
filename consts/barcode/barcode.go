package barcode

type Code string

const (
	Code128   Code = "code128"
	EAN       Code = "ean"
	Codabar   Code = "codabar"
	Code39    Code = "code39"
	Code93    Code = "code93"
	PDF417    Code = "pdf417"
	TwoOfFive Code = "twooffive"
)
