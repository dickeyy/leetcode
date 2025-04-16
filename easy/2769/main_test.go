package main

import "testing"

func Test1(t *testing.T) {
	want := theMaximumAchievableX(4, 1)
	got := 6

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
