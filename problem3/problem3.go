package main

import (
	"fmt"
)

func PrimeNumber(number int) bool {
	var count int
	for i := 1; i <= number; i++{
		if number%i == 0{
			count++
		}
	}
	
	if count <= 2{
		return true
	} 

	return false
}

func main() {
	fmt.Println(PrimeNumber(3))
	fmt.Println(PrimeNumber(7))
	fmt.Println(PrimeNumber(10))
}
