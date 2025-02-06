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

		assertError(t, err, ErrNotFound)
		assertStrings(t, err.Error(), want)
	})
}

func TestAdd(t *testing.T) {
	dictionary := Dictionary{}
	dictionary.Add("test", "this is a test")

	want := "this is a test"
	got, err := dictionary.Search("test")
	if err != nil {
		t.Fatal("should find added work: ", err)
	}

	assertStrings(t, got, want)
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Fatalf("got %q want %q given, %q", got, want, "test")
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Fatalf("got error %q want %q", got, want)
	}
}
