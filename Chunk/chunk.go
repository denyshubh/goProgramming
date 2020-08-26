/*--- Directions
Given an array and chunk size, divide the array into many subarrays
where each subarray is of length size
--- Examples
chunk([1, 2, 3, 4], 2) --> [[ 1, 2], [3, 4]]
chunk([1, 2, 3, 4, 5], 2) --> [[ 1, 2], [3, 4], [5]]
chunk([1, 2, 3, 4, 5, 6, 7, 8], 3) --> [[ 1, 2, 3], [4, 5, 6], [7, 8]]
chunk([1, 2, 3, 4, 5], 4) --> [[ 1, 2, 3, 4], [5]]
chunk([1, 2, 3, 4, 5], 10) --> [[ 1, 2, 3, 4, 5]]
*/
package main

import (
	"fmt"
	"math"
)

//Chunk Divides the array into n parts
func Chunk(arr []int, size int) [][]int {
	f := float64(len(arr)) / float64(size)
	n := int(math.Ceil(f))
	a := make([][]int, n)
	i := 0
	for j := 0; j < n-1; j++ {
		a[j] = arr[i : i+size]
		i += size
	}
	a[len(a)-1] = arr[i:len(arr)]
	return a
}
func main() {
	n := 5
	fmt.Printf("Divide array into %d subarray\n", n)
	num := []int{22, 3, 5, 7, 11, 13, 53, 85, 17, 211, 133}
	a := Chunk(num, n)
	fmt.Println(a)
}
