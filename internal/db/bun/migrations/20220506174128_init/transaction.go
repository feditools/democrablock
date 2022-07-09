package models

import "time"

// Transaction represents a transaction lgo entry.
type Transaction struct {
	ID        int64     `validate:"-" bun:"id,pk,autoincrement"`
	CreatedAt time.Time `validate:"-" bun:",nullzero,notnull,default:current_timestamp"`
	Type      string    `validate:"-" bun:",nullzero,notnull"`
	MetaData  string    `validate:"-" bun:",nullzero,notnull"`
}
