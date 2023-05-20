package main

import (
	"testing"
)

//On peut changer want pour voir le test fail.

func TestHello(t *testing.T) {
	got := Hello()

	want := "Hello, Word!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
