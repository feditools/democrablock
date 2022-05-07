package redis

import (
	"github.com/feditools/democrablock/internal/log"
)

type empty struct{}

var logger = log.WithPackageField(empty{})
