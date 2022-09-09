package main

import (
	"fmt"
)

// CEK NILAI SEPARASI IS PRIMA
func PrimaCek(value int) bool{
	var count int
	var cond bool

	if value == 2{
		return true
	} else if value == 1{
		return false
	} else {
		for i := 1; i <= value; i++{
			if value%i == 0{
				count++
			}
		}
		// CEK FAKTOR BILANGAN > 2
		if count  > 2{
			cond = false
		} else {
			cond = true
		}
	}
	return cond
}

func FullPrima(n int) bool {
	var result bool
	var value1, value2, lenSlice int

	// SEPARASI VALUE
	slice := []int{}
	for n > 0 {
		slice = append(slice, n%10)
		n = n / 10
	} 
	lenSlice =  len(slice)
	if lenSlice < 2{
		value1 = slice[0]
	} else{
		value1 = slice[1]
		value2 = slice[0]
	}

	// CEK APAKAH KEDUANYA BILANGAN PRIMA
	if value1 == 2 && PrimaCek(value2) == false{
		result = false
	} else if (PrimaCek(value1) == true) && (PrimaCek(value2) == true){
		result = true
	} else {
		result = false
	}
	return result 
}

func main() {
	fmt.Println(FullPrima(2))  // true 
	fmt.Println(FullPrima(3))  // true
	fmt.Println(FullPrima(11)) // false
	fmt.Println(FullPrima(13)) // false
	fmt.Println(FullPrima(23)) // true
	fmt.Println(FullPrima(29)) // false
	fmt.Println(FullPrima(37)) // true
	fmt.Println(FullPrima(41)) // false
	fmt.Println(FullPrima(43)) // false
	fmt.Println(FullPrima(53)) // true
}
