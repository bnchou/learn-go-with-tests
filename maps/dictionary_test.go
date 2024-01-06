package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertString(got, want, t)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")

		if err == nil {
			t.Fatal("expected to get an error.")
		}

		assertError(err, ErrNotFound, t)
	})
}

func TestAdd(t *testing.T) {
	dictionary := Dictionary{}
	dictionary.Add("test", "this is just a test")

	want := "this is just a test"
	got, err := dictionary.Search("test")
	if err != nil {
		t.Fatal("should find added word:", err)
	}

	assertString(got, want, t)
}

func assertError(got, want error, t *testing.T) {
	if got != want {
		t.Fatalf("expected %q but got %q", want, got)
	}
}

func assertString(got, want string, t *testing.T) {
	if got != want {
		t.Errorf("got %q want %q given, %q", got, want, "test")
	}
}
