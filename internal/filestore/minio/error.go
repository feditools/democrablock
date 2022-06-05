package minio

import (
	"github.com/feditools/democrablock/internal/filestore"
)

// ProcessError replaces any known values with our own db.Error types.
func (m *Module) ProcessError(err error) filestore.Error {
	switch {
	case err == nil:
		return nil
	default:
		return err
	}
}
