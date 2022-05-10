package webapp

const (
	// COAnonymous is an anonymous cross origin.
	COAnonymous = "anonymous"

	// DirStatic is the location of static assets.
	DirStatic = "static"
)

// SessionKey is a key used for storing data in a web session.
type SessionKey int

const (
	// SessionKeyAccountID contains the id of the currently logged-in user.
	SessionKeyAccountID SessionKey = iota
	// SessionKeyLoginRedirect contains the url to be redirected too after logging in.
	SessionKeyLoginRedirect
)
