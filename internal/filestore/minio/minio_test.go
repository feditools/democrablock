package minio

import (
	"testing"

	"github.com/feditools/democrablock/internal/filestore"
)

func TestClient_ImplementsDB(t *testing.T) {
	t.Parallel()

	var _ filestore.FileStore = (*Module)(nil)
}
