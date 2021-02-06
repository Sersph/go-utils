package sql

import (
	"fmt"
	"testing"
)

type User struct {
}

func TestIdUtils(t *testing.T) {
	sqlBuilder := NewSelectBuilder()

	sqlBuilder.From.TableByEntity(User{})
	sqlBuilder.Where.Eq("Name", "winily")
	sqlBuilder.Where.IsEq(func(value interface{}) bool {
		return value != nil
	}, "Age", 0)
	sqlBuilder.Where.IsEq(func(value interface{}) bool {
		return value != nil
	}, "Sex", nil)

	fmt.Println(sqlBuilder.ToSQl())
}
