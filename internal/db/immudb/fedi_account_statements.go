package immudb

import (
	"fmt"
	"time"

	"github.com/feditools/democrablock/internal/models"
)

const fediAccountAllColumns = "id, " + // 0
	"created_at, " + // 1
	"updated_at, " + // 2
	"username, " + // 3
	"instance_id, " + // 4
	"actor_uri, " + // 5
	"display_name, " + // 6
	"last_finger, " + // 7
	"access_token, " + // 8
	"is_admin" // 9

const (
	fediAccountColumnID int64 = iota
	fediAccountColumnCreatedAt
	fediAccountColumnUpdatedAt
	fediAccountColumnUsername
	fediAccountColumnInstanceID
	fediAccountColumnActorURI
	fediAccountColumnDisplayName
	fediAccountColumnLastFinger
	fediAccountColumnAccessToken
	fediAccountColumnIsAdmin
)

const countFediAccountsStatement = `
SELECT COUNT(*) FROM %s;`

func countFediAccounts() string {
	return fmt.Sprintf(
		countFediAccountsStatement,
		TableNameFediAccounts, // Table Name
	)
}

const countFediAccountsForInstanceStatement = `
SELECT COUNT(*) FROM %s WHERE instance_id = %d;`

func countFediAccountsForInstance(instanceID int64) string {
	return fmt.Sprintf(
		countFediAccountsForInstanceStatement,
		TableNameFediAccounts, // Table Name
		instanceID,
	)
}

const insertFediAccountStatement = `
INSERT INTO %s (
    created_at,
    updated_at,
    username,
    instance_id,
    actor_uri,
    display_name,
    last_finger,
    is_admin
)
VALUES (
    CAST(%d AS TIMESTAMP),
    CAST(%d AS TIMESTAMP),
    '%s',
    %d,
    '%s',
    '%s',
    CAST(%d AS TIMESTAMP),
    %t
);`

func insertFediAccount(account *models.FediAccount, createdAt time.Time) string {
	return fmt.Sprintf(
		insertFediAccountStatement,
		TableNameFediAccounts,           // Table Name
		createdAt.Unix(),                // created_at
		createdAt.Unix(),                // updated_at
		account.Username,                // username
		account.InstanceID,              // instance_id
		account.ActorURI,                // actor_uri
		account.DisplayName,             // display_name
		account.LastFinger.UTC().Unix(), // last_finger
		account.Admin,                   // is_admin
	)
}

const selectFediAccountStatement = `
SELECT %s FROM %s WHERE id = %d;`

func selectFediAccount(accountID int64) string {
	return fmt.Sprintf(
		selectFediAccountStatement,
		fediAccountAllColumns, // Columns
		TableNameFediAccounts, // Table Name
		accountID,             // id
	)
}

const selectFediAccountByUsernameStatement = `
SELECT %s FROM %s WHERE username = '%s' AND instance_id = %d;`

func selectFediAccountByUsername(instanceID int64, username string) string {
	return fmt.Sprintf(
		selectFediAccountByUsernameStatement,
		fediAccountAllColumns, // Columns
		TableNameFediAccounts, // Table Name
		username,              // username
		instanceID,            // instance_id
	)
}

const selectFediAccountsPageStatement = `
SELECT %s FROM %s WHERE id > %d ORDER BY %s %s LIMIT %d;`

func selectFediAccountsPage(lastReadID int64, count int, orderBy string, ascending bool) string {
	return fmt.Sprintf(
		selectFediAccountsPageStatement,
		fediAccountAllColumns, // Columns
		TableNameFediAccounts, // Table Name
		lastReadID,            // Where id
		orderBy,               // Order By
		sortOrder(ascending),  // Sorting Order
		count,                 // Limit
	)
}

const upsertFediAccountStatement = `
UPSERT INTO %s (
    id,
    updated_at,
    username,
    instance_id,
    actor_uri,
    display_name,
    last_finger,
    is_admin
)
VALUES (
    %d,
    CAST(%d AS TIMESTAMP),
    '%s',
    %d,
    '%s',
    '%s',
    CAST(%d AS TIMESTAMP),
    %t
);`

func upsertFediAccount(account *models.FediAccount, updatedAt time.Time) string {
	return fmt.Sprintf(
		upsertFediAccountStatement,
		TableNameFediAccounts,           // Table Name
		account.ID,                      // id
		updatedAt.Unix(),                // updated_at
		account.Username,                // username
		account.InstanceID,              // instance_id
		account.ActorURI,                // actor_uri
		account.DisplayName,             // display_name
		account.LastFinger.UTC().Unix(), // last_finger
		account.Admin,                   // is_admin
	)
}
