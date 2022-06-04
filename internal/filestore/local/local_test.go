package local

import (
	"testing"

	"github.com/feditools/democrablock/internal/filestore"
)

func TestModule_ImplementsFileStore(t *testing.T) {
	t.Parallel()

	var _ filestore.FileStore = (*Module)(nil)
}
