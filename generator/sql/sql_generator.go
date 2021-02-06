package sql

import (
	"github.com/winily/go-utils/conversion"
	"reflect"
)

/**
使用队列方式进行 sql 词组进行暂存
利用先进先出方式进行拼接 sql
*/
type SelectBuilder struct {
	// 查询字段队列
	selectPhrase *Queue
	// 查询表队列
	fromPhrase *Queue
	// 查询条件队列
	wherePhrase *Queue
}

type ConditionFunc func(interface{}) bool

func NewSelectBuilder() *SelectBuilder {
	sb := SelectBuilder{}
	sb.selectPhrase = NewQueue()
	sb.fromPhrase = NewQueue()
	sb.wherePhrase = NewQueue()
	return &sb
}

func (sb *SelectBuilder) From(table interface{}) *SelectBuilder {
	tableName := reflect.TypeOf(table).Name()
	sb.fromPhrase.Enqueue(tableName)
	return sb
}

func (sb *SelectBuilder) FromByName(table interface{}, tableName string) *SelectBuilder {
	sb.fromPhrase.Enqueue(tableName)
	return sb
}

func (sb *SelectBuilder) Eq(column string, val interface{}) *SelectBuilder {
	where := ConstituteWherePhrase(column, EQ, conversion.ToString(val))
	sb.wherePhrase.Enqueue(where)
	return sb
}

func (sb *SelectBuilder) IsEq(condition ConditionFunc, column string, val interface{}) *SelectBuilder {
	if condition(val) {
		sb.Eq(column, val)
	}
	return sb
}

func (sb *SelectBuilder) AllEq(params map[string]interface{}) *SelectBuilder {
	for column, val := range params {
		sb.Eq(column, val)
	}
	return sb
}

func (sb *SelectBuilder) IsAllEq(condition ConditionFunc, params map[string]interface{}) *SelectBuilder {
	for column, val := range params {
		sb.IsEq(condition, column, val)
	}
	return sb
}

func (sb *SelectBuilder) ToSQl() string {
	if sb.fromPhrase.Size() == 0 {
		panic("查询表不能为空")
	}
	return sb.ToSelect() + sb.ToFrom() + sb.ToWhere()
}

func (sb *SelectBuilder) ToSelect() string {
	if sb.selectPhrase.Size() == 0 {
		return SELECT + " * "
	}

	selectStr := SELECT
	for sb.selectPhrase.Size() > 0 {
		selectStr += " " + sb.selectPhrase.Dequeue() + " "
	}
	return selectStr
}

func (sb *SelectBuilder) ToWhere() string {
	if sb.wherePhrase.Size() == 0 {
		return ""
	}

	where := WHERE
	for sb.wherePhrase.Size() > 0 {
		where += " " + sb.wherePhrase.Dequeue() + " "
		if sb.wherePhrase.Size() >= 1 {
			where += AND
		}
	}
	return where
}

func (sb *SelectBuilder) ToFrom() string {
	from := FROM
	for sb.fromPhrase.Size() > 0 {
		from += " " + sb.fromPhrase.Dequeue() + " "
		if sb.fromPhrase.Size() >= 1 {
			from += ","
		}
	}
	return from
}
