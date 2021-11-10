package common

const (
	Monthly = "01"
	Day     = "02"
	Date    = "2006-01-02"
	Time    = "2006-01-02 15:04:05"

	Local = "Local"

	Two_Thousand_Years = 2000
)

var (
	MonthDay = map[string]int{
		"January":   1,
		"February":  2,
		"March":     3,
		"April":     4,
		"May":       5,
		"June":      6,
		"July":      7,
		"August":    8,
		"September": 9,
		"October":   10,
		"November":  11,
		"December":  12,
	}

	WeekDay = map[string]int{
		"Monday":    1,
		"Tuesday":   2,
		"Wednesday": 3,
		"Thursday":  4,
		"Friday":    5,
		"Saturday":  6, // 周六
		"Sunday":    7, // 周日
	}
)
