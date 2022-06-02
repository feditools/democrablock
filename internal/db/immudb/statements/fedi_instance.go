package statements

import (
	"fmt"
)

const (
	FediInstancesTableName = "fedi_instances"

	FediInstanceColumnNameID             = ColumnNameID
	FediInstanceColumnNameCreatedAt      = "created_at"
	FediInstanceColumnNameUpdatedAt      = "updated_at"
	FediInstanceColumnNameDomain         = "domain"
	FediInstanceColumnNameActorURI       = "actor_uri"
	FediInstanceColumnNameServerHostname = "server_hostname"
	FediInstanceColumnNameSoftware       = "software"
	FediInstanceColumnNameClientID       = "client_id"
	FediInstanceColumnNameClientSecret   = "client_secret"
)

const (
	FediInstanceColumnIndexID int64 = iota
	FediInstanceColumnIndexCreatedAt
	FediInstanceColumnIndexUpdatedAt
	FediInstanceColumnIndexDomain
	FediInstanceColumnIndexActorURI
	FediInstanceColumnIndexServerHostname
	FediInstanceColumnIndexSoftware
	FediInstanceColumnIndexClientID
	FediInstanceColumnIndexClientSecret
)

const fediInstanceAllColumns = FediInstanceColumnNameID + ", " + // 0
	FediInstanceColumnNameCreatedAt + ", " + // 1
	FediInstanceColumnNameUpdatedAt + ", " + // 2
	FediInstanceColumnNameDomain + ", " + // 3
	FediInstanceColumnNameActorURI + ", " + // 4
	FediInstanceColumnNameServerHostname + ", " + // 5
	FediInstanceColumnNameSoftware + ", " + // 6
	FediInstanceColumnNameClientID + ", " + // 7
	FediInstanceColumnNameClientSecret // 8

const countFediInstancesStatement = `
SELECT COUNT(*) FROM %[1]s;`

func CountFediInstances() string {
	return fmt.Sprintf(
		countFediInstancesStatement,
		FediInstancesTableName, // 1-Table Name
	)
}

const insertFediInstanceStatement = `
INSERT INTO %[1]s (
    %[2]s,
    %[3]s,
    %[4]s,
    %[5]s,
    %[6]s,
    %[7]s,
    %[8]s,
    %[9]s
)
VALUES (
    @%[2]s,
    @%[3]s,
    @%[4]s,
    @%[5]s,
    @%[6]s,
    @%[7]s,
    @%[8]s,
    @%[9]s
);`

func InsertFediInstance() string {
	return fmt.Sprintf(
		insertFediInstanceStatement,
		FediInstancesTableName,               // 1-Table Name
		FediInstanceColumnNameCreatedAt,      // 2
		FediInstanceColumnNameUpdatedAt,      // 3
		FediInstanceColumnNameDomain,         // 4
		FediInstanceColumnNameActorURI,       // 5
		FediInstanceColumnNameServerHostname, // 6
		FediInstanceColumnNameSoftware,       // 7
		FediInstanceColumnNameClientID,       // 8
		FediInstanceColumnNameClientSecret,   // 9
	)
}

const selectFediInstanceStatement = `
SELECT %[2]s FROM %[1]s WHERE %[3]s = @%[3]s;`

func SelectFediInstance() string {
	return fmt.Sprintf(
		selectFediInstanceStatement,
		FediInstancesTableName,   // 1-Table Name
		fediInstanceAllColumns,   // 2-Columns
		FediInstanceColumnNameID, // 3
	)
}

const selectFediInstanceByDomainStatement = `
SELECT %[2]s FROM %[1]s WHERE %[3]s = @%[3]s;`

func SelectFediInstanceByDomain() string {
	return fmt.Sprintf(
		selectFediInstanceByDomainStatement,
		FediInstancesTableName,       // 1-Table Name
		fediInstanceAllColumns,       // 2-Columns
		FediInstanceColumnNameDomain, // 3
	)
}

const selectFediInstancesPageStatement = `
SELECT %[2]s FROM %[1]s WHERE %[3]s > @%[4]s ORDER BY %[3]s %[5]s LIMIT %[6]d;`

func SelectFediInstancesPage(asc bool, limit int) string {
	return fmt.Sprintf(
		selectFediInstancesPageStatement,
		FediInstancesTableName,   // 1-Table Name
		fediInstanceAllColumns,   // 2-Columns
		FediInstanceColumnNameID, // 3
		ParamLastReadID,          // 4
		sortOrder(asc),           // 5
		limit,                    // 6
	)
}

const upsertFediInstanceStatement = `
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

func UpsertFediInstance() string {
	return fmt.Sprintf(
		upsertFediInstanceStatement,
		FediInstancesTableName,               // 1-Table Name
		FediInstanceColumnNameID,             // 2
		FediInstanceColumnNameCreatedAt,      // 3
		FediInstanceColumnNameUpdatedAt,      // 4
		FediInstanceColumnNameDomain,         // 5
		FediInstanceColumnNameActorURI,       // 6
		FediInstanceColumnNameServerHostname, // 7
		FediInstanceColumnNameSoftware,       // 8
		FediInstanceColumnNameClientID,       // 9
		FediInstanceColumnNameClientSecret,   // 10
	)
}
