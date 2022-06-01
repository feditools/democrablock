package immudb

import "fmt"

const selectPageHelperStatement = `
SELECT id FROM %s WHERE id > %d ORDER BY %s %s LIMIT %d;`

func selectPageHelper(tableName string, lastReadID int64, count int, orderBy string, ascending bool) string {
	return fmt.Sprintf(
		selectPageHelperStatement,
		tableName,            // Table Name
		lastReadID,           // Where id
		orderBy,              // Order By
		sortOrder(ascending), // Sorting Order
		count,                // Limit
	)
}
