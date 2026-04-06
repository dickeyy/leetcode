package main

import "testing"

func Test1(t *testing.T) {
	got := judgeCircle("UD")
	want := true

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func Test2(t *testing.T) {
	got := judgeCircle("LL")
	want := false

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
