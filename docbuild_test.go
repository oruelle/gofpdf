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
	"testing"
)

func TestDocBuild(t *testing.T) {
	pdf, err := New("P", "mm", "A4", "")
	if err != nil {
		t.Fatalf("Failed to create PDF: %v", err)
	}
	pdf.SetFooterWithPageNumber("Left text", "Right text")
	pdf.SetHeader("Left header", "center header", "Right header")

	err = pdf.SetFont("Arial", "", 10) // If header or footer add page after setting font
	if err != nil {
		t.Fatalf("Failed to create Set font: %v", err)
	}

	err = pdf.AddPage()
	if err != nil {
		t.Fatalf("Failed to add new page: %v", err)
	}

	pdf.SetCellMargin(5)

	ww, _ := pdf.GetWorkingSize()

	white := NewRGBColor(255, 255, 255)
	blue := NewRGBColor(141, 179, 226)

	// Create title
	pdf.Title("My title", 0, white, blue)

	pdf.Title("My title 1", 1, white, blue)
	pdf.SetFontSize(10)

	str := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aliquam pulvinar lectus mi, ac blandit sem finibus ultricies. Etiam ac nunc vitae velit efficitur finibus. Nulla sit amet pharetra massa. Quisque vehicula quis eros et elementum. Donec vitae erat vitae tellus facilisis condimentum. Aenean varius purus lacinia bibendum blandit. Maecenas aliquam venenatis mi, at condimentum velit eleifend a. Vivamus ut luctus nibh, eu dictum nulla. Fusce mattis lorem vel lectus vulputate, in bibendum diam volutpat. Quisque varius a nunc sed convallis. Nunc et sollicitudin urna. Integer egestas pulvinar nulla, sed auctor dolor. Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas. Nullam dolor velit, faucibus nec facilisis ut, dictum eget orci."

	//Paragraphs
	pdf.Parag(ww, 0, str, ALIGN_LEFT)

	pdf.Parag(ww/2, 0, str, ALIGN_LEFT)

	pdf.ParagXCenter(ww/2, 0, str, ALIGN_LEFT)

	pdf.ParagXRight(ww/2, 0, str, ALIGN_LEFT)

	pdf.Parag(ww, 0, str, ALIGN_JUSTIFY)

	pdf.Parag(ww/2, 0, str, ALIGN_JUSTIFY)

	pdf.ParagXCenter(ww/2, 0, str, ALIGN_JUSTIFY)

	pdf.ParagXRight(ww/2, 0, str, ALIGN_JUSTIFY)

	//Table
	h1 := []string{"head1", "head2", "head3", "head4", "head5"}
	r1 := []string{"row11", "row12", "row13", "row14", "row15"}
	r2 := []string{"row21", "row22", "row23", "row24", "long long long long cell"}
	r3 := []string{"row31", "row32", "row33", "row34", "row35"}
	_table := [][]string{h1, r1, r2, r3}
	_table_align := []string{"LM", "CM", "CM", "CM", "CM"}
	pdf.Table(ww/2, _table, _table_align, false, false)

	pdf.TableXCenter(ww/2, _table, _table_align, true, true)

	// Another Paragraph
	pdf.SetHomeX()
	pdf.Parag(ww, 0, str, ALIGN_JUSTIFY)

	pdf.Parag(ww, 0, str, ALIGN_JUSTIFY)

	pdf.ParagXCenter(ww/2, 0, str, ALIGN_JUSTIFY)

	pdf.BulletedList(0, r1, 0)
	pdf.BulletedList(0, r2, 1)
	pdf.NextLine(0)

	x, y := pdf.GetXY()
	err = pdf.ImageOptions("image/golang-gopher.png", x, y, 50, 50, false, ImageOptions{}, 0, "")
	if err != nil {
		t.Fatal(err)
	}

	pdf.Parag(ww, 0, str, ALIGN_JUSTIFY)

	err = pdf.OutputFileAndClose("pdf/Fpdf_DocBuildSimply.pdf")
	if err != nil {
		t.Fatalf("%v", err)
	}
}
