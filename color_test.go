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

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestNewHexColor(t *testing.T) {
	color := "bbef15"
	c, err := NewHexColor(color)
	if err != nil {
		t.Fatalf(`NewHexColor(%s) = %v`, color, err)
	}
	if c.R() != 0xbb || c.G() != 0xef || c.B() != 0x15 {
		t.Fatalf(`NewHexColor(%s) = %d,%d,%d, %v, want match for %d, %d, %d, nil`, color, c.R(), c.G(), c.B(), err, 0xbb, 0xef, 0x15)
	}

	color = "0x123456"
	c, err = NewHexColor(color)
	if err != nil {
		t.Fatalf(`NewHexColor(%s) = %v`, color, err)
	}
	if c.R() != 0x12 || c.G() != 0x34 || c.B() != 0x56 {
		t.Fatalf(`NewHexColor(%s) = %d,%d,%d, %v, want match for %d, %d, %d, nil`, color, c.R(), c.G(), c.B(), err, 0x12, 0x34, 0x56)
	}
}
