package arrays

import (
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {
	arr := make([]int, 0)
	arr = append(arr, []int{1, 2, 3, 4, 5, 6, 7, 7, 6, 8, 9, 9, 4, 3}...)
	fmt.Println("stats: ", arr)

	strArrs := []string{"bbb"}
	strArr := []string{"aaa", "bbb", "ccc", "eee", "ddd", "aaa", "ccc"}

	//arrs := make([]interface{}, 0)
	//
	//for _, value := range arr {
	//	arrs = append(arrs, value)
	//}

	//arrs = SliceAdd(arrs, 8, 0)

	//arrs  = SliceAdds(arrs, []interface{}{11,22,33}, 6)

	//var err error
	//arrs, err = SliceDel(arrs, 2, 5)
	//fmt.Println("err: ", err)
	//
	//fmt.Println("end: ", arrs)

	//fmt.Println("-:", conversion.ToString(arr))

	//fmt.Println(7/2)

	fmt.Println("并集", Aggregate(strArr))
	fmt.Println("交集", BeMixed(strArr, strArrs))
	fmt.Println("差集", DifferenceSet(strArr))

}
