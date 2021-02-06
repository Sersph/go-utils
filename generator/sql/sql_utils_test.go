package sql

import (
	"fmt"
	"testing"
)

type User struct {
}

func TestIdUtils(t *testing.T) {
	sqlBuilder := NewSelectBuilder()

	sqlBuilder.From(User{}).Eq("name", "winily").Eq("age", 20).Eq("sex", 1)

	fmt.Println(sqlBuilder.ToSQl())
}
