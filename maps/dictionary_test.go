package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := map[string]string{"test": "just a test!"}

	got := Search(dictionary, "test")
	want := "just a test!"

	if got != want {
		t.Fatalf("got %q want %q given, %q", got, want, "test")
	}
}
