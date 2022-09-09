package main

import "fmt"

func FaktorBilanganDesc(n int) string {
	var result string
	for i := n ; i >= 1; i--{
		if n % i == 0{
			result += fmt.Sprintln(i)
		}
	}
	return result
}

func main() {
	var number int
	fmt.Scanf("%d", &number)
	fmt.Println(FaktorBilanganDesc(number))
} 
