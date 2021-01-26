package generator

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

//获取年月日 例：2020.08.28 = 200828
func GetYTD() string {
	year := strconv.Itoa(time.Now().Year() - 2000)
	months := time.Now().Format("01")
	days := time.Now().Format("02")

	k := year + months + days
	return k
}

//获取流水号：000001-999999
func generatorFourCode(index int) (int, string) {
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

var (
	exits = 1
	front string
)

//生成单号：年月日+6位流水号 列：201029000001
func SteamID() string {
	//判断是不是 新的一天
	if front != time.Now().Format("2006-01-02") {
		exits = 1
	}

	_, days := generatorFourCode(exits)
	id := GetYTD() + days
	//记录 当前日期
	front = time.Now().Format("2006-01-02")
	exits += 1
	return id
}

//指定位数的随机数 <7（允许批量生成）
func GetRandomInteger(bit int) string {
	var result string
	time.Sleep(1)
	switch bit {
	case 1:
		result = fmt.Sprintf("%01v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(10))
	case 2:
		result = fmt.Sprintf("%02v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(100))
	case 3:
		result = fmt.Sprintf("%03v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000))
	case 4:
		result = fmt.Sprintf("%04v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(10000))
	case 5:
		result = fmt.Sprintf("%05v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(100000))
	case 6:
		result = fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
	default:
		result = ""
	}
	return result
}
