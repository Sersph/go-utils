package conversion

import (
	"fmt"
	"testing"
)

func TestIdUtils(t *testing.T) {
	c := Color16BitToRGBA("#FF708AF0")
	fmt.Println(c.R, c.G, c.B, c.A)
	b := RGBAToColor16Bit(c)
	fmt.Println(b)
}
