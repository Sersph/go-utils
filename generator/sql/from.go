package sql

import "reflect"

type From struct {
	from *Queue
}

func NewFrom() *From {
	from := From{}
	from.from = NewQueue()
	return &from
}

func (from *From) TableByEntity(table interface{}) *From {
	tableName := reflect.TypeOf(table).Name()
	from.Table(tableName)
	return from
}

func (from *From) Table(tableName string) *From {
	from.from.Enqueue(tableName)
	return from
}

func (from *From) ToString() string {
	fromSql := FROM
	for from.from.Size() > 0 {
		fromSql += " " + from.from.Dequeue() + " "
		if from.from.Size() >= 1 {
			fromSql += ","
		}
	}
	return fromSql
}

func (from *From) IsNull() bool {
	return from.from.Size() <= 0
}
