package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Saying hello with a name specified", func(t *testing.T) {
		got := Hello("Ahmet")
		want := "hi, Ahmet"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Saying hello without a name specified", func(t *testing.T) {
		got := Hello("")
		want := "hi, SUPRA"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) { // testing is for both tests and benchmarks
	t.Helper() // original function that called this will be reported
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
