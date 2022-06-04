package kv

const (
	keyBase = "democrablock:"

	keyFileStore             = keyBase + "fs:"
	keyFileStorePresignedURL = keyFileStore + "psu:"

	keyFedi         = keyBase + "fedi:"
	keyFediNodeInfo = keyFedi + "ni:"

	keySession = keyBase + "session:"
)

// KeyFileStorePresignedURL returns the kv key which holds a filestore
// presigned url tokens.
func KeyFileStorePresignedURL(t string) string { return keyFileStorePresignedURL + t }

// KeyFediNodeInfo returns the kv key which holds cached nodeinfo.
func KeyFediNodeInfo(d string) string { return keyFediNodeInfo + d }

// KeySession returns the base kv key prefix.
func KeySession() string { return keySession }
