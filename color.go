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
	"strconv"
)

var (
	GREY, _         = NewHexColor("bdc3c7")
	LIGHT_GREY, _   = NewHexColor("e5e7e9")
	DARK_GREY, _    = NewHexColor("626567")
	BLACK, _        = NewHexColor("000000")
	WHITE, _        = NewHexColor("ffffff")
	BLUE, _         = NewHexColor("3498db")
	LIGHT_BLUE, _   = NewHexColor("aed6f1")
	GREEN, _        = NewHexColor("2ecc71")
	LIGHT_GREEN, _  = NewHexColor("abebc6")
	RED, _          = NewHexColor("e74c3c")
	LIGHT_RED, _    = NewHexColor("f5b7b1")
	YELLOW, _       = NewHexColor("f1c40f")
	LIGHT_YELLOW, _ = NewHexColor("f9e79f")
	ORANGE, _       = NewHexColor("e67e22")
	LIGHT_ORANGE, _ = NewHexColor("f5cba7")
)

type Color struct {
	r     int
	g     int
	b     int
	alpha int
}

// NewColor returns a Color object.
func NewColor() (c *Color) {
	c = new(Color)

	return
}

// NewRGBColor returns a Color object from given r, g, b values.
func NewRGBColor(r int, g int, b int) (c *Color) {
	c = new(Color)

	c.r = r
	c.b = b
	c.g = g

	return
}

// NewHexColor returns a Color object from a hex string representation
// or an error if the parsing failed.
func NewHexColor(hex string) (c *Color, err error) {
	c = new(Color)

	// Delete "0x" if present at the beginning of the string
	if hex[:2] == "0x" {
		hex = hex[2:]
	}

	value, err := strconv.ParseInt(hex, 16, 64)
	if err != nil {
		return nil, err
	}

	c.r = int(value>>16) & 0xff
	c.g = int(value>>8) & 0xff
	c.b = int(value) & 0xff

	return
}

// GetRGB returns values of red, green and blue components
// from the color
func (c *Color) GetRGB() (r, g, b int) {
	return c.r, c.g, c.b
}

// R() returns the red value of the color
func (c *Color) R() int {
	return c.r
}

// G() returns the green value of the color
func (c *Color) G() int {
	return c.g
}

// B() returns the blue value of the color
func (c *Color) B() int {
	return c.b
}

// ToFillColor() set Fill color with the color
// of the current color instance.
func (c *Color) ToFillColor(fp *Fpdf) {
	fp.SetFillColor(c.r, c.g, c.b)
}

// ToTextColor() set Text color with the color
// of the current color instance.
func (c *Color) ToTextColor(fp *Fpdf) {
	fp.SetTextColor(c.r, c.g, c.b)
}

// FromFillColor() set current color instance
// with the current Fill color.
func (c *Color) FromFillColor(fp *Fpdf) *Color {
	c.r, c.g, c.b = fp.GetFillColor()
	return c
}

// FromFillColor() set current color instance
// with the current Text color.
func (c *Color) FromTextColor(fp *Fpdf) *Color {
	c.r, c.g, c.b = fp.GetTextColor()
	return c
}
