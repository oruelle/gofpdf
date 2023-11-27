/*
 * Copyright (c) 2023-2025 Olivier Ruelle (github.com/oruelle)
 *
 * Permission to use, copy, modify, and distribute this software for any
 * purpose with or without fee is hereby granted, provided that the above
 * copyright notice and this permission notice appear in all copies.
 *
 * THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
 * WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
 * MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
 * ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
 * WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
 * ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
 * OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
 */

package gofpdf

import (
	"fmt"
	"math"
	"strings"
)

const ALIGN_LEFT = "left"
const ALIGN_RIGHT = "right"
const ALIGN_CENTER = "right"
const ALIGN_JUSTIFY = "justify"

// Title() insert a title in the current page. Lvl parameter adjust font size (0 : Master title).
// Fill color can be precise (nil otherwise)
func (fp *Fpdf) Title(title string, lvl uint8, fontColor, fillColor *Color) (err error) {
	coef_font := (2.0 - float64(lvl)*0.25) // coef for font

	// Backup colors
	_text_color := NewColor().FromTextColor(fp)
	_fill_color := NewColor().FromFillColor(fp)

	// Before spacing
	ww, _ := fp.GetWorkingSize()
	//fp.CellFormat(ww, fp.fontSize * coef_fon, "", "0", 1, "LM", false, 0, "")

	// Set title colors and font
	if fontColor == nil {
		fontColor = NewRGBColor(0, 0, 0)
	}
	fontColor.ToTextColor(fp)
	if fillColor == nil {
		fillColor = NewRGBColor(255, 255, 255)
	}
	fillColor.ToFillColor(fp)

	font_size_pt := coef_font * fp.fontSizePt
	if font_size_pt < fp.fontSizePt {
		font_size_pt = fp.fontSizePt
	}
	err = fp.SetFont(fp.fontFamily, "B", font_size_pt)
	if err != nil {
		return
	}

	// Creat title
	if lvl == 0 {
		err = fp.MultiCell(ww, fp.fontSize*2, title, "", "CM", true)
	} else {
		err = fp.MultiCell(ww, fp.fontSize*2, title, "", "LM", true)
	}
	if err != nil {
		return
	}

	// After spacing
	err = fp.CellFormat(ww, fp.fontSize*2, "", "0", 1, "LM", false, 0, "")
	if err != nil {
		return
	}

	//Reset colors and font
	_text_color.ToTextColor(fp)
	_fill_color.ToFillColor(fp)

	err = fp.SetFont(fp.fontFamily, "", fp.fontSize)

	return
}

// MultiCellOnFixedHeight enable to fix height of the Cell, which means
// if the text is multiline the height of the Cell will not change
// and text height line will shrink.
func (fp *Fpdf) MultiCellOnFixedHeight(width, height float64, text, border, align string, fill bool) {
	w := fp.GetStringWidth(text)
	cellHeight := height
	if w > width-2*fp.GetCellMargin() {
		cellHeight /= math.Floor(w/(width-2*fp.GetCellMargin())) + 1
	}

	fp.MultiCell(width, cellHeight, text, border, align, fill)
}

// Table builds a table with the given content with the given width.
func (fp *Fpdf) Table(width float64, table [][]string, align []string, header, evenOdd bool) (err error) {

	x_orig := fp.GetX()
	height := fp.fontSize * 1.2

	curent_fill := NewColor().FromFillColor(fp)
	orig_font := NewFontFromCurrent(fp)
	new_font := orig_font.Copy()

	for n_row, row := range table {
		nb_col := float64(len(row))
		for n, cell := range row {
			ln := 0
			// Line return
			if n == 0 {
				fp.SetX(x_orig)
			}
			// Line return
			if n == len(row)-1 {
				ln = 1
			}
			// for alignment
			_align := "CM"
			if len(row) == len(align) {
				_align = align[n]
			}
			if evenOdd && n_row%2 == 0 {
				LIGHT_GREY.ToFillColor(fp)
			} else {
				WHITE.ToFillColor(fp)
			}
			if header && n_row == 0 {
				new_font.SetStyle("B").ToCurrentFont(fp)
			} else {
				new_font.SetStyle("").ToCurrentFont(fp)
			}

			err = fp.CellFormat(width/nb_col, height, cell, "1", ln, _align, true, 0, "")
			if err != nil {
				return
			}
		}
	}

	orig_font.ToCurrentFont(fp)
	curent_fill.ToFillColor(fp)

	var x, y float64
	x, y = fp.GetXY()
	fp.SetXY(x, y+height*2)

	return
}

// TableX is similar to Table but is placed at the given
// X position not to the current one.
func (fp *Fpdf) TableX(x, width float64, table [][]string, align []string, header, evenOdd bool) {
	// Backup current position
	x_orig := fp.GetX()

	// Set new position
	fp.SetX(x)

	// Create Table
	fp.Table(width, table, align, header, evenOdd)

	// Reset Position
	fp.SetX(x_orig)
}

