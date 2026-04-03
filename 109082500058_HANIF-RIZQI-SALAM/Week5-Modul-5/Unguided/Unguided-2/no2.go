package main

import "fmt"

func cariFaktor(i, n int) {
	if i > n {
		return
	}
	if n%i == 0 {
		fmt.Printf("%d ", i)
	}
	cariFaktor(i+1, n)
}

func main() {
	var n int
	fmt.Print("Masukkan nilai N: ")
	fmt.Scan(&n)

	fmt.Print("Keluaran: ")
	cariFaktor(1, n)
	fmt.Println()
}
