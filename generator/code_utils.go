package generator

import "strings"

// 获取流水号：000001-999999
func SerialCode(index int) (int, string) {
	var c = [...]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	i1 := (index / (10 * 10 * 10 * 10 * 10)) % 10
	i2 := (index / (10 * 10 * 10 * 10)) % 10
	i3 := (index / (10 * 10 * 10)) % 10
	i4 := (index / (10 * 10)) % 10
	i5 := (index / (10)) % 10
	i6 := (index) % 10

	var str strings.Builder
	str.WriteString(c[i1])
	str.WriteString(c[i2])
	str.WriteString(c[i3])
	str.WriteString(c[i4])
	str.WriteString(c[i5])
	str.WriteString(c[i6])
	return index, str.String()
}
