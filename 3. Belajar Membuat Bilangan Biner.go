package main

import "fmt"

func main() {
	// Loop untuk menampilkan representasi biner dari angka 1 hingga 10
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d - %b \n", i, i) // Menampilkan angka dalam format desimal dan biner
	}
}
