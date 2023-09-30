package main

import "fmt"

func main() {
	pelanggan := map[string]int{
		"Alice":   30,
		"Bob":     25,
		"Charlie": 35,
	}

	nama := "Bob"
	usia, ada := pelanggan[nama]

	if ada {
		fmt.Printf("%s memiliki usia %d tahun.\n", nama, usia)
	} else {
		fmt.Printf("Pelanggan dengan nama %s tidak ditemukan.\n", nama)
	}

	pelanggan["David"] = 28

	fmt.Println("Data Pelanggan:")
	for nama, usia := range pelanggan {
		fmt.Printf("%s - Usia: %d tahun\n", nama, usia)
	}
}
