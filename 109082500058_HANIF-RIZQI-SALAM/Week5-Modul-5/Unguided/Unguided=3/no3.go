package main

import "fmt"

func cetakBilanganGanjil(i, n int) {
	if i > n {
		return
	}
	if i%2 != 0 {
		fmt.Printf("%d ", i)
	}
	cetakBilanganGanjil(i+1, n)
}

func main() {
	var n int
	fmt.Print("Masukkan nilai N: ")
	fmt.Scan(&n)

	fmt.Print("Keluaran: ")
	cetakBilanganGanjil(1, n)
	fmt.Println()
}
