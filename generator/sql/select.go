package sql

type Select struct {
	sel *Queue
}

func NewSelect() *Select {
	sel := Select{}
	sel.sel = NewQueue()
	return &sel
}

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
