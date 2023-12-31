package main

import "testing"

func TestHelloWorld(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris")
		wanted := "Hello, Chris"

		if got != wanted {
			t.Errorf("got %q but have %q", got, wanted)
		}
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("")
		wanted := "Hello, World"

		if got != wanted {
			t.Errorf("got %q but have %q", got, wanted)
		}
	})
}
