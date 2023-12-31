package helloworld

import "testing"

func TestHelloWorld(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		wanted := "Hello, Chris"

		assertMessage(got, wanted, t)
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		wanted := "Hello, World"

		assertMessage(got, wanted, t)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		wanted := "Hola, Elodie"

		assertMessage(got, wanted, t)
	})
}

func assertMessage(got, wanted string, t testing.TB) {
	t.Helper()
	if got != wanted {
		t.Errorf("got %q but have %q", got, wanted)
	}
}
