package db

import (
	"context"

	"github.com/feditools/democrablock/internal/models"
)

type TxID string

// DB represents a database client.
type DB interface {
	// Close closes the db connections
	Close(ctx context.Context) Error
	// DoMigration runs database migrations
	DoMigration(ctx context.Context) Error
	// LoadTestData adds test data to the database
	LoadTestData(ctx context.Context) Error
	// ResetCache clears any caches in the module
	ResetCache(ctx context.Context) Error

	// db tx

	TxNew(ctx context.Context) (TxID, Error)
	TxCommit(ctx context.Context, id TxID) Error
	TxRollback(ctx context.Context, id TxID) Error

	// FediAccount

	// CountFediAccounts returns the number of federated social account
	CountFediAccounts(ctx context.Context) (count int64, err Error)
	// CountFediAccountsWithCouncil returns the number of federated social accounts which are on the council
	CountFediAccountsWithCouncil(ctx context.Context) (count int64, err Error)
	// CountFediAccountsForInstance returns the number of federated social account for an instance
	CountFediAccountsForInstance(ctx context.Context, instanceID int64) (count int64, err Error)
	// CreateFediAccount stores the federated social account
	CreateFediAccount(ctx context.Context, account *models.FediAccount) (err Error)
	// IncFediAccountLoginCount updates the login count of a stored federated instance
	IncFediAccountLoginCount(ctx context.Context, account *models.FediAccount) (err Error)
	// ReadFediAccount returns one federated social account
	ReadFediAccount(ctx context.Context, id int64) (account *models.FediAccount, err Error)
	// ReadFediAccountByUsername returns one federated social account
	ReadFediAccountByUsername(ctx context.Context, instanceID int64, username string) (account *models.FediAccount, err Error)
	// ReadFediAccountsPage returns a page of federated social accounts
	ReadFediAccountsPage(ctx context.Context, index, count int) (instances []*models.FediAccount, err Error)
	// UpdateFediAccount updates the stored federated instance
	UpdateFediAccount(ctx context.Context, account *models.FediAccount) (err Error)

	// FediInstance

	// CountFediInstances returns the number of federated instances
	CountFediInstances(ctx context.Context) (count int64, err Error)
	// CreateFediInstance stores the federated instance
	CreateFediInstance(ctx context.Context, instance *models.FediInstance) (err Error)
	// ReadFediInstance returns one federated social instance
	ReadFediInstance(ctx context.Context, id int64) (instance *models.FediInstance, err Error)
	// ReadFediInstanceByDomain returns one federated social instance
	ReadFediInstanceByDomain(ctx context.Context, domain string) (instance *models.FediInstance, err Error)
	// ReadFediInstancesPage returns a page of federated social instances
	ReadFediInstancesPage(ctx context.Context, index, count int) (instances []*models.FediInstance, err Error)
	// UpdateFediInstance updates the stored federated instance
	UpdateFediInstance(ctx context.Context, instance *models.FediInstance) (err Error)

	// Transaction Log Entry

	// CountTransactions returns the number of federated instances
	CountTransactions(ctx context.Context) (count int64, err Error)
	// CreateTransaction stores the federated instance
	CreateTransaction(ctx context.Context, transaction *models.Transaction) (err Error)
	// CreateTransactionTX stores the federated instance using a transaction
	CreateTransactionTX(ctx context.Context, txID TxID, transaction *models.Transaction) (err Error)
	// ReadTransaction returns one federated social instance
	ReadTransaction(ctx context.Context, id int64) (transaction *models.Transaction, err Error)
	// ReadTransactionsPage returns a page of federated social instances
	ReadTransactionsPage(ctx context.Context, index, count int) (transactions []*models.Transaction, err Error)
}
