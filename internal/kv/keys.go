package kv

import "strconv"

const (
	keyBase = "democrablock:"

	keyFedi         = keyBase + "fedi:"
	keyFediNodeInfo = keyFedi + "ni:"

	keyFileStore             = keyBase + "fs:"
	keyFileStorePresignedURL = keyFileStore + "psu:"

	keySession = keyBase + "session:"

	keyUser            = keyBase + "user:"
	keyUserAccessToken = keyUser + "at:"
)

// KeyFileStorePresignedURL returns the kv key which holds a filestore
// presigned url tokens.
func KeyFileStorePresignedURL(t string) string { return keyFileStorePresignedURL + t }

// KeyFediNodeInfo returns the kv key which holds cached nodeinfo.
func KeyFediNodeInfo(d string) string { return keyFediNodeInfo + d }

// KeySession returns the base kv key prefix.
func KeySession() string { return keySession }

// KeyUserAccessToken returns the kv key which holds a user's access token.
func KeyUserAccessToken(i int64) string { return keyUserAccessToken + strconv.FormatInt(i, 10) }
