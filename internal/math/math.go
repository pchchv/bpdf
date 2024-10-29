package math

import "github.com/pchchv/bpdf/core/entity"

type math struct{}

// New create a Math.
func New() *math {
	return &math{}
}

// Resize adjusts the internal dimension of an element to occupy a percentage of the available space
//   - inner: The inner dimensions of the element
//   - outer: The outer dimensions of the element
//   - percent: The percentage of the external dimension that can be occupied
//   - justReferenceWidth: Indicates whether resizing should be done only in relation to width or in relation to width and height
func (s *math) Resize(inner *entity.Dimensions, outer *entity.Dimensions, percent float64, justReferenceWidth bool) *entity.Dimensions {
	percent /= 100.0
	innerProportion := inner.Height / inner.Width
	outerProportion := outer.Height / outer.Width
	newInnerWidth := 0.0
	if innerProportion > outerProportion && !justReferenceWidth {
		newInnerWidth = outer.Height / innerProportion * percent
	} else {
		newInnerWidth = outer.Width * percent
	}

	newInnerHeight := newInnerWidth * innerProportion
	if justReferenceWidth && newInnerHeight > outer.Height {
		newInnerWidth = outer.Height / innerProportion * 1
		newInnerHeight = newInnerWidth * innerProportion
	}

	return &entity.Dimensions{Width: newInnerWidth, Height: newInnerHeight}
}
