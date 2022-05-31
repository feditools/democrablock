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
	// ReAdmin matches the admin page.
	ReAdmin = regexp.MustCompile(fmt.Sprintf(`^?/%s$`, PartAdmin))

	// ReAdminFediversePre matches the admin fediverse page prefix.
	ReAdminFediversePre = regexp.MustCompile(fmt.Sprintf(`^?/%s/%s`, PartAdmin, PartFediverse))
	// ReAdminFediverseAccountsPre matches the admin fediverse page prefix.
	ReAdminFediverseAccountsPre = regexp.MustCompile(fmt.Sprintf(`^?/%s/%s/%s`, PartAdmin, PartFediverse, PartAccounts))
	// ReAdminFediverseInstancesPre matches the admin fediverse page prefix.
	ReAdminFediverseInstancesPre = regexp.MustCompile(fmt.Sprintf(`^?/%s/%s/%s`, PartAdmin, PartFediverse, PartInstances))

	// ReAdminSystemPre matches the admin system page prefix.
	ReAdminSystemPre = regexp.MustCompile(fmt.Sprintf(`^?/%s/%s`, PartAdmin, PartSystem))

	// ReHome matches the Home page.
	ReHome = regexp.MustCompile(fmt.Sprintf(`^?%s$`, Home))
	// ReList matches the List page.
	ReList = regexp.MustCompile(fmt.Sprintf(`^?%s$`, List))
)
