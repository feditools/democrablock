package kv

import "testing"

func TestKeySession(t *testing.T) {
	want := "democrablock:session:"

	if v := KeySession(); v != want {
		t.Errorf("enexpected value for TestKeySession, got: '%s', want: '%s'.", v, want)
	}
}
