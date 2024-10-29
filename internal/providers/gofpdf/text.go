package gofpdf

import (
	"strings"
	"unicode"

	"github.com/pchchv/bpdf/consts/align"
	"github.com/pchchv/bpdf/consts/fontfamily"
	"github.com/pchchv/bpdf/core"
	"github.com/pchchv/bpdf/internal/providers/gofpdf/fpdfwrapper"
	"github.com/pchchv/bpdf/properties"
)

type text struct {
	pdf  fpdfwrapper.Fpdf
	math core.Math
	font core.Font
}

// NewText create a Text.
func NewText(pdf fpdfwrapper.Fpdf, math core.Math, font core.Font) *text {
	return &text{
		pdf,
		math,
		font,
	}
}

func (s *text) textToUnicode(txt string, props *properties.Text) string {
	if props.Family == fontfamily.Arial ||
		props.Family == fontfamily.Helvetica ||
		props.Family == fontfamily.Symbol ||
		props.Family == fontfamily.ZapBats ||
		props.Family == fontfamily.Courier {
		translator := s.pdf.UnicodeTranslatorFromDescriptor("")
		return translator(txt)
	}

	return txt
}

func (s *text) addLine(textProp *properties.Text, xColOffset, colWidth, yColOffset, textWidth float64, text string) {
	left, top, _, _ := s.pdf.GetMargins()
	fontHeight := s.font.GetHeight(textProp.Family, textProp.Style, textProp.Size)
	if textProp.Align == align.Left {
		s.pdf.Text(xColOffset+left, yColOffset+top, text)
		if textProp.Hyperlink != nil {
			s.pdf.LinkString(xColOffset+left, yColOffset+top-fontHeight, textWidth, fontHeight, *textProp.Hyperlink)
		}

		return
	}

	if textProp.Align == align.Justify {
		const spaceString = " "
		const emptyString = ""
		text = strings.TrimRight(text, spaceString)
		textNotSpaces := strings.ReplaceAll(text, spaceString, emptyString)
		textWidth = s.pdf.GetStringWidth(textNotSpaces)
		defaultSpaceWidth := s.pdf.GetStringWidth(spaceString)
		words := strings.Fields(text)
		numSpaces := max(len(words)-1, 1)
		spaceWidth := (colWidth - textWidth) / float64(numSpaces)
		x := xColOffset + left
		if isIncorrectSpaceWidth(textWidth, spaceWidth, defaultSpaceWidth, textNotSpaces) {
			spaceWidth = defaultSpaceWidth
		}

		var finishX float64
		initX := x
		for _, word := range words {
			s.pdf.Text(x, yColOffset+top, word)
			finishX = x + s.pdf.GetStringWidth(word)
			x = finishX + spaceWidth
		}

		if textProp.Hyperlink != nil {
			s.pdf.LinkString(initX, yColOffset+top-fontHeight, finishX-initX, fontHeight, *textProp.Hyperlink)
		}

		return
	}

	modifier := 2.0
	if textProp.Align == align.Right {
		modifier = 1
	}

	dx := (colWidth - textWidth) / modifier
	if textProp.Hyperlink != nil {
		s.pdf.LinkString(dx+xColOffset+left, yColOffset+top-fontHeight, textWidth, fontHeight, *textProp.Hyperlink)
	}

	s.pdf.Text(dx+xColOffset+left, yColOffset+top, text)
}

func isIncorrectSpaceWidth(textWidth, spaceWidth, defaultSpaceWidth float64, text string) bool {
	if textWidth <= 0 || spaceWidth <= defaultSpaceWidth*10 {
		return false
	}

	lastChar := rune(text[len(text)-1])
	return !unicode.IsLetter(lastChar) && !unicode.IsNumber(lastChar)
}
