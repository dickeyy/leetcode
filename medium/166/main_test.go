package main

import "testing"

func Test1(t *testing.T) {
	got := fractionToDecimal(1, 2)
	want := "0.5"

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func Test2(t *testing.T) {
	got := fractionToDecimal(2, 1)
	want := "2"

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func Test3(t *testing.T) {
	got := fractionToDecimal(4, 333)
	want := "0.(012)"

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func Test4(t *testing.T) {
	got := fractionToDecimal(1, 6)
	want := "0.1(6)"

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func Test5(t *testing.T) {
	got := fractionToDecimal(-50, 8)
	want := "-6.25"

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func Test6(t *testing.T) {
	got := fractionToDecimal(0, 3)
	want := "0"

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func Test7(t *testing.T) {
	got := fractionToDecimal(1, -2)
	want := "-0.5"

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func Test8(t *testing.T) {
	got := fractionToDecimal(-1, -2)
	want := "0.5"

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func Test9(t *testing.T) {
	got := fractionToDecimal(-2147483648, 1)
	want := "-2147483648"

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
