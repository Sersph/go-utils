package generator

import (
	"fmt"
	"math/rand"
	"time"
)

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

