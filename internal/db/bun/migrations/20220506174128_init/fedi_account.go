package models

import "time"

// FediAccount represents a federated social account.
type FediAccount struct {
	ID          int64         `validate:"-" bun:"id,pk,autoincrement"`
	CreatedAt   time.Time     `validate:"-" bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt   time.Time     `validate:"-" bun:",nullzero,notnull,default:current_timestamp"`
	ActorURI    string        `validate:"url" bun:",nullzero,notnull"`
	Username    string        `validate:"-" bun:",unique:unique_fedi_user,nullzero,notnull"`
	InstanceID  int64         `validate:"-" bun:",unique:unique_fedi_user,nullzero,notnull"`
	Instance    *FediInstance `validate:"-" bun:"rel:belongs-to,join:instance_id=id"`
	DisplayName string        `validate:"-" bun:",nullzero"`
	LastFinger  time.Time     `validate:"-" bun:",notnull"`
	LogInCount  int64         `validate:"-" bun:",notnull"`
	LogInLast   time.Time     `validate:"-" bun:",nullzero"`

	// login stuff
	IsAdmin   bool `validate:"-" bun:",notnull"`
	IsCouncil bool `validate:"-" bun:",notnull"`
}
