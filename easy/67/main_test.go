package main

import "testing"

func Test1(t *testing.T) {
	got := addBinary("11", "1")
	want := "100"

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func Test2(t *testing.T) {
	got := addBinary("1010", "1011")
	want := "10101"

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func Test3(t *testing.T) {
	got := addBinary(
		"10100000100100110110010000010101111011011001101110111111111101000000101111001110001111100001101",
		"110101001011101110001111100110001010100001101011101010000011011011001011101111001100000011011110011",
	)
	want := "110111101100010011000101110110100000011101000101011001000011011000001100011110011010010011000000000"

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
