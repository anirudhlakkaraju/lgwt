package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "just a test!"}

	t.Run("key exists", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "just a test!"

		assertStrings(t, got, want)
	})

	t.Run("key does not exist", func(t *testing.T) {
		_, err := dictionary.Search("unkown")
		want := "could not find word you were looking for"

		if err == nil {
			t.Fatalf("expected to get an error")
		}

		assertStrings(t, err.Error(), want)
	})
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Fatalf("got %q want %q given, %q", got, want, "test")
	}
}
