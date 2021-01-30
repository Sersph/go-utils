package conversion

import (
	"fmt"
	"testing"
)

func TestToMap(t *testing.T) {
	str := struct {
		name string
		age  int
	}{
		name: "kay",
		age:  19,
	}

	fmt.Println("str: ", str)

	maps := ToMap(str)

	fmt.Println(maps)
}
