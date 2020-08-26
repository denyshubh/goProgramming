/*--- Directions
Check to see if two provided strings are anagrams of eachother.
One string is an anagram of another if it uses the same characters
in the same quantity. Only consider characters, not spaces
or punctuation.  Consider capital letters to be the same as lower case
--- Examples
  anagrams('rail safety', 'fairy tales') --> True
  anagrams('RAIL! SAFETY!', 'fairy tales') --> True
  anagrams('Hi there', 'Bye there') --> False
*/
package main

import (
	"fmt"
	"strings"
)

func helper(s string) []int {
	s = strings.ToUpper(s)
	charCount := make([]int, 26)
	for _, char := range s {
		if char >= 65 && char < 91 {
			ascii := uint8(char) - 65
			charCount[ascii]++
		}
	}
	return charCount
}

//Anagram function to check if the string is Anagram or not
func Anagram(a, b string) bool {
	charCount1 := helper(a)
	charCount2 := helper(b)

	for i := 0; i < 26; i++ {
		if charCount1[i] != charCount2[i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("-------- Anagram Test ----------")
	a, b := "rail safety", "faiy tales"
	res := Anagram(a, b)
	if res {
		fmt.Println("Anagram Words!!")
	} else {
		fmt.Println("Not Anagram Words!!")
	}

}
