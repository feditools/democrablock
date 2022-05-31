package immudb

import (
	"fmt"
	"testing"
	"time"
)

//revive:disable:add-constant

func TestTsToTime(t *testing.T) {
	t.Parallel()

	tables := []struct {
		ts   int64
		time time.Time
	}{
		{1653885379429868, time.Date(2022, 05, 30, 4, 36, 19, 429868000, time.UTC)},
	}

	for i, table := range tables {
		i := i
		table := table

		name := fmt.Sprintf("[%d] Running tsToTime for %d", i, table.ts)
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			newTime := tsToTime(table.ts)
			if !newTime.Equal(table.time) {
				t.Errorf("[%d] invalid time, got: '%s', want: '%s'", i, newTime, table.time)
			}
		})
	}
}

//revive:enable:add-constant
