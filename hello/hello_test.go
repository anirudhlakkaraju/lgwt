package hello

import "testing"

func TestHelloWorld(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "English")
		want := "Hello, Chris"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello World!' when empty string is provided", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Martinez", "Spanish")
		want := "Hola, Martinez"

		assertCorrectMessage(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got := Hello("Madame", "French")
		want := "Bonjour, Madame"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	// means current method is a helper
	// so test suite will track og method instead of this
	// in cases of test failure
	t.Helper()
	if got != want {
		t.Errorf("expected=%v. got=%v", want, got)
	}
}
