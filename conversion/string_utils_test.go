package conversion

import (
	"fmt"
	"testing"
)

func TestString(t *testing.T) {
	fmt.Println(HumpToUnderscore("AbcdEfgOZ"))
}
