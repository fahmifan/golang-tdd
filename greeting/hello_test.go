package greeting

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMsg := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"

		assertCorrectMsg(t, got, want)
	})

	t.Run("empty string defaults to 'World'", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMsg(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", spanish)
		want := "Hola, Elodie"

		assertCorrectMsg(t, got, want)
	})

	t.Run("in Franch", func(t *testing.T) {
		got := Hello("Peter", french)
		want := "Bonjour, Peter"

		assertCorrectMsg(t, got, want)
	})
}
