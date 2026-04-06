package main

import "testing"

func Test1(t *testing.T) {
	got := decodeCiphertext("ch   ie   pr", 3)
	want := "cipher"

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
