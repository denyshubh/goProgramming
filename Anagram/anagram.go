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

//Anagram function to check if the string is Anagram or not
func Anagram(s string) []int {
	s = strings.ToUpper(s)
	charCount := make([]int, 26)
	for _, char := range s {
		if char != 32 {
			ascii := uint8(char) - 65
			charCount[ascii]++
		}
	}
	return charCount
}

func main() {
	fmt.Println("-------- Anagram Test ----------")
	a, b := "rail safety", "fairy tales"
	charCount1 := Anagram(a)
	charCount2 := Anagram(b)
	flag := 0
	for i := 0; i < 26; i++ {
		if charCount1[i] != charCount2[i] {
			fmt.Println("Not Anagram")
			flag = 1
			break
		}
	}
	if flag == 0 {
		fmt.Println("Anagram Word!!")
	}

}
