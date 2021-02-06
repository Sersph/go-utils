package sql

import "github.com/winily/go-utils/conversion"

/**
使用队列方式进行 sql 词组进行暂存
利用先进先出方式进行拼接 sql
*/
type SelectBuilder struct {
	// 查询字段队列
	SelectPhrase *Queue
	// 查询表队列
	FromPhrase *Queue
	// 查询条件队列
	WherePhrase *Queue
}

type ConditionFunc func(interface{}) bool

func NewSelectBuilder() *SelectBuilder {
	sb := SelectBuilder{}
	sb.SelectPhrase = NewQueue()
	sb.FromPhrase = NewQueue()
	sb.WherePhrase = NewQueue()
	return &sb
}

func (sb *SelectBuilder) Eq(column string, val interface{}) *SelectBuilder {
	where := ColumnProcess(column) + EQ + conversion.ToString(val)
	sb.WherePhrase.Enqueue(where)
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
