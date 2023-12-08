package main

import (
	"fmt"
	"strings"
)

// Main function of the program
func main() {
	// Simple string printing
	fmt.Println("Halo Vinzel") // Output: Halo Vinzel

	// Length of a string
	fmt.Println("String length:", len("Halo Vinzel")) // Output: 9

	// Accessing a specific character (byte) in a string
	// Note: Indexing returns the byte value (ASCII), not the character itself.
	fmt.Println("Second character (ASCII):", "Nasi goreng"[1]) // Output: ASCII value of 'e'

	// Printing numbers with strings
	fmt.Println("Satu:", 1) // Output: Satu: 1
	fmt.Println("Tiga koma lima:", 3.5) // Output: Tiga koma lima: 3.5

	// Boolean data type
	fmt.Println("Boolean example:", true) // Output: Boolean example: true

	// Variable declaration
	var name = "Vinzel"
	fmt.Println("Name:", name)

	// Short variable declaration
	country := "Singapore"
	fmt.Println("Country:", country)

	// Declaring integer variable
	var age int64 = 38
	fmt.Println("Age:", age)

	// Declaring multiple variables
	var (
		firstname = "Vinzel"
		lastname  = "Xenoasia"
	)
	fmt.Println("Firstname:", firstname)
	fmt.Println("Lastname:", lastname)

	// Memanggil Fungsi Custom
	printUpperCase("Hello, GoLang!") // Fungsi Custom
}

// Custom function to print a string in uppercase
func printUpperCase(s string) {
	upper := strings.ToUpper(s)
	fmt.Println("Upper case string:", upper)
}
