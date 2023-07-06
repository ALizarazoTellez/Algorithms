package color

import (
	"fmt"
	"math"
	"strconv"
)

type RGB struct {
	R, G, B uint8
}

func (rgb RGB) String() string {
	return fmt.Sprintf("rgb(%d, %d, %d)", rgb.R, rgb.G, rgb.B)
}

func (rgb RGB) Hex() RGBHex {
	return RGBHex{fmt.Sprintf("%X%X%X", rgb.R, rgb.G, rgb.B)}
}

func (rgb RGB) HSL() HSL {
	r := float32(rgb.R) / 255.0
	g := float32(rgb.G) / 255.0
	b := float32(rgb.B) / 255.0

	var cmax float32

	if r > g && r > b {
		cmax = r
	} else if g > r && g > b {
		cmax = g
	} else {
		cmax = b
	}

	var cmin float32

	if r < g && r < b {
		cmin = r
	} else if g < r && g < b {
		cmin = g
	} else {
		cmin = b
	}

	delta := cmax - cmin

	// Get Hue.
	var hue uint16

	if delta == 0 {
		hue = 0
	} else if cmax == r {
		hue = uint16(60*math.Mod(float64((g-b)/delta), 6) + 0.5)
	} else if cmax == g {
		hue = uint16(60*(((b-r)/delta)+2) + 0.5)
	} else if cmax == b {
		hue = uint16(60*(((r-g)/delta)+4) + 0.5)
	}

	// Get Lightness.
	lightness := (cmax + cmin) / 2

	// Get Saturation.
	var sat float32

	if delta == 0 {
		sat = 0
	} else {
		sat = delta / float32(1-math.Abs(float64(2*lightness-1)))
	}

	return HSL{hue, sat * 100, lightness * 100}.Normalize()
}

func (rgb RGB) Mix(c RGB) RGB {
	r1 := float32(rgb.R)
	r2 := float32(c.R)
	g1 := float32(rgb.G)
	g2 := float32(c.G)
	b1 := float32(rgb.B)
	b2 := float32(c.B)

	return RGB{
		uint8((r1+r2)/2 + 0.5),
		uint8((g1+g2)/2 + 0.5),
		uint8((b1+b2)/2 + 0.5),
	}
}

type RGBHex struct {
	Hex string
}

func (h RGBHex) String() string {
	return "#" + h.Hex
}

func (h RGBHex) RGB() RGB {
	var rgb RGB

	r, _ := strconv.ParseUint(h.Hex[0:2], 16, 8)
	rgb.R = uint8(r)

	g, _ := strconv.ParseUint(h.Hex[2:4], 16, 8)
	rgb.G = uint8(g)

	b, _ := strconv.ParseUint(h.Hex[4:6], 16, 8)
	rgb.B = uint8(b)

	return rgb
}

type HSL struct {
	H uint16
	S float32
	L float32
}

func (hsl HSL) String() string {
	return fmt.Sprintf("hsl(%d, %.1f, %.1f)", hsl.H, hsl.S, hsl.L)
}

func (hsl HSL) RGB() RGB {
	s := hsl.S / 100
	l := hsl.L / 100
	h := hsl.H / 60

	c := float32((1 - math.Abs(float64((2*l - 1)))) * float64(s))
	x := (c * float32(1-math.Abs(float64((h%2)-1))))
	m := l - c/2

	var r, g, b float32

	switch {
	case hsl.H < 60:
		r = c
		g = x
		b = 0
	case hsl.H < 120:
		r = x
		g = c
		b = 0
	case hsl.H < 180:
		r = 0
		g = c
		b = x
	case hsl.H < 240:
		r = 0
		g = x
		b = c
	case hsl.H < 300:
		r = x
		g = 0
		b = c
	default:
		r = c
		g = 0
		b = x
	}

	r = (r+m)*255 + 0.5
	g = (g+m)*255 + 0.5
	b = (b+m)*255 + 0.5

	return RGB{
		uint8(r),
		uint8(g),
		uint8(b),
	}
}

func (hsl HSL) Normalize() HSL {
	return HSL{
		hsl.H,
		float32(math.Round(float64(hsl.S*10)) / 10),
		float32(math.Round(float64(hsl.L*10)) / 10),
	}
}

func ConvertToPastel(rgb RGB) RGB {
	hsl := rgb.HSL()
	hsl.S += 10

	rgb = hsl.RGB()

	return rgb.Mix(RGB{255, 255, 255})
}
