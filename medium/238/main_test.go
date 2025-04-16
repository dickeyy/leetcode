package main

import (
	"reflect"
	"testing"
)

func Test1(t *testing.T) {
	got := productExceptSelf([]int{1, 2, 3, 4})
	want := []int{24, 12, 8, 6}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func Test2(t *testing.T) {
	got := productExceptSelf([]int{-1, 1, 0, -3, 3})
	want := []int{0, 0, 9, 0, 0}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
