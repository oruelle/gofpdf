package gofpdf

type Font struct {
	family string
	style  string
	size   float64
}

// Create a new Font struct given family style and size (user unit)
func NewFont(family, style string, size float64) *Font {
	f := new(Font)

	f.family = family
	f.style = style
	f.size = size

	return f
}

// Create a new Font struct given family style and size (user unit)
func NewFontFromCurrent(fp *Fpdf) *Font {
	f := new(Font)

	f.FromCurrentFont(fp)

	return f
}

// Copy this instance font to new one.
func (f *Font) Copy() *Font {
	_new_font := new(Font)

	_new_font.family, _new_font.style, _new_font.size = f.family, f.style, f.size

	return _new_font
}

// Copy this instance font to current font.
func (f *Font) ToCurrentFont(fp *Fpdf) {
	if fp != nil {
		fp.SetFont(f.family, f.style, 0)
		fp.SetFontUnitSize(f.size)
	}
}

// Copy the current font to this instance font.
func (f *Font) FromCurrentFont(fp *Fpdf) *Font {
	if fp != nil {
		f.family, f.style, f.size = fp.GetFont()
	}

	return f
}

// Family setter
func (f *Font) SetFamily(family string) *Font {
	f.family = family
	return f
}

// Style setter
func (f *Font) SetStyle(style string) *Font {
	f.style = style
	return f
}

// Size setter (unit size)
func (f *Font) SetSize(size float64) *Font {
	f.size = size
	return f
}
