package hello

import "testing"

func TestHello(t *testing.T) {
	t.Run("default language and name", func(t *testing.T) {
		got := Hello("Pranav", "")
		want := "Hello, Pranav"

		if got != want {
			t.Errorf("got: %s want: %s", got, want)
		}
	})

	t.Run("default language and without name", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		if got != want {
			t.Errorf("got: %s want: %s", got, want)
		}
	})

	t.Run("french language and name", func(t *testing.T) {
		got := Hello("Pranav", "French")
		want := "Bonjour, Pranav"

		if got != want {
			t.Errorf("got: %s want: %s", got, want)
		}
	})

	t.Run("spanish language and name", func(t *testing.T) {
		got := Hello("Pranav", "Spanish")
		want := "Hola, Pranav"

		if got != want {
			t.Errorf("got: %s want: %s", got, want)
		}
	})

}
