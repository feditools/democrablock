package statements

import (
	"fmt"
)

const (
	FediAccountsTableName = "fedi_accounts"

	FediAccountColumnNameID          = ColumnNameID
	FediAccountColumnNameCreatedAt   = "created_at"
	FediAccountColumnNameUpdatedAt   = "updated_at"
	FediAccountColumnNameUsername    = "username"
	FediAccountColumnNameInstanceID  = "instance_id"
	FediAccountColumnNameActorURI    = "actor_uri"
	FediAccountColumnNameDisplayName = "display_name"
	FediAccountColumnNameLastFinger  = "last_finger"
	FediAccountColumnNameAccessToken = "access_token"
	FediAccountColumnNameIsAdmin     = "is_admin"
)

const (
	FediAccountColumnIndexID int64 = iota
	FediAccountColumnIndexCreatedAt
	FediAccountColumnIndexUpdatedAt
	FediAccountColumnIndexUsername
	FediAccountColumnIndexInstanceID
	FediAccountColumnIndexActorURI
	FediAccountColumnIndexDisplayName
	FediAccountColumnIndexLastFinger
	FediAccountColumnIndexAccessToken
	FediAccountColumnIndexIsAdmin
)

const fediAccountAllColumns = FediAccountColumnNameID + ", " + // 0
	FediAccountColumnNameCreatedAt + ", " + // 1
	FediAccountColumnNameUpdatedAt + ", " + // 2
	FediAccountColumnNameUsername + ", " + // 3
	FediAccountColumnNameInstanceID + ", " + // 4
	FediAccountColumnNameActorURI + ", " + // 5
	FediAccountColumnNameDisplayName + ", " + // 6
	FediAccountColumnNameLastFinger + ", " + // 7
	FediAccountColumnNameAccessToken + ", " + // 8
	FediAccountColumnNameIsAdmin // 9

const countFediAccountsStatement = `
SELECT COUNT(*) FROM %[1]s;`

func CountFediAccounts() string {
	return fmt.Sprintf(
		countFediAccountsStatement,
		FediAccountsTableName, // Table Name
	)
}

const countFediAccountsForInstanceStatement = `
SELECT COUNT(*) FROM %[1]s WHERE %[2]s = @%[2]s;`

func CountFediAccountsForInstance() string {
	return fmt.Sprintf(
		countFediAccountsForInstanceStatement,
		FediAccountsTableName, // Table Name
		FediAccountColumnNameInstanceID,
	)
}

const insertFediAccountStatement = `
INSERT INTO %[1]s (
    %[2]s,
    %[3]s,
    %[4]s,
    %[5]s,
    %[6]s,
    %[7]s,
    %[8]s,
    %[9]s,
    %[10]s
)
VALUES (
    @%[2]s,
    @%[3]s,
    @%[4]s,
    @%[5]s,
    @%[6]s,
    @%[7]s,
    @%[8]s,
    @%[9]s,
    @%[10]s
);`

func InsertFediAccount() string {
	return fmt.Sprintf(
		insertFediAccountStatement,
		FediAccountsTableName,            // 1-Table Name
		FediAccountColumnNameCreatedAt,   // 2
		FediAccountColumnNameUpdatedAt,   // 3
		FediAccountColumnNameUsername,    // 4
		FediAccountColumnNameInstanceID,  // 5
		FediAccountColumnNameActorURI,    // 6
		FediAccountColumnNameDisplayName, // 7
		FediAccountColumnNameLastFinger,  // 8
		FediAccountColumnNameAccessToken, // 9
		FediAccountColumnNameIsAdmin,     // 10
	)
}

const selectFediAccountStatement = `
SELECT %[2]s FROM %[1]s WHERE %[3]s = @%[3]s;`

func SelectFediAccount() string {
	return fmt.Sprintf(
		selectFediAccountStatement,
		FediAccountsTableName,   // 1-Table Name
		fediAccountAllColumns,   // 2-Columns
		FediAccountColumnNameID, // 3
	)
}

const selectFediAccountByUsernameStatement = `
SELECT %[2]s FROM %[1]s WHERE %[3]s = @%[3]s AND %[4]s = @%[4]s;`

func SelectFediAccountByUsername() string {
	return fmt.Sprintf(
		selectFediAccountByUsernameStatement,
		FediAccountsTableName,           // 1-Table Name
		fediAccountAllColumns,           // 2-Columns
		FediAccountColumnNameUsername,   // 3
		FediAccountColumnNameInstanceID, // 4
	)
}

const selectFediAccountsPageStatement = `
SELECT %[2]s FROM %[1]s WHERE %[3]s > @%[4]s ORDER BY %[3]s %[5]s LIMIT %[6]d;`

func SelectFediAccountsPage(asc bool, limit int) string {
	return fmt.Sprintf(
		selectFediAccountsPageStatement,
		FediAccountsTableName,   // 1-Table Name
		fediAccountAllColumns,   // 2-Columns
		FediAccountColumnNameID, // 3
		ParamLastReadID,         // 4
		sortOrder(asc),          // 5
		limit,                   // 6
	)
}

const upsertFediAccountStatement = `
UPSERT INTO %[1]s (
    %[2]s,
    %[3]s,
    %[4]s,
    %[5]s,
    %[6]s,
    %[7]s,
    %[8]s,
    %[9]s,
    %[10]s
)
VALUES (
    @%[2]s,
    @%[3]s,
    @%[4]s,
    @%[5]s,
    @%[6]s,
    @%[7]s,
    @%[8]s,
    @%[9]s,
    @%[10]s
);`

func UpsertFediAccount() string {
	return fmt.Sprintf(
		upsertFediAccountStatement,
		FediAccountsTableName,            // 1- Table Name
		FediAccountColumnNameID,          // 2
		FediAccountColumnNameUpdatedAt,   // 3
		FediAccountColumnNameUsername,    // 4
		FediAccountColumnNameInstanceID,  // 5
		FediAccountColumnNameActorURI,    // 6
		FediAccountColumnNameDisplayName, // 7
		FediAccountColumnNameLastFinger,  // 8
		FediAccountColumnNameAccessToken, // 9
		FediAccountColumnNameIsAdmin,     // 10
	)
}
