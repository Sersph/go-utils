package generator

import (
	"github.com/sony/sonyflake"
	"strconv"
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