// TableX is similar to Table but is placed at the given
// X position not to the current one.
func (fp *Fpdf) TableXY(x, y, width float64, table [][]string, align []string, header, evenOdd bool) {
	// Backup current position
	x_orig, y_orig := fp.GetXY()

	// Set new position
	fp.SetXY(x, y)

	// Create Table
	fp.Table(width, table, align, header, evenOdd)

	// Reset Position
	fp.SetXY(x_orig, y_orig)
}

// TableX is similar to Table but is placed at the given
// X position not to the current one.
func (fp *Fpdf) TableXCenter(width float64, table [][]string, align []string, header, evenOdd bool) {
	// Create Table
	fp.TableX(fp.workingCenter-width/2, width, table, align, header, evenOdd)
}

// Parag build a paragraph with the given text on the given width
// from the current position and with the given line height
// (if height <=0 the default is set to 120% of font size).
func (fp *Fpdf) Parag(width, height float64, textStr, alignStr string) {
	if height < fp.fontSize {
		height = fp.fontSize * 1.2
	}
	var lines []string

	if fp.isCurrentUTF8 {
		lines = fp.SplitText(textStr, width)
	} else {
		lines = fp.SplitText(textStr, width)
	}

	var x, y float64
	x, y = fp.GetXY()
	leftpos := x

	for n, line := range lines {

		lineWidth := fp.GetStringWidth(line)

		switch alignStr {
		case ALIGN_LEFT:
			// Write text
			fp.Text(x, y, line)
			// Reset position for next line
			x = leftpos
			y += height

			// Page break
			if y > fp.workingHeight+fp.tMargin {
				fp.AddPage()
				//fp.SetHomeXY()
				x, y = fp.GetXY()

				x = leftpos
				y += height
			}
		case ALIGN_JUSTIFY:
			words := strings.Split(line, " ")
			nb_words := len(words)
			words_width := 0.0
			for _, word := range words {
				words_width += fp.GetStringWidth(word)
			}
			// original space width
			orig_space_width := (lineWidth - words_width) / float64(nb_words-1)
			// space size to justify text
			space_width := (width - words_width) / float64(nb_words-1)

			// if last line then no justify
			if n == len(lines)-1 {
				space_width = orig_space_width
			}

			for _, word := range words {
				fp.Text(x, y, word)
				wordWidth := fp.GetStringWidth(word)
				x += space_width + wordWidth
			}
			// Reset position for next line
			x = leftpos
			y += height

			// Page break
			if y > fp.workingHeight+fp.tMargin {
				fp.AddPage()
				//fp.SetHomeXY()
				x, y = fp.GetXY()

				x = leftpos
				y += height
			}
		}
	}

	// Reset Position with y space
	fp.SetXY(x, y+height)
}

// ParagXY is similar to Parag but is placed at the given X, Y
// position not to the current one.
func (fp *Fpdf) ParagXY(x, y, width, height float64, textStr, alignStr string) {
	// Backup current position
	x_orig, y_orig := fp.GetXY()

	// Set new position
	fp.SetXY(x, y)

	// Create Paragraph
	fp.Parag(width, height, textStr, alignStr)

	// Reset Position
	fp.SetXY(x_orig, y_orig)
}

// ParagX is similar to Parag but is placed at the given
// X position not to the current one.
func (fp *Fpdf) ParagX(x, width, height float64, textStr string, alignStr string) {
	// Backup current position
	x_orig := fp.GetX()

	// Set new position
	fp.SetX(x)

	// Create Paragraph
	fp.Parag(width, height, textStr, alignStr)

	// Reset Position
	fp.SetX(x_orig)
}

// ParagXCenter is similar to Parag but is horizontaly centered in the page.
func (fp *Fpdf) ParagXCenter(width, height float64, textStr string, alignStr string) {
	// Create Paragraph
	fp.ParagX(fp.workingCenter-width/2, width, height, textStr, alignStr)
}

// ParagXRight is similar to Parag but is horizontaly position at right
// in the page.
func (fp *Fpdf) ParagXRight(width, height float64, textStr string, alignStr string) {
	// Create Paragraph
	fp.ParagX(fp.workingRight-width, width, height, textStr, alignStr)
}

