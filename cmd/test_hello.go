package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello()
	want := "hi"

	if got != want {
		t.Errorf("Expected: %q,  received: %q", want, got)
	}
}
