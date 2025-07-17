package main

import "testing"

func Test1(t *testing.T) {
	got := reverseBits(43261596)
	want := uint32(964176192)

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
