package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := map[string]string{"test": "this is just a test"}

	got := Search(dictionary, "test")
	want := "this is just a test"

	assertString(got, want, t)
}

func assertString(got string, want string, t *testing.T) {
	if got != want {
		t.Errorf("got %q want %q given, %q", got, want, "test")
	}
}
