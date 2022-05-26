package models

import (
	"github.com/feditools/go-lib/fedihelper"
	"github.com/sirupsen/logrus"
	"time"
)

// GetActorURI returns the account's actor uri.
func (f *FediAccount) GetActorURI() (actorURI string) {
	return f.ActorURI
}

// GetDisplayName returns the account's display name.
func (f *FediAccount) GetDisplayName() (displayName string) {
	return f.DisplayName
}

// GetInstance returns the instance of the account.
func (f *FediAccount) GetInstance() (instance fedihelper.Instance) {
	return f.Instance
}

// GetLastFinger returns the time of the last finger.
func (f *FediAccount) GetLastFinger() (lastFinger time.Time) {
	return f.LastFinger
}

// GetUsername returns the account's username.
func (f *FediAccount) GetUsername() (username string) {
	return f.ActorURI
}

// SetActorURI sets the account's actor uri.
func (f *FediAccount) SetActorURI(actorURI string) {
	f.ActorURI = actorURI
}

// SetDisplayName sets the account's display name.
func (f *FediAccount) SetDisplayName(displayName string) {
	f.DisplayName = displayName
}

// SetInstance sets the instance of the account.
func (f *FediAccount) SetInstance(instanceI fedihelper.Instance) {
	l := logger.WithFields(logrus.Fields{
		"struct": "FediAccount",
		"func":   "SetInstance",
	})

	instance, ok := instanceI.(*FediInstance)
	if !ok {
		l.Warnf("instance not type *FediInstance")

		return
	}

	f.InstanceID = instance.ID
	f.Instance = instance
}

// SetLastFinger sets the time of the last finger.
func (f *FediAccount) SetLastFinger(lastFinger time.Time) {
	f.LastFinger = lastFinger
}

// SetUsername sets the account's username.
func (f *FediAccount) SetUsername(username string) {
	f.Username = username
}
