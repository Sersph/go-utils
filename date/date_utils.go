package date

import (
	"log"
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

/*
	获取时间
	参数：years 年、 months 月、 days 天
	列： 0为当时， 整数表示往后， 负数表示往前
	返回：2021-02-03
*/
func GetTime(years, months, days int) string {
	return time.Now().AddDate(years, months, days).Format("2006-01-02")
}

//string 2006-10-02 10:00:00 转time
func TimeParseString(data string) *time.Time {
	times, err := time.Parse("2006-01-02 15:04:05", data)
	if err != nil {
		log.Fatalln("time parse err:", err)
	}
	return &times
}
