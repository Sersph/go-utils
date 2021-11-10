package arrays

//切片求集：交集、差集、并集

//str 公共部分
func BeMixed(slice, slice2 []string) []string {
	if len(slice) <= 0 || len(slice) <= 0 {
		return nil
	}

	sliceMap := strToMap(slice)

	slice2Map := strToMap(slice2)

	newSlice := make([]string, 0)
	for key := range sliceMap {
		newSlice = append(newSlice, slice2Map[key])
	}

	return Aggregate(newSlice)
}

//差集
func DifferenceSet(slice []string) []string {
	if len(slice) <= 1 {
		return slice
	}

	newMap := make(map[string]int)
	for _, value := range slice {
		newMap[value]++
	}

	newSlice := make([]string, 0)
	for _, value := range slice {
		if newMap[value] == 1 {
			newSlice = append(newSlice, value)
		}
	}

	return newSlice
}

// str 去重
func Aggregate(slice []string) []string {
	if len(slice) <= 1 {
		return slice
	}

	exitMap := strToMap(slice)

	newSlice := make([]string, 0)
	for str := range exitMap {
		newSlice = append(newSlice, str)
	}

	return newSlice
}

func strToMap(slice []string) map[string]string {
	exitMap := make(map[string]string)
	for _, v := range slice {
		exitMap[v] = v
	}
	return exitMap
}
