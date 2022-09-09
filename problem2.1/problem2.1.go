package main

import "fmt"

func FaktorBilangan(n int) string {
	var result string
	for i := 1; i <= n; i++{
		if n % i == 0{
			result += fmt.Sprintln(i)
		}
	}
	return result
}

func main() {
	var number int
	fmt.Scanf("%d", &number)
	fmt.Println(FaktorBilangan(number))
}