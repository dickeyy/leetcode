package main

import "testing"

func Test1(t *testing.T) {
	got := countPrimes(10)
	want := 4

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
