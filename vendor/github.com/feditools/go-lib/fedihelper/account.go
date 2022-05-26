package fedihelper

import "time"

type Account interface {
	GetActorURI() (actorURI string)
	GetAccessToken() (accessToken string, err error)
	GetDisplayName() (displayName string)
	GetInstance() (instance Instance)
	GetLastFinger() (lastFinger time.Time)
	GetUsername() (username string)

	SetActorURI(actorURI string)
	SetAccessToken(accessToken string) (err error)
	SetDisplayName(displayName string)
	SetInstance(instance Instance)
	SetLastFinger(lastFinger time.Time)
	SetUsername(username string)
}
