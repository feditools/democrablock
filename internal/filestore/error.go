package filestore

import "fmt"

// Error represents a database specific error.
type Error error

var (
	// ErrNotFound is returned when the file value is not found.
	ErrNotFound Error = fmt.Errorf("not found")
)
