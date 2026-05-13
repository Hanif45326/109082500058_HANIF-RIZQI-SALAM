package main

import (
	"fmt"
	"sort"
)

func main() {
	var input int
	suaraMasuk := 0
	suaraSah := 0

	rekapSuara := make(map[int]int)

	for {
		_, err := fmt.Scan(&input)
		if err != nil || input == 0 {
			break
		}

		suaraMasuk++

		if input >= 1 && input <= 20 {
			suaraSah++
			rekapSuara[input]++
		}
	}

	fmt.Printf("Suara masuk: %d\n", suaraMasuk)
	fmt.Printf("Suara sah: %d\n", suaraSah)

	var candidates []int
	for candidate := range rekapSuara {
		candidates = append(candidates, candidate)
	}
	sort.Ints(candidates)

	for _, c := range candidates {
		fmt.Printf("%d: %d\n", c, rekapSuara[c])
	}
}
