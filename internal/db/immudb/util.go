package immudb

import (
	"context"
	"time"

	"github.com/feditools/democrablock/internal/db/immudb/statements"
)

const (
	microsecondsToNanoseconds = 1000
)

func (c *Client) PageHelper(ctx context.Context, tableName string, index, count int) (int64, error) {
	l := logger.WithField("func", "PageHelper")

	lastReadID := int64(0)
	for i := 0; i < index; i++ {
		// prep params
		params := map[string]interface{}{
			statements.ParamLastReadID: lastReadID,
		}

		// run query
		resp, err := c.db.SQLQuery(
			ctx,
			statements.SelectPageHelper(tableName, true, count),
			params,
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

/* func isNull(v *schema.SQLValue) bool {
	return reflect.TypeOf(v.GetValue()) == reflect.TypeOf(&schema.SQLValue_Null{})
} */

func tsToTime(ts int64) time.Time {
	return time.Unix(0, ts*microsecondsToNanoseconds).UTC()
}
