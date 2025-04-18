package main

import (
	"reflect"
	"testing"
)

func Test1(t *testing.T) {
	got := countAndSay(4)
	want := "1211"

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func Test2(t *testing.T) {
	got := countAndSay(1)
	want := "1"

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestMapPairs(t *testing.T) {
	got := mapPairs("223314444411")
	want := [][]int{{2, 2}, {3, 2}, {1, 1}, {4, 5}, {1, 2}}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestConvToStr(t *testing.T) {
	got := convToStr([][]int{{2, 2}, {3, 2}, {1, 1}, {4, 5}, {1, 2}})
	want := "2223115421"

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
