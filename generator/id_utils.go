package generator

import (
	"github.com/sony/sonyflake"
	dateUtils "github.com/winily/go-utils/date"
	"strconv"
	"time"
)

func AmazonEC2MachineID() (uint16, error) {
	var c uint16
	c = 65
	return c, nil
}

// 采用sony的雪花算法生成主键id
func GetSonyId() string {
	var st sonyflake.Settings
	st.MachineID = AmazonEC2MachineID
	sf := sonyflake.NewSonyflake(st)
	if sf == nil {
		panic("sony flake not created")
	}
	id, err := sf.NextID()
	if err != nil {
		return ""
	}
	return strconv.FormatUint(id, 10)
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
	_, days := SerialCode(exits)
	id := dateUtils.GetYTD() + days
	//记录 当前日期
	front = time.Now().Format("2006-01-02")
	exits += 1
	return id
}