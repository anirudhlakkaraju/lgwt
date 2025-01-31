package integers

import "testing"

func TestAddr(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Fatalf("expected '%d', got '%d'", expected, sum)
	}
}
