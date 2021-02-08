package sql

import "github.com/winily/go-utils/conversion"

type Where struct {
	where *Queue
}

func NewWhere() *Where {
	w := Where{}
	w.where = NewQueue()
	return &w
}

func (where *Where) Eq(column string, val interface{}) *Where {
	columnWhere := ConstituteWherePhrase(column, EQ, conversion.ToString(val))
	where.where.Enqueue(columnWhere)
	return where
}

func (where *Where) IsEq(condition ConditionFunc, column string, val interface{}) *Where {
	if condition(val) {
		where.Eq(column, val)
	}
	return where
}

func (where *Where) AllEq(params map[string]interface{}) *Where {
	for column, val := range params {
		where.Eq(column, val)
	}
	return where
}

func (where *Where) IsAllEq(condition ConditionFunc, params map[string]interface{}) *Where {
	for column, val := range params {
		where.IsEq(condition, column, val)
	}
	return where
}

/**
TODO 缺少大量操作函数，参考 Mybatis plus 生成器篇，后面考虑兼容多表查询操作，as 别名问题
*/

/**
用户手写的 sql 拼接到 where 上
按照调用顺序拼接
前面语句自动添加 AND
*/
func (where *Where) SQL(sql string) *Where {
	where.where.Enqueue(sql)
	return where
}

func (where *Where) ToString() string {
	if where.where.Size() == 0 {
		return ""
	}

	whereSQL := WHERE
	for where.where.Size() > 0 {
		whereSQL += " " + where.where.Dequeue() + " "
		if where.where.Size() >= 1 {
			whereSQL += AND
		}
	}
	return whereSQL
}
