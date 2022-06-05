package local

import (
	"fmt"
	"testing"
)

func TestCheckPathSanity(t *testing.T) {
	t.Parallel()

	tables := []struct {
		hash string
		bits []string

		sane bool
	}{
		{
			"e4781f87a0bcaad14bef739d03d59b397abc0b6bfb5adfc1b0af8745c61b2c92",
			[]string{"e4", "78", "1f"},

			true,
		},
		{
			"77ec511b86d1c91739fe90d31f85353b853f1e4dd51fcd0c1dd91d5f61a26558",
			[]string{"77", "ec", "51"},

			true,
		},
		{
			"fd3541528480a292e39fb1b89b1fbea076077b9f82092a4e71273d5482bac88c",
			[]string{"77", "ec", "51"},

			false,
		},
		{
			"fd3541528480a292e39fb1b89b1fbea076077b9f82092a4e71273d5482bac88c",
			[]string{"77", "ec"},

			false,
		},
	}

	for i, table := range tables {
		i := i
		table := table

		name := fmt.Sprintf("[%d] Getting id", i)
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			resp := checkPathSanity(table.hash, table.bits...)
			if resp != table.sane {
				t.Errorf("[%d] got wrong sanity, got: '%t', want: '%t'", i, resp, table.sane)
			}
		})
	}
}
