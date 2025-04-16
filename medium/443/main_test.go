package main

import "testing"

func Test1(t *testing.T) {
	got := compress([]byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'})
	want := 6

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func Test2(t *testing.T) {
	got := compress([]byte{'a'})
	want := 1

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func Test3(t *testing.T) {
	got := compress([]byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'})
	want := 4

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
