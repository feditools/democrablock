package models

import "time"

type Block struct {
	ID        int64     `validate:"-" bun:"id,pk,autoincrement"`
	CreatedAt time.Time `validate:"-" bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time `validate:"-" bun:",nullzero,notnull,default:current_timestamp"`
}
