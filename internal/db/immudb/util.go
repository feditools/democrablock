package immudb

import (
	"reflect"
	"time"

	"github.com/codenotary/immudb/pkg/api/schema"
)

const (
	microsecondsToNanoseconds = 1000
)

func isNull(v *schema.SQLValue) bool {
	return reflect.TypeOf(v.GetValue()) == reflect.TypeOf(&schema.SQLValue_Null{})
}

func tsToTime(ts int64) time.Time {
	return time.Unix(0, ts*microsecondsToNanoseconds).UTC()
}
