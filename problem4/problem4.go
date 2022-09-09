package main

import (
	"fmt"
)

func Palindrome(input string) bool {
	// Reverse String dari inputan
	runeStr := []rune(input)
	for i, j := 0, len(runeStr)-1; i<j; i, j = i+1, j-1{
		runeStr[i], runeStr[j] = runeStr[j], runeStr[i]
	}
	result := string(runeStr)
	
	// Cek apakah inputan == reverse inputan
	if result == input{
		return true
	} else if result != input {
		return false
	} else {
		panic("bukan huruf")
	}
}

func main() {
	fmt.Println(Palindrome("civic"))       // true
	fmt.Println(Palindrome("katak"))       // true
	fmt.Println(Palindrome("kasur rusak")) // true
	fmt.Println(Palindrome("kupu-kupu"))   // false
	fmt.Println(Palindrome("lion"))        // false
}
