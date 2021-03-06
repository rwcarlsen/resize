/*
Copyright (c) 2012, Jan Schlicht <jan.schlicht@gmail.com>

Permission to use, copy, modify, and/or distribute this software for any purpose
with or without fee is hereby granted, provided that the above copyright notice
and this permission notice appear in all copies.

THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES WITH
REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF MERCHANTABILITY AND
FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR ANY SPECIAL, DIRECT,
INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES WHATSOEVER RESULTING FROM LOSS
OF USE, DATA OR PROFITS, WHETHER IN AN ACTION OF CONTRACT, NEGLIGENCE OR OTHER
TORTIOUS ACTION, ARISING OUT OF OR IN CONNECTION WITH THE USE OR PERFORMANCE OF
THIS SOFTWARE.
*/

package resize

import (
	"image"
	"image/color"
)

type colorArray [4]float32

// converter allows to retrieve
// a colorArray for points of an image
type converter interface {
	at(x, y int) colorArray
}

type genericConverter struct {
	src image.Image
}

func (c *genericConverter) at(x, y int) colorArray {
	r, g, b, a := c.src.At(x, y).RGBA()
	return colorArray{
		float32(r),
		float32(g),
		float32(b),
		float32(a),
	}
}

type rgbaConverter struct {
	src *image.RGBA
}

func (c *rgbaConverter) at(x, y int) colorArray {
	if !(image.Point{x, y}.In(c.src.Rect)) {
		return colorArray{0, 0, 0, 0}
	}
	i := c.src.PixOffset(x, y)
	return colorArray{
		float32(uint16(c.src.Pix[i+0])<<8 | uint16(c.src.Pix[i+0])),
		float32(uint16(c.src.Pix[i+1])<<8 | uint16(c.src.Pix[i+1])),
		float32(uint16(c.src.Pix[i+2])<<8 | uint16(c.src.Pix[i+2])),
		float32(uint16(c.src.Pix[i+3])<<8 | uint16(c.src.Pix[i+3])),
	}
}

type rgba64Converter struct {
	src *image.RGBA64
}

func (c *rgba64Converter) at(x, y int) colorArray {
	if !(image.Point{x, y}.In(c.src.Rect)) {
		return colorArray{0, 0, 0, 0}
	}
	i := c.src.PixOffset(x, y)
	return colorArray{
		float32(uint16(c.src.Pix[i+0])<<8 | uint16(c.src.Pix[i+1])),
		float32(uint16(c.src.Pix[i+2])<<8 | uint16(c.src.Pix[i+3])),
		float32(uint16(c.src.Pix[i+4])<<8 | uint16(c.src.Pix[i+5])),
		float32(uint16(c.src.Pix[i+6])<<8 | uint16(c.src.Pix[i+7])),
	}
}

type grayConverter struct {
	src *image.Gray
}

func (c *grayConverter) at(x, y int) colorArray {
	if !(image.Point{x, y}.In(c.src.Rect)) {
		return colorArray{0, 0, 0, 0}
	}
	i := c.src.PixOffset(x, y)
	g := float32(uint16(c.src.Pix[i])<<8 | uint16(c.src.Pix[i]))
	return colorArray{
		g,
		g,
		g,
		float32(0xffff),
	}
}

type gray16Converter struct {
	src *image.Gray16
}

func (c *gray16Converter) at(x, y int) colorArray {
	if !(image.Point{x, y}.In(c.src.Rect)) {
		return colorArray{0, 0, 0, 0}
	}
	i := c.src.PixOffset(x, y)
	g := float32(uint16(c.src.Pix[i+0])<<8 | uint16(c.src.Pix[i+1]))
	return colorArray{
		g,
		g,
		g,
		float32(0xffff),
	}
}

type ycbcrConverter struct {
	src *image.YCbCr
}

func (c *ycbcrConverter) at(x, y int) colorArray {
	if !(image.Point{x, y}.In(c.src.Rect)) {
		return colorArray{0, 0, 0, 0}
	}
	yi := c.src.YOffset(x, y)
	ci := c.src.COffset(x, y)
	r, g, b := color.YCbCrToRGB(c.src.Y[yi], c.src.Cb[ci], c.src.Cr[ci])
	return colorArray{
		float32(uint16(r) * 0x101),
		float32(uint16(g) * 0x101),
		float32(uint16(b) * 0x101),
		float32(0xffff),
	}
}
