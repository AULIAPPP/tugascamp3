package main

import "fmt"

func main() {

	data := make(map[string]int)

	data["satu"] = 1
	data["dua"] = 2
	data["tiga"] = 3

	fmt.Println("Data dalam map:")
	for k, v := range data {
		fmt.Printf("Kunci: %s, Nilai: %d\n", k, v)
	}

	nama := make(map[int]string)

	nama[1] = "Alice"
	nama[2] = "Bob"
	nama[3] = "Charlie"

	fmt.Println("\nData dalam map nama:")
	for k, v := range nama {
		fmt.Printf("Kunci: %d, Nilai: %s\n", k, v)
	}
}
