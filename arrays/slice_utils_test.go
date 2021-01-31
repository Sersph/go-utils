package arrays

import (
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {
	arr := make([]int, 0)
	arr = append(arr, []int{1, 2, 3, 4, 5, 6, 7}...)
	fmt.Println("stats: ", arr)

	arrs := make([]interface{}, 0)

	for _, value := range arr {
		arrs = append(arrs, value)
	}

	//arrs = SliceAdd(arrs, 8, 0)

	//arrs  = SliceAdds(arrs, []interface{}{11,22,33}, 6)

	var err error
	arrs, err = SliceDel(arrs, 2, 5)
	fmt.Println("err: ", err)

	fmt.Println("end: ", arrs)
}
