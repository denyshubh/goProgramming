package main

import (
	"testing"
)

func TestAnagram(t *testing.T) {

	res := Anagram("hello", "llohe")
	if !res {
		t.Error("Expected words to be Anagram")
	}

	res = Anagram("Whoa! Hi!", "Hi! Whoa!")
	if !res {
		t.Error("Expected words to be Anagram")
	}

	res = Anagram("A tree, a life, a bench", "A tree, a fence, a yard")
	if res {
		t.Error("Expected words to be Anagram")
	}
}
