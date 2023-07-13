package color

import (
	"fmt"
	"testing"
)

func TestColorConvert(t *testing.T) {
	tests := []struct {
		rgb RGB
		hex RGBHex
		hsl HSL
	}{
		{RGB{26, 93, 26}, RGBHex{"1A5D1A"}, HSL{120, 56.3, 23.3}},
		{RGB{241, 201, 59}, RGBHex{"F1C93B"}, HSL{47, 86.7, 58.8}},
		{RGB{251, 216, 93}, RGBHex{"FBD85D"}, HSL{47, 95.2, 67.5}},
		{RGB{250, 227, 146}, RGBHex{"FAE392"}, HSL{47, 91.2, 77.6}},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("convert %q to %q", tt.rgb, tt.hex), func(t *testing.T) {
			got := tt.rgb.Hex()
			if got != tt.hex {
				t.Errorf("want %q, got %q", tt.hex, got)
			}
		})

		t.Run(fmt.Sprintf("convert %q to %q", tt.hex, tt.rgb), func(t *testing.T) {
			got := tt.hex.RGB()
			if got != tt.rgb {
				t.Errorf("want %q, got %q", tt.rgb, got)
			}
		})

		t.Run(fmt.Sprintf("convert %q to %q", tt.rgb, tt.hsl), func(t *testing.T) {
			got := tt.rgb.HSL()
			if got != tt.hsl {
				t.Errorf("want %q, got %q", tt.hsl, got)
			}
		})

		t.Run(fmt.Sprintf("convert %q to %q", tt.hsl, tt.rgb), func(t *testing.T) {
			got := tt.hsl.RGB()
			if got != tt.rgb {
				t.Errorf("want %q, got %q", tt.rgb, got)
			}
		})
	}
}

func TestColorMix(t *testing.T) {
	tests := []struct {
		color1 RGB
		color2 RGB
		mix    RGB
	}{
		{RGB{26, 93, 26}, RGB{0, 0, 0}, RGB{13, 47, 13}},
		{RGB{100, 24, 57}, RGB{255, 255, 255}, RGB{178, 140, 156}},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("mix %q with %q", tt.color1, tt.color2), func(t *testing.T) {
			got := tt.color1.Mix(tt.color2)
			if got != tt.mix {
				t.Errorf("want %q, got %q", tt.mix, got)
			}
		})
	}
}

func TestConvertToPastel(t *testing.T) {
	tests := []struct {
		color  RGB
		pastel RGB
	}{
		{RGB{26, 93, 26}, RGB{178, 140, 156}},
		{RGB{106, 67, 11}, RGB{184, 162, 131}},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("convert to pastel %q", tt.color), func(t *testing.T) {
			got := ConvertToPastel(tt.color)
			if tt.pastel != got {
				t.Errorf("want %q got %q", tt.pastel, got)
			}
		})
	}
}
