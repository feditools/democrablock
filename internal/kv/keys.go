package kv

const (
	keyBase = "democrablock:"

	keySession = keyBase + "session:"
)

// KeySession returns the base kv key prefix.
func KeySession() string { return keySession }
