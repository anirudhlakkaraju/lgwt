package main

import "testing"

func TestHelloWorld(t *testing.T) {
	got := Hello("Chris")
	want := "Hello Chris!"

	if got != want {
		t.Errorf("expected=%v. got=%v", want, got)
	}
}
