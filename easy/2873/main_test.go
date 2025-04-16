package main

import "testing"

func Test1(t *testing.T) {
	got := maximumTripletValue([]int{12, 6, 1, 2, 7})
	want := int64(77)

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func Test2(t *testing.T) {
	got := maximumTripletValue([]int{1, 10, 3, 4, 19})
	want := int64(133)

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func Test3(t *testing.T) {
	got := maximumTripletValue([]int{1, 2, 3})
	want := int64(0)

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
