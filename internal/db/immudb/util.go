package immudb

import (
	"context"
	"reflect"
	"time"

	"github.com/codenotary/immudb/pkg/api/schema"
)

const (
	microsecondsToNanoseconds = 1000
)

func (c *Client) PageHelper(ctx context.Context, tableName string, index, count int) (int64, error) {
	l := logger.WithField("func", "FediAccountPageHelper")

	lastReadID := int64(0)
	for i := 0; i < index; i++ {
		resp, err := c.db.SQLQuery(
			ctx,
			selectPageHelper(tableName, lastReadID, count, "id", true),
			nil,
			true,
		)
		if err != nil {
			l.Errorf("SQLQuery: %s", err.Error())

			return 0, err
		}

		if len(resp.GetRows()) == 0 {
			return lastReadID, err
		}

		lastReadID = resp.GetRows()[len(resp.GetRows())-1].GetValues()[0].GetN()
	}

	return lastReadID, nil
}

func isNull(v *schema.SQLValue) bool {
	return reflect.TypeOf(v.GetValue()) == reflect.TypeOf(&schema.SQLValue_Null{})
}

//revive:disable:flag-parameter

func sortOrder(ascending bool) string {
	if ascending {
		return "ASC"
	}

	return "DESC"
}

//revive:enable:flag-parameter

func tsToTime(ts int64) time.Time {
	return time.Unix(0, ts*microsecondsToNanoseconds).UTC()
}
