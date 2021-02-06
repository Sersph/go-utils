package sql

/**
使用队列方式进行 sql 词组进行暂存
利用先进先出方式进行拼接 sql
*/
type SelectBuilder struct {
	// 查询字段队列
	Select *Select
	// 查询表队列
	From *From
	// 查询条件队列
	Where *Where
}

type ConditionFunc func(interface{}) bool

func NewSelectBuilder() *SelectBuilder {
	sb := SelectBuilder{}
	sb.Select = NewSelect()
	sb.From = NewFrom()
	sb.Where = NewWhere()
	return &sb
}

func (sb *SelectBuilder) ToSQl() string {
	if sb.From.IsNull() {
		panic("查询表不能为空")
	}
	return sb.Select.ToString() + sb.From.ToString() + sb.Where.ToString()
}
