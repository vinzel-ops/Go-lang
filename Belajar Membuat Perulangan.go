package main

import "fmt"

func main() {
	// Loop dari 1000000 hingga 1000099
	for i := 1000000; i < 1000100; i++ {
		// Mencetak angka dalam format desimal, biner, dan heksadesimal
		// %d untuk desimal, %b untuk biner, %x untuk heksadesimal
		fmt.Printf("%d \t %b \t %x \n", i, i, i)
	}
}
