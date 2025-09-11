package main

import (
	"reflect"
	"testing"
)

func Test1(t *testing.T) {
	got := solveNQueens(4)
	want := [][]string{{".Q..", "...Q", "Q...", "..Q."}, {"..Q.", "Q...", "...Q", ".Q.."}}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func Test2(t *testing.T) {
	got := solveNQueens(1)
	want := [][]string{{"Q"}}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
