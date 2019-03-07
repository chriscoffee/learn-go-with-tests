package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%s' but wanted '%s'", got, want)
		}
	}

	t.Run("Say Hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"

		assertCorrectMessage(t, got, want)

	})

	t.Run("Empty string defaults to 'World'", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)

	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
		
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Alex", "french")
		want := "Bonjour, Alex"
		assertCorrectMessage(t, got, want)

	})

	t.Run("in Chinese", func(t *testing.T) {
		got := Hello("Olivia", "chinese")
		want := "你好, Olivia"
		assertCorrectMessage(t, got, want)

	})
}
