package main

import (
	"reflect"
	"testing"
)

func TestChunk(t *testing.T) {
	//chunk divides an array of 10 elements with chunk size 2
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	chunked := Chunk(arr, 2)
	expected := [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}, {9, 10}}
	if !reflect.DeepEqual(chunked, expected) {
		t.Error("Test 1: Output Mismatch")
	}

	//chunk divides an array of 3 elements with chunk size 1
	arr = []int{1, 2, 3}
	chunked = Chunk(arr, 1)
	expected = [][]int{{1}, {2}, {3}}
	if !reflect.DeepEqual(chunked, expected) {
		t.Error("Test 2: Output Mismatch")
	}

	//chunk divides an array of 5 elements with chunk size 3
	arr = []int{1, 2, 3, 4, 5}
	chunked = Chunk(arr, 3)
	expected = [][]int{{1, 2, 3}, {4, 5}}
	if !reflect.DeepEqual(chunked, expected) {
		t.Error("Test 3: Output Mismatch")
	}
	//chunk divides an array of 13 elements with chunk size 5
	arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}
	chunked = Chunk(arr, 5)
	expected = [][]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}, {11, 12, 13}}
	if !reflect.DeepEqual(chunked, expected) {
		t.Error("Test 4: Output Mismatch")
	}
}
