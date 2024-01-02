package main

import "testing"

func TestFindFirstNumberInString(t *testing.T) {

	var got = FindFirstNumberInString("one2threefour5678")
	var want = 1

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestFindLastNumberInString(t *testing.T) {

	var got = FindLastNumberInString("one234567three")
	var want = 3

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
