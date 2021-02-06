package arrays

//切片求集：交集、差集、并集

//交集
func BeMixed(slice, slice2 []string) []string {
	if len(slice) <= 0 || len(slice) <= 0 {
		return nil
	}

	newSlice := make([]string, 0)
	for _, value := range slice {
		exits := false
		for _, value2 := range slice2 {
			if value == value2 {
				exits = true
				break
			}
		}
		if exits {
			newSlice = append(newSlice, value)
		}
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

//并集 去重
func Aggregate(slice []string) []string {
	if len(slice) <= 1 {
		return slice
	}

	newSlice := make([]string, 0)
	for _, value := range slice {
		exits := true
		for _, newValue := range newSlice {
			if value == newValue {
				exits = false
			}
		}
		if exits {
			newSlice = append(newSlice, value)
		}
	}
	return newSlice
}
