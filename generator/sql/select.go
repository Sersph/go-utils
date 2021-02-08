package sql

type Select struct {
	sel *Queue
}

func NewSelect() *Select {
	sel := Select{}
	sel.sel = NewQueue()
	return &sel
}

/**
TODO 字段，设置查询字段，结构体解析字段
TODO 解析字段问题，配置转换字段名称 全小写，全大写或者不变
*/

func (sel *Select) ToString() string {
	if sel.sel.Size() == 0 {
		return SELECT + " * "
	}

	selectStr := SELECT
	for sel.sel.Size() > 0 {
		selectStr += " " + sel.sel.Dequeue() + " "
	}
	return selectStr
}
