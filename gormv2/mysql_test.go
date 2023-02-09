package gormv2

import (
	"fmt"
	"testing"
)

func Test_mysql(t *testing.T) {
	InitMysql("root:123456tcp@(localhost:3306)")

	var count int64
	GetMainDB().Table("test.user").Count(&count)

	fmt.Println("count:", count)
}
