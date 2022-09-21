package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("say hello to people", func(t *testing.T) {
		got := Hello("Chrystopher", "")
		want := "Hello, Chrystopher"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say in Spanish", func(t *testing.T) {
		got := Hello("Chrystopher", "spanish")
		want := "Hola, Chrystopher"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hola, Mundo' in spanish when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "spanish")
		want := "Hola, Mundo"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say in French", func(t *testing.T) {
		got := Hello("Chrystopher", "french")
		want := "Bonjour, Chrystopher"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hola, Mundo' in french when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "french")
		want := "Bonjour, Monde"
		assertCorrectMessage(t, got, want)
	})

}

func assertCorrectMessage(t testing.TB, got, want string) {
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
