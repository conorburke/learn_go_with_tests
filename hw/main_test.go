package main

import "testing"

func TestHello(t *testing.T) {
	// got := sayHello("Neo")
	// want := "Hello, Neo"

	// if got != want {
	// 	t.Errorf("got %q want %q", got, want)
	// }
	assertCorrectMessage := func(t testing.TB, got, want string) {
		// when test fails, line in actual test code will be
		// outputted instead of line in this helepr function
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
	t.Run("saying hello to name", func(t *testing.T) {
		got := SayHello("Neo", "")
		want := "Hello, Neo"

		// if got != want {
		// 	t.Errorf("got %q want %q", got, want)
		// }
		assertCorrectMessage(t, got, want)
	})
	t.Run("saying hello to noone", func(t *testing.T) {
		got := SayHello("", "")
		want := "Hello, World"

		// if got != want {
		// 	t.Errorf("got %q want %q", got, want)
		// }
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := SayHello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})
}
