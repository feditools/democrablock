package immudb

import "fmt"

const (
	keyFediAccount           = "fediaccount:"
	keyFediAccountLoginCount = keyFediAccount + "logincount:"
	keyFediAccountLoginLast  = keyFediAccount + "loginlast:"
)

func KeyFediAccountLoginCount(accountID int64) []byte {
	return []byte(fmt.Sprintf("%s%d", keyFediAccountLoginCount, accountID))
}

func KeyFediAccountLoginLast(accountID int64) []byte {
	return []byte(fmt.Sprintf("%s%d", keyFediAccountLoginLast, accountID))
}
