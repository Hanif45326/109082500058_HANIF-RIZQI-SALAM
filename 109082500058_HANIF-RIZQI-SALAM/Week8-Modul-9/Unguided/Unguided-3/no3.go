package main

import "fmt"

func main() {
	var klubA, klubB string
	var skorA, SkorB int
	var pemenang []string

	fmt.Print("klub A: ")
	fmt.Scanln(&klubA)
	fmt.Print("klub B: ")
	fmt.Scanln(&klubB)

	i := 1
	for {
		fmt.Printf("Pertandingan %d : ", i)
		fmt.Scan(&skorA, &SkorB)
		if skorA < 0 || SkorB < 0 {
			break
		}

		if skorA > SkorB {
			pemenang = append(pemenang, klubA)
		} else if SkorB > skorA {
			pemenang = append(pemenang, klubB)
		} else {
			pemenang = append(pemenang, "Draw")
		}
		i++
	}
	fmt.Println()
	for j, hasil := range pemenang {
		fmt.Printf("Hasil %d : %s\n", j+1, hasil)
	}
	fmt.Println("Pertandingan selesai")
}
