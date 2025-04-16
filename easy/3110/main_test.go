package main

import "testing"

func Test1(t *testing.T) {
	got := scoreOfString("hello")
	want := 13

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func Test2(t *testing.T) {
	got := scoreOfString("zaz")
	want := 50

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
