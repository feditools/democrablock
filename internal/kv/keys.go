package kv

import "strconv"

const (
	keyBase = "democrablock:"

	keyAccount            = keyBase + "acct:"
	keyAccountAccessToken = keyAccount + "at:"

	keyFedi         = keyBase + "fedi:"
	keyFediNodeInfo = keyFedi + "ni:"

	keyFileStore             = keyBase + "fs:"
	keyFileStorePresignedURL = keyFileStore + "psu:"

	keyInstance      = keyBase + "instance:"
	keyInstanceOAuth = keyInstance + "oauth:"

	keySession = keyBase + "session:"
)

// KeyAccountAccessToken returns the kv key which holds a user's access token.
func KeyAccountAccessToken(i int64) string { return keyAccountAccessToken + strconv.FormatInt(i, 10) }

// KeyFileStorePresignedURL returns the kv key which holds a filestore
// presigned url tokens.
func KeyFileStorePresignedURL(t string) string { return keyFileStorePresignedURL + t }

// KeyFediNodeInfo returns the kv key which holds cached nodeinfo.
func KeyFediNodeInfo(d string) string { return keyFediNodeInfo + d }

// KeyInstanceOAuth returns the kv key which holds an instance's oauth tokens.
func KeyInstanceOAuth(i int64) string { return keyInstanceOAuth + strconv.FormatInt(i, 10) }

// KeySession returns the base kv key prefix.
func KeySession() string { return keySession }
