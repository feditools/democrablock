package models

import (
	"context"
	"errors"
	"time"

	"github.com/uptrace/bun"
)

type TransactionType string

const (
	TransactionTypeCouncilInit TransactionType = "COUNCIL_INIT"
)

// Transaction represents a transaction lgo entry.
type Transaction struct {
	ID        int64           `validate:"-" bun:"id,pk,autoincrement"`
	CreatedAt time.Time       `validate:"-" bun:",nullzero,notnull,default:current_timestamp"`
	Type      TransactionType `validate:"-" bun:",nullzero,notnull"`
}

var _ bun.BeforeAppendModelHook = (*Transaction)(nil)

// BeforeAppendModel runs before a bun append operation.
func (f *Transaction) BeforeAppendModel(_ context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		now := time.Now()
		f.CreatedAt = now

		err := validate.Struct(f)
		if err != nil {
			return err
		}
	case *bun.UpdateQuery:
		return errors.New("transaction log entries cannot be updated")
	}

	return nil
}
