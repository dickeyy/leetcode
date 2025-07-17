package main

import "testing"

func Test1(t *testing.T) {
	got := lemonadeChange([]int{5, 5, 5, 10, 20})
	want := true

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func Test2(t *testing.T) {
	got := lemonadeChange([]int{5, 5, 10, 10, 20})
	want := false

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
