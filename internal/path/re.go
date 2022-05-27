package path

import (
	"fmt"
	"regexp"
)

const (
	// reToken is regex to match a token.
	reToken = `[a-zA-Z0-9_]{16,}` //#nosec G101
)

var (
	// ReHome matches the Home page.
	ReHome = regexp.MustCompile(fmt.Sprintf(`^?%s$`, Home))
	// ReList matches the List page.
	ReList = regexp.MustCompile(fmt.Sprintf(`^?%s$`, List))
)
