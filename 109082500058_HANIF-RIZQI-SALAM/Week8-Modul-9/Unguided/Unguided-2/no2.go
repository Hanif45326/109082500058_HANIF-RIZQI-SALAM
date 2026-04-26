package main

import (
	"fmt"
	"math"
)

func main() {
	var n, x, idxHapus, cariFrekuensi int
	const kapasitas = 100
	var arr [kapasitas]int

	fmt.Print("Masukkan jumlah elemen (N): ")
	fmt.Scan(&n)

	fmt.Printf("Masukkan %d elemen angka:\n", n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	fmt.Print("\na. Isi seluruh array: ")
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", arr[i])
	}
	fmt.Println()

	fmt.Print("b. Elemen indeks ganjil: ")
	for i := 1; i < n; i += 2 {
		fmt.Printf("%d ", arr[i])
	}
	fmt.Println()

	fmt.Print("c. Elemen indeks genap: ")
	for i := 0; i < n; i += 2 {
		fmt.Printf("%d ", arr[i])
	}
	fmt.Println()
	fmt.Print("d. Masukkan nilai x untuk kelipatan indeks: ")
	fmt.Scan(&x)
	fmt.Printf("Elemen indeks kelipatan %d: ", x)
	for i := 0; i < n; i++ {
		if i%x == 0 {
			fmt.Printf("%d ", arr[i])
		}
	}
	fmt.Println()

	var total float64
	for i := 0; i < n; i++ {
		total += float64(arr[i])
	}
	rataRata := total / float64(n)
	fmt.Printf("f. Rata-rata: %.2f\n", rataRata)

	var jumlahKuadratSelisih float64
	for i := 0; i < n; i++ {
		jumlahKuadratSelisih += math.Pow(float64(arr[i])-rataRata, 2)
	}
	stdDev := math.Sqrt(jumlahKuadratSelisih / float64(n))
	fmt.Printf("g. Standar Deviasi: %.2f\n", stdDev)
	fmt.Print("h. Masukkan bilangan yang ingin dicari frekuensinya: ")
	fmt.Scan(&cariFrekuensi)
	count := 0
	for i := 0; i < n; i++ {
		if arr[i] == cariFrekuensi {
			count++
		}
	}
	fmt.Printf(" Frekuensi bilangan %d: %d kali\n", cariFrekuensi, count)
	fmt.Print("e. Masukkan indeks yang ingin dihapus: ")
	fmt.Scan(&idxHapus)
	for i := idxHapus; i < n-1; i++ {
		arr[i] = arr[i+1]
	}
	n--
	fmt.Print(" Array setelah dihapus: ")
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", arr[i])
	}
	fmt.Println()
}