// SetHeader sets the header simplify with given left,
// center and right text.
func (fp *Fpdf) SetHeader(left, center, right string) {
	// Update top margin for header
	fp.SetTopMargin(fp.tMargin * 2.0)

	fp.SetHeaderFunc(func() {
		// Backup font color
		font_color_orig := NewColor().FromTextColor(fp)
		// Set font color
		NewRGBColor(150, 150, 150).ToTextColor(fp)
		ptsize, _ := fp.GetFontSize()
		fp.SetFontSize(ptsize * 0.8)

		y_pos := fp.tMargin / 2

		// Left text
		fp.Text(fp.lMargin, y_pos, left)

		// Center text
		rc := fp.GetStringWidth(center)
		fp.Text(fp.workingWidth/2+fp.lMargin-rc/2, y_pos, center)

		// Right text
		rw := fp.GetStringWidth(right)
		fp.Text(fp.workingWidth+fp.lMargin-rw, y_pos, right)

		// Restore font size and color
		fp.SetFontSize(ptsize)
		font_color_orig.ToTextColor(fp)
	})
}

// SetFooter sets the footer simplify with given left,
// center and right text.
func (fp *Fpdf) SetFooter(left, center, right string) {
	// Update bottom margin for footer
	if fp.bMargin < 20.0 {
		fp.SetAutoPageBreak(true, 20.0)
	}

	fp.SetFooterFunc(func() {
		// Backup font color
		font_color_orig := NewColor().FromTextColor(fp)
		// Set font color
		NewRGBColor(150, 150, 150).ToTextColor(fp)
		ptsize, _ := fp.GetFontSize()
		fp.SetFontSize(ptsize * 0.8)

		y_pos := fp.workingHeight + fp.tMargin + fp.bMargin/2

		// Left text
		fp.Text(fp.lMargin, y_pos, left)

		// Center text
		rc := fp.GetStringWidth(center)
		fp.Text(fp.workingWidth/2+fp.lMargin-rc/2, y_pos, center)

		// Right text
		rw := fp.GetStringWidth(right)
		fp.Text(fp.workingWidth+fp.lMargin-rw, y_pos, right)

		// Restore font size and color
		fp.SetFontSize(ptsize)
		font_color_orig.ToTextColor(fp)
	})
}

// SetFooterWithPageNumber sets the footer simplify with given left,
// right text. Page number is automatically add at bottom center.
func (fp *Fpdf) SetFooterWithPageNumber(left, right string) {
	// Update bottom margin for footer
	if fp.bMargin < 20.0 {
		fp.SetAutoPageBreak(true, 20.0)
	}

	fp.SetFooterFunc(func() {
		// Backup font color
		font_color_orig := NewColor().FromTextColor(fp)
		// Set font color
		NewRGBColor(150, 150, 150).ToTextColor(fp)
		ptsize, _ := fp.GetFontSize()
		fp.SetFontSize(ptsize * 0.8)

		y_pos := fp.workingHeight + fp.tMargin + fp.bMargin/2

		// Left text
		fp.Text(fp.lMargin, y_pos, left)

		// Set page number text
		page_text := fmt.Sprintf("%d/{nb}", fp.PageNo())
		rc := fp.GetStringWidth(page_text)
		fp.Text(fp.workingWidth/2+fp.lMargin-rc/2, y_pos, page_text)

		// Right text
		rw := fp.GetStringWidth(right)
		fp.Text(fp.workingWidth+fp.lMargin-rw, y_pos, right)

		// Restore font size and color
		fp.SetFontSize(ptsize)
		font_color_orig.ToTextColor(fp)
	})
}

// BulletedListXY insert a bullet list at the specified
// position with the specified list of string.
func (fp *Fpdf) BulletedList(height float64, list []string, lvl int) {
	if height < fp.fontSize {
		height = fp.fontSize * 1.2
	}
	x, y := fp.GetXY()
	fp.BulletedListXY(x, y, height, list, lvl)
	fp.SetXY(x, y+float64(len(list))*height)
}

// BulletedListXY insert a bullet list at the specified
// position with the specified list of string.
func (fp *Fpdf) BulletedListXY(x, y, height float64, list []string, lvl int) {
	if height < fp.fontSize {
		height = fp.fontSize * 1.2
	}
	radius := 0.5
	for n, str := range list {
		fp.Line(x, y, x+5, y)
		fp.Line(x, y, x, y+5)
		fp.Bullet(x+5*float64(lvl+1), y+float64(n)*height-radius-fp.fontSize*0.1, radius, lvl)
		fp.Text(x+5*float64(lvl+1)+4*radius, y+float64(n)*height, str)
	}
}

func (fp *Fpdf) Bullet(x, y, size float64, lvl int) {
	switch lvl {
	case 0:
		fp.Circle(x, y, size, "D")
	case 1:
		fp.Circle(x, y, size, "FD")
	case 2:
		poly := []PointType{PointType{x, y}, PointType{x + size, y}, PointType{x + size, y + size}, PointType{x, y + size}}
		fp.Polygon(poly, "D")
	case 3:
		poly := []PointType{PointType{x, y}, PointType{x + size, y}, PointType{x + size, y + size}, PointType{x, y + size}}
		fp.Polygon(poly, "FD")
	}
}
