package main

import (
	"testing"
)

func TestEachIndex(t *testing.T) {
	if index := eachRunePos("hoge fuga", "noindex"); len(index) != 0 {
		t.Error("Expect no hits, actual: ", index)
	}

	index := eachRunePos("abracadabra", "a")
	expected := []int{0, 3, 5, 7, 10}

	if len(index) != len(expected) {
		t.Error("Expect 5 hits, actual: ", index)
	}
	for i, pos := range index {
		if pos != expected[i] {
			t.Errorf("Expect %d, actual: %d", expected[i], pos)
		}
	}

	index = eachRunePos("アブラカダブラ", "ブラ")
	expected = []int{1, 5}
	if len(index) != len(expected) {
		t.Error("Expect 2 hits, actual: ", index)
	}
	for i, pos := range index {
		if pos != expected[i] {
			t.Errorf("Expect %d, actual: %d", expected[i], pos)
		}
	}
}
