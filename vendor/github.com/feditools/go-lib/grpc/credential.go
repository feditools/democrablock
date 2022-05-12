package grpc

import (
	"context"
)

// Credential contains feditools grpc token.
type Credential struct {
	token string
}

// NewCredential creates a new feditools grpc token.
func NewCredential(token string) *Credential {
	return &Credential{
		token: token,
	}
}

// GetRequestMetadata add the token to the request.
func (c *Credential) GetRequestMetadata(_ context.Context, _ ...string) (map[string]string, error) {
	return map[string]string{
		"authorization": c.token,
	}, nil
}

// RequireTransportSecurity returns false since transport security is not required.
func (Credential) RequireTransportSecurity() bool {
	return false
}
