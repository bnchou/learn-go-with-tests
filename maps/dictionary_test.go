package maps

import "testing"

func TestSearch(t *testing.T) {
	Dictionary := Dictionary{"test": "this is just a test"}

	got := Dictionary.Search("test")
	want := "this is just a test"

	assertString(got, want, t)
}

func assertString(got string, want string, t *testing.T) {
	if got != want {
		t.Errorf("got %q want %q given, %q", got, want, "test")
	}
}
