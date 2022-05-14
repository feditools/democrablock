package webapp

// SessionKey is a key used for storing data in a web session.
type SessionKey int

const (
	// SessionKeyAccountID contains the id of the currently logged-in user.
	SessionKeyAccountID SessionKey = iota
	// SessionKeyLoginRedirect contains the url to be redirected too after logging in.
	SessionKeyLoginRedirect
	// SessionKeyOAuthState contains the state sent to the oauth server.
	SessionKeyOAuthState
	// SessionKeyOAuthCode contains the code sent to the oauth server.
	SessionKeyOAuthCode
	// SessionKeyOAuthToken contains the token returned by the oauth server.
	SessionKeyOAuthToken
)
