package conversion

import (
	"fmt"
	"testing"
)

func TestString(t *testing.T) {
	value := 9.888
	fmt.Println("toString: ", ToString(value))

	arrays := []interface{}{
		90, "toString", 9.88, []int{1, 2, 3}, map[string]interface{}{
			"name": "kay",
			"age":  18,
		},
	}

	fmt.Println("arrayToString: ", ArrayToString(arrays, " - "))

}
