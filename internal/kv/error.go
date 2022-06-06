package kv

import "fmt"

// Error represents a database specific error.
type Error error

var (
	// ErrNil is returned when the kv value is nil.
	ErrNil Error = fmt.Errorf("nil")
)
