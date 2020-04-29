package hello

import "testing"

func TestHello(t *testing.T) {
	want := "Bonjour le monde."
    if got := Hello(); got != want {
        t.Errorf("Hello() = %q, want %q", got, want)
    }
}

func TestProverb(t *testing.T) {
    want := "Concurrency is not parallelism."
    if got := Proverb(); got != want {
        t.Errorf("Proverb() = %q, want %q", got, want)
    }
}

func TestPerson(t *testing.T) {
    want := "Personne !?"
    if got := Person(); got != want {
        t.Errorf("Person() = %q, want %q", got, want)
    }
}