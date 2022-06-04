package filestore

import (
	"fmt"
	"testing"

	"github.com/tmthrgd/go-hex"
)

func TestMakeHashDirs(t *testing.T) {
	t.Parallel()

	tables := []struct {
		input  []byte
		output string
	}{
		{
			hex.MustDecodeString("000000000000"),
			"00/00/00/",
		},
		{
			hex.MustDecodeString("000100010001"),
			"00/01/00/",
		},
		{
			hex.MustDecodeString("123456789abc"),
			"12/34/56/",
		},
	}

	for i, table := range tables {
		i := i
		table := table

		name := fmt.Sprintf("[%d] Running MakeHashDirs", i)
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			response := MakeHashDirs(table.input)
			if response != table.output {
				t.Errorf("[%d] got wrong string, got: '%s', want: '%s'", i, response, table.output)
			}
		})
	}
}
