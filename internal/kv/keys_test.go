package kv

import "testing"

func TestKeyFileStorePresignedURL(t *testing.T) {
	want := "democrablock:fs:psu:25729f65-7dbd-40dc-8859-784874f13390"

	if v := KeyFileStorePresignedURL("25729f65-7dbd-40dc-8859-784874f13390"); v != want {
		t.Errorf("enexpected value for KeyFileStorePresignedURL, got: '%s', want: '%s'.", v, want)
	}
}

func TestKeyFediNodeInfo(t *testing.T) {
	want := "democrablock:fedi:ni:example.com"

	if v := KeyFediNodeInfo("example.com"); v != want {
		t.Errorf("enexpected value for KeyFediNodeInfo, got: '%s', want: '%s'.", v, want)
	}
}

func TestKeySession(t *testing.T) {
	want := "democrablock:session:"

	if v := KeySession(); v != want {
		t.Errorf("enexpected value for TestKeySession, got: '%s', want: '%s'.", v, want)
	}
}
