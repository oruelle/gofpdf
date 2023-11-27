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

type BatchInterface interface {
	Insert() error
}

type BatchParag struct {
	BatchInterface
	text       string  // text to insert
	x          float64 // y position to insert element
	y          float64 // y position to insert element
	width      float64 // width of paragraph
	lheight    float64 // line height
	fontFamily string
	fontStyle  string
	fontSizePt int    // font size in points
	fontColor  Color  // font color
	align      string // alignement to set
}

func (p *BatchParag) Insert(fp *Fpdf) error {
	fp.ParagXY(p.x, p.y, p.width, p.lheight, p.text, p.align)
	if fp.Error() != nil {
		return fp.Error()
	}
	return nil
}
