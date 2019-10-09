package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Matt")
	want := "Hello, Matt"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
