package sql

func ColumnProcess(column string) string {
	return BACK_QUOTE + column + BACK_QUOTE
}
