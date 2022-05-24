package webapp

// SessionKey is a key used for storing data in a web session.
type SessionKey int

const (
	// SessionKeyLoginRedirect contains the url to be redirected too after logging in.
	SessionKeyLoginRedirect SessionKey = iota
	// SessionKeyOAuthToken contains the token returned by the oauth server.
	SessionKeyOAuthToken
	// SessionKeyOAuthJWT contains the jwt token returned by the oauth server.
	SessionKeyOAuthJWT
)
