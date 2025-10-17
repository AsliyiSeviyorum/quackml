package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Ahmet")
	want := "hi, Ahmet"

	if got != want {
		t.Errorf("Expected: %q,  received: %q", want, got)
	}
}
