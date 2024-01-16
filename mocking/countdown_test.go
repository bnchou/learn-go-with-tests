package main

import (
	"bytes"
	"testing"
)

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	spySleep := MockSleep{}

	Countdown(buffer, &spySleep)

	got := buffer.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

	if spySleep.call != 3 {
		t.Errorf("expected 3 calls but got %v", spySleep.call)
	}
}
