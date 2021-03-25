package main

import "testing"

func TestHello(t *testing.T) {

    assertCorrectMessage := func(t testing.TB, got, want string) {
        // just a helper to help in reporting failure stack trace
        t.Helper()
        if got != want {
            t.Errorf("got %q while we wanted %q", got, want)
        }
    }

    t.Run("say hello to people", func(t *testing.T) {
        got := Hello("Chris", "English")
        want := "Hello, Chris"
        assertCorrectMessage(t, got, want)
    })

    t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
        got := Hello("", "English")
        want := "Hello, World"
        assertCorrectMessage(t, got, want)
    })

    t.Run("in Spanish", func(t *testing.T) {
        got := Hello("Elodie", "Spanish")
        want := "Hola, Elodie"
        assertCorrectMessage(t, got, want)
    })

    t.Run("in French", func(t *testing.T) {
        got := Hello("Elodie", "French")
        want := "Bonjour, Elodie"
        assertCorrectMessage(t, got, want)
    })
}
