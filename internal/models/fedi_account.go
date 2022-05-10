package models

// FediAccount is a stub of a federated social account.
type FediAccount struct {
	ID          int64
	ActorURI    string
	Username    string
	InstanceID  int64
	Instance    *FediInstance
	DisplayName string
	Admin       bool
}
