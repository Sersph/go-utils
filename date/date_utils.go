package date


import (
	"strconv"
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