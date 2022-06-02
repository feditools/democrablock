package immudb

import (
	"testing"

	"github.com/feditools/democrablock/internal/db"
)

func TestClient_ImplementsDB(t *testing.T) {
	t.Parallel()

	var _ db.DB = (*Client)(nil)
}
