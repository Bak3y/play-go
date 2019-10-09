package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Matt", "")
		want := "Hello, Matt"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("", "spanish")
		want := "Hola, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish Matteo", func(t *testing.T) {
		got := Hello("Matteo", "spanish")
		want := "Hola, Matteo"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("", "french")
		want := "Bonjour, World"
		assertCorrectMessage(t, got, want)
	})

}
