package immudb

import (
	"fmt"
	"time"

	"github.com/feditools/democrablock/internal/models"
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
    last_finger
)
VALUES (
    CAST(%d AS TIMESTAMP),
    CAST(%d AS TIMESTAMP),
    '%s',
    %d,
    '%s',
    '%s',
    CAST(%d AS TIMESTAMP)
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
	)
}
