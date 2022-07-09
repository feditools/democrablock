package models

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/uptrace/bun"
)

type TransactionType string

// Transaction represents a transaction lgo entry.
type Transaction struct {
	ID        int64           `validate:"-" bun:"id,pk,autoincrement"`
	CreatedAt time.Time       `validate:"-" bun:",nullzero,notnull,default:current_timestamp"`
	Type      TransactionType `validate:"-" bun:",nullzero,notnull"`
	MetaData  string          `validate:"-" bun:",nullzero,notnull"`
}

var _ bun.BeforeAppendModelHook = (*Transaction)(nil)

// BeforeAppendModel runs before a bun append operation.
func (t *Transaction) BeforeAppendModel(_ context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		now := time.Now()
		t.CreatedAt = now

		err := validate.Struct(t)
		if err != nil {
			return err
		}
	case *bun.UpdateQuery:
		return errors.New("transaction log entries cannot be updated")
	}

	return nil
}

// GetMetaData marshals metadata into json and adds it to transaction.
func (t *Transaction) GetMetaData(i interface{}) error {
	return json.Unmarshal([]byte(t.MetaData), i)
}

// SetMetaData marshals metadata into json and adds it to transaction.
func (t *Transaction) SetMetaData(i interface{}) error {
	metaData, err := json.Marshal(i)
	if err != nil {
		return err
	}
	t.MetaData = string(metaData)

	return nil
}
