package conversion

import "strconv"

// 将数字转为科学计数法
// precision 精度 保留几位小数点 四舍五入
// in 120000
// out 1.2e+5
func ToScientificNotation(number string, precision int) string {
	// TODO
	return ""
}

// 将科学计数法的数字转为标准数字
// precision 精度 保留几位小数点 四舍五入
// in 1.2e+5
// out 120000
func ToNormal(number string, precision int) string {
	// TODO
	return ""
}

// 格式化数字位
// 如果位数等于要格式的位就直接返回
// 如果超过位，将头部超出位删除
// 如果不足，前面补零
func BitFormatInt(number int, bit int) string {
	numberStr := strconv.Itoa(number)
	if len(numberStr) == bit {
		return numberStr
	}

	if len(numberStr) > bit {
		return numberStr[len(numberStr) - bit:]
	}

	zero := ""
	for i := 0; i < (bit - len(numberStr)); i++ {
		zero += "0"
	}
	return zero + numberStr
}