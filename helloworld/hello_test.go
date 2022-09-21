package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Chrystopher")
	want := "Hello, Chrystopher"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
