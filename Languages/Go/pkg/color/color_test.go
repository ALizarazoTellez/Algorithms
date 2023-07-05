package color

import (
	"fmt"
	"testing"
)

func TestColorConvert(t *testing.T) {
	tests := []struct {
		rgb RGB
		hex Hex
	}{
		{RGB{26, 93, 26}, Hex{"1A5D1A"}},
		{RGB{241, 201, 59}, Hex{"F1C93B"}},
		{RGB{251, 216, 93}, Hex{"FBD85D"}},
		{RGB{250, 227, 146}, Hex{"FAE392"}},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("convert %v to %s", tt.rgb, tt.hex), func(t *testing.T) {
			got := tt.rgb.ToHex()
			if got != tt.hex {
				t.Errorf("want %q, got %q", tt.hex, got)
			}
		})

		t.Run(fmt.Sprintf("convert %s to %v", tt.hex, tt.rgb), func(t *testing.T) {
			got := tt.hex.ToRGB()
			if got != tt.rgb {
				t.Errorf("want %q, got %q", tt.rgb, got)
			}
		})
	}
}
