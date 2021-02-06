package sql

/**
所有的关键词都加上前后空格，防止拼接sql 关键字粘连出错
*/
const (
	SELECT   = " SELECT "
	CREATE   = "CREATE "
	TABLE    = " TABLE "
	INSERT   = " INSERT "
	INTO     = " INTO "
	FROM     = " FROM "
	VALUES   = " VALUES "
	UPDATE   = " UPDATE "
	SET      = " SET "
	WHERE    = " WHERE "
	DROP     = " DROP "
	ALTER    = " ALTER "
	ADD      = " ADD "
	OR       = " OR "
	ALL      = " ALL "
	JOIN     = " JOIN "
	AND      = " AND "
	NOT      = " NOT "
	NULL     = " NULL "
	DISTINCT = " DISTINCT "
	AS       = " AS "
	ORDER    = " ORDER "
	DESC     = " DESC "
	ASC      = " ASC "
	IS       = " IS "
	BY       = " BY "
	DELETE   = " DELETE "
	BETWEEN  = " BETWEEN "
	GROUP    = " GROUP "
	AVG      = " AVG "
	SUM      = " SUM "
	MIN      = " MIN "
	MAX      = " MAX "
	COUNT    = " COUNT "
	HAVING   = " HAVING "
	LIKE     = " LIKE "
)

/**
组合词
*/

const (
	CREATE_TABLE = CREATE + TABLE
	INSERT_INTO  = INSERT + INTO
	DELETE_FROM  = DELETE + FROM
	DROP_TABLE   = DROP + TABLE
	ALTER_TABLE  = ALTER + TABLE
	ORDER_BY     = ORDER + BY
	NOT_BETWEEN  = NOT + BETWEEN
	IS_NULL      = IS + NULL
	IS_NOT_NULL  = IS + NOT + NULL
	GROUP_BY     = GROUP + BY
	NOT_LIKE     = NOT + LIKE
)

/**
运算符
*/

const (
	EQ      = " = "
	NE      = " <> "
	GT      = " > "
	GE      = " >= "
	LT      = " < "
	LE      = " <= "
	PERCENT = "%" // 这个符号用来模糊搜索等操作所以前后不加空格，防止查询出错
)

/**
标点符号
*/
const (
	LEFT_BRACKETS  = " ( "
	RIGHT_BRACKETS = " ) "
	BACK_QUOTE     = "`"
	COMMA          = ", "
	PERIOD         = "."
)
