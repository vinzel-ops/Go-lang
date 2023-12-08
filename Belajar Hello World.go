package main

import (
	"fmt"
)

func main() {
	// Menampilkan pesan selamat datang
	printWelcomeMessage("Hello world!")
}

// printWelcomeMessage menampilkan pesan selamat datang ke konsol.
// Fungsi ini menerima sebuah parameter 'message' yang berjenis string, 
// yang merupakan pesan yang akan ditampilkan.
func printWelcomeMessage(message string) {
	// Menggunakan Printf untuk output yang terformat
	fmt.Printf("Welcome Message: %s\n", message)
}
