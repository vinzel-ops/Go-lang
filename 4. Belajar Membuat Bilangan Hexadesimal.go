package main

import "fmt"

func main() {
	// Loop dari 1 hingga 10
	for i := 1; i <= 10; i++ {
		// Mencetak angka dalam format heksadesimal
		// %#X untuk heksadesimal dalam huruf besar dengan prefix 0X
		fmt.Printf("%d \t %#X \n", i, i)
	}
}
