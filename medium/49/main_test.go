package main

import (
	"reflect"
	"testing"
)

func Test1(t *testing.T) {
	got := groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
	want := [][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func Test2(t *testing.T) {
	got := groupAnagrams([]string{"", ""})
	want := [][]string{{"", ""}}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
