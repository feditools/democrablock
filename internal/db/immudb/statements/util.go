package statements

import (
	"fmt"
)

const selectPageHelperStatement = `
SELECT %[2]s FROM %[1]s WHERE %[2]s > @%[3]s ORDER BY %[2]s %[4]s LIMIT %[5]d;`

func SelectPageHelper(tableName string, asc bool, limit int) string {
	return fmt.Sprintf(
		selectPageHelperStatement,
		tableName,       // 1-Table Name
		ColumnNameID,    // 2
		ParamLastReadID, // 3
		sortOrder(asc),  // 4
		limit,           // 5
	)
}

//revive:disable:flag-parameter

func sortOrder(asc bool) string {
	if asc {
		return "ASC"
	}

	return "DESC"
}

//revive:enable:flag-parameter
