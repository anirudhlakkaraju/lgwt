package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	numbers := []int{1, 2, 3, 4, 5}

	got := Sum(numbers)
	want := 15

	if got != want {
		t.Errorf("want %d, got %d, given %v", want, got, numbers)
	}
}

func TestSumAll(t *testing.T) {

	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestSumAllTails(t *testing.T) {

	t.Run("Input slices with values", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{0, 9, 11})
		want := []int{5, 20}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("want %d, got %d", want, got)
		}
	})

	t.Run("Input slices with no values", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{0, 9, 11})
		want := []int{0, 20}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("want %d, got %d", want, got)
		}
	})
}
