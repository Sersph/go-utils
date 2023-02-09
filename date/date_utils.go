package date

import (
	"github.com/winily/go-utils/common"
	"log"
	"strconv"
	"time"
)

// GetYTD 获取年月日 例：2020.08.28 = 200828
func GetYTD() string {
	t := time.Now()
	year := strconv.Itoa(t.Year() - common.Two_Thousand_Years)
	months := t.Format(common.Monthly)
	days := t.Format(common.Day)

	return year + months + days
}

/*
	获取时间
	参数：years 年、 months 月、 days 天
	列： 0为当时， 整数表示往后， 负数表示往前
	返回：2021-02-03
*/
func GetAddTime(years, months, days int) string {
	return time.Now().AddDate(years, months, days).Format(common.Date)
}

// GetTimeHourDiffer 计算小时相差几小时
func GetTimeHourDiffer(t2, t1 time.Time) int64 {
	var hour int64
	if t1.Before(t2) {
		diff := t2.Unix() - t1.Unix()
		hour = diff / 3600
		return hour
	} else {
		return hour
	}
}

// GetTimeMinDiffer 计算分钟相差几小时
func GetTimeMinDiffer(t2, t1 time.Time) int64 {
	var min int64
	if t1.Before(t2) {
		diff := t2.Unix() - t1.Unix()
		min = diff / (3600 * 60)
		return min
	} else {
		return min
	}
}

//TimeParseString 2006-10-02 10:00:00 转time
func TimeParseString(data string) *time.Time {
	times, err := time.Parse(common.Time, data)
	if err != nil {
		log.Fatalln("time parse err:", err)
	}
	return &times
}

// GetWeekTimeByTime 通过日期字符串 获取本周的开始时间 和结束时间
func GetWeekTimeByTime(timeStr string) (startDay, endDay time.Time) {

	if timeStr == "" {
		return
	}

	loc, _ := time.LoadLocation("Local")
	rTime, _ := time.ParseInLocation(common.Date, timeStr, loc)

	end := (time.Saturday + 1 - rTime.Weekday()).String()

	owe := int64(86400 * (7 - common.WeekDay[end]))
	plus := int64(86400 * (common.WeekDay[end] - 1))
	d, _ := time.ParseDuration("24h")

	startDay = time.Unix(rTime.Unix()-owe, 0).Add(d)
	endDay = time.Unix(rTime.Unix()+plus, 0).Add(d)
	return
}

// GetMonthTimeByTime 获取指定日期的开始月时间和结束月时间
func GetMonthTimeByTime(timeStr string) (startDay, endDay time.Time) {

	loc, _ := time.LoadLocation(common.Local)
	rTime, _ := time.ParseInLocation(common.Date, timeStr, loc)

	newMonth := rTime.Month()
	startDay = time.Date(rTime.Year(), newMonth, 1, 0, 0, 0, 0, loc)
	endDay = time.Date(rTime.Year(), newMonth+1, 0, 0, 0, 0, 0, loc)

	return
}

// GetBetweenDates 根据开始日期和结束日期计算出时间段内所有日期
// 参数为日期格式，如：2020-01-01
func GetBetweenDates(startDay, endDay string) []string {
	d := make([]string, 0)
	timeFormatTpl := common.Time
	if len(timeFormatTpl) != len(startDay) {
		timeFormatTpl = timeFormatTpl[0:len(startDay)]
	}
	date, err := time.Parse(timeFormatTpl, startDay)
	if err != nil {
		log.Println("GetBetweenDates 时间解析:", err)
		return d
	}
	date2, err := time.Parse(timeFormatTpl, endDay)
	if err != nil {
		log.Println("GetBetweenDates 时间解析", err)
		return d
	}
	if date2.Before(date) {
		log.Println("GetBetweenDates 结束时间小于开始时间")
		return d
	}
	// 输出日期格式固定
	timeFormatTpl = common.Date
	date2Str := date2.Format(timeFormatTpl)
	d = append(d, date.Format(timeFormatTpl))
	for {
		date = date.AddDate(0, 0, 1)
		dateStr := date.Format(timeFormatTpl)
		d = append(d, dateStr)
		if dateStr == date2Str {
			break
		}
	}
	return d
}

func StrToTime(timeStr string) time.Time {
	timeStr = timeStr + " 00:00:00"
	local, _ := time.LoadLocation(common.Local)
	t, _ := time.ParseInLocation(common.Time, timeStr, local)
	return t
}
