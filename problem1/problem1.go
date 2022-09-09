package main

import "fmt"

func KonversiNilai(nilai int) string {
	if nilai <= 34 {
		return "D"
	} else if nilai <= 49 {
		return "C"
	} else if nilai <= 64 {
		return "B"
	}else if nilai <= 79 {
		return "B+"
	}else if nilai > 79 {
		return "A"
	}else {
		return "Harus input angka"
	}
}

func main() {
	var nilai int = 80
	
	fmt.Println(KonversiNilai(nilai))
}
 