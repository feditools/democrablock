package path

import (
	"fmt"
	"regexp"
)

var (
	// ReHome matches the Home page.
	ReHome = regexp.MustCompile(fmt.Sprintf(`^?%s$`, Home))
	// ReList matches the List page.
	ReList = regexp.MustCompile(fmt.Sprintf(`^?%s$`, List))
)
