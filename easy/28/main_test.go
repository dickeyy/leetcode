package main

import "testing"

func Test1(t *testing.T) {
	got := strStr("sadbutsad", "sad")
	want := 0

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
