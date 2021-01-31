package arrays

import "errors"

func lengthPass(slice []interface{}, index int) bool {
	if len(slice) <= index {
		//print("Array out of bounds")
		return true
	}
	return false

}

//slice : 目标切片、data : 需要添加的数(一个)、index : 添加到指定的下标
func SliceAdd(slice []interface{}, data interface{}, index int) ([]interface{}, error) {
	if lengthPass(slice, index) {
		return nil, errors.New("array out of bounds")
	}

	slice = append(slice, 0)

	copy(slice[index+1:], slice[index:])
	slice[index] = data

	return slice, nil
}

//slice : 目标切片、data : 需要添加的数(多个)、index : 添加到指定的下标
func SliceAdds(slice []interface{}, array []interface{}, index int) ([]interface{}, error) {
	if lengthPass(slice, index) {
		return nil, errors.New("array out of bounds")
	}
	slice = append(slice, array...) //扩展足够多的位置

	copy(slice[index+len(array):], slice[index:]) //slice[index:] 向后移动len(array)个位置
	copy(slice[index:], array)                    //复制新添加的切片
	return slice, nil
}

//slice : 目标切片、
func SliceDel(slice []interface{}, index, quantity int) ([]interface{}, error) {
	if lengthPass(slice, index) {
		return nil, errors.New("array out of bounds")
	}
	if len(slice) < quantity+index {
		return nil, errors.New("array out of bounds")
	}
	slice = slice[:index+copy(slice[index:], slice[index+quantity:])]

	return slice, nil
}
