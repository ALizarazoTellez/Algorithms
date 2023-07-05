package color

import (
	"fmt"
	"strconv"
)

type RGB struct {
	R, G, B uint8
}

func (rgb RGB) String() string {
	return fmt.Sprintf("rgb(%d, %d, %d)", rgb.R, rgb.G, rgb.B)
}

func (rgb RGB) ToHex() Hex {
	return Hex{fmt.Sprintf("%X%X%X", rgb.R, rgb.G, rgb.B)}
}

type Hex struct {
	Hex string
}

func (h Hex) ToRGB() RGB {
	var rgb RGB

	r, _ := strconv.ParseUint(h.Hex[0:2], 16, 8)
	rgb.R = uint8(r)

	g, _ := strconv.ParseUint(h.Hex[2:4], 16, 8)
	rgb.G = uint8(g)

	b, _ := strconv.ParseUint(h.Hex[4:6], 16, 8)
	rgb.B = uint8(b)

	return rgb
}
