package util

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/tmthrgd/go-hex"
)

//revive:disable:add-constant

func TestInt64ToBytes(t *testing.T) {
	t.Parallel()

	tables := []struct {
		input  int64
		output []byte
	}{
		{
			0,
			hex.MustDecodeString("0000000000000000"),
		},
		{
			1,
			hex.MustDecodeString("0000000000000001"),
		},
		{
			286854,
			hex.MustDecodeString("0000000000046086"),
		},
		{
			98432185886,
			hex.MustDecodeString("00000016eb03f61e"),
		},
		{
			1155483248994321,
			hex.MustDecodeString("00041ae7e899a811"),
		},
		{
			9223372036854775807,
			hex.MustDecodeString("7fffffffffffffff"),
		},
	}

	for i, table := range tables {
		i := i
		table := table

		name := fmt.Sprintf("[%d] Running Int64ToBytes", i)
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			result := Int64ToBytes(table.input)

			if len(result) != len(table.output) {
				t.Errorf("[%d] invalid length, got: '%d', want: '%d'", i, len(result), len(table.output))

				return
			}

			if !bytes.Equal(result, table.output) {
				t.Errorf("[%d] invalid bytes, got: '%x', want: '%x'", i, result, table.output)
			}
		})
	}
}

func TestBytesToInt64(t *testing.T) {
	t.Parallel()

	tables := []struct {
		input  []byte
		output int64
	}{
		{
			hex.MustDecodeString("0000000000000000"),
			0,
		},
		{
			hex.MustDecodeString("0000000000000001"),
			1,
		},
		{
			hex.MustDecodeString("0000000000046086"),
			286854,
		},
		{
			hex.MustDecodeString("00000016eb03f61e"),
			98432185886,
		},
		{
			hex.MustDecodeString("00041ae7e899a811"),
			1155483248994321,
		},
		{
			hex.MustDecodeString("7fffffffffffffff"),
			9223372036854775807,
		},
	}

	for i, table := range tables {
		i := i
		table := table

		name := fmt.Sprintf("[%d] Running Int64ToBytes", i)
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			result := BytesToInt64(table.input)

			if result != table.output {
				t.Errorf("[%d] invalid int64, got: '%x', want: '%x'", i, result, table.output)
			}
		})
	}
}

//revive:enable:add-constant
