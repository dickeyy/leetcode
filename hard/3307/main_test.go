package main

import (
	"reflect"
	"testing"
)

func Test1(t *testing.T) {
	got := kthCharacter(5, []int{0, 0, 0})
	want := byte('a')

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
