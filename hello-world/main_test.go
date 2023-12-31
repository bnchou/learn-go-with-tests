package main

import "testing"

func TestHelloWorld(t *testing.T) {
	got := Hello()
	wanted := "Hello world"

	if got != wanted {
		t.Errorf("got %q but have %q", got, wanted)
	}
}
