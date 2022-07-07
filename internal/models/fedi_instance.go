package models

import (
	"context"
	"time"

	"github.com/uptrace/bun"
)

// FediInstance represents a federated social instance.
type FediInstance struct {
	ID             int64     `validate:"-" bun:"id,pk,autoincrement"`
	CreatedAt      time.Time `validate:"-" bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt      time.Time `validate:"-" bun:",nullzero,notnull,default:current_timestamp"`
	Domain         string    `validate:"-" bun:",nullzero,notnull,unique"`
	ActorURI       string    `validate:"url" bun:",nullzero,notnull"`
	ServerHostname string    `validate:"-" bun:",nullzero,notnull,unique"`
	Software       string    `validate:"-" bun:",nullzero,notnull"`
}

var _ bun.BeforeAppendModelHook = (*FediInstance)(nil)

// BeforeAppendModel runs before a bun append operation.
func (f *FediInstance) BeforeAppendModel(_ context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		now := time.Now()
		f.CreatedAt = now
		f.UpdatedAt = now

		err := validate.Struct(f)
		if err != nil {
			return err
		}
	case *bun.UpdateQuery:
		f.UpdatedAt = time.Now()

		err := validate.Struct(f)
		if err != nil {
			return err
		}
	}

	return nil
}
