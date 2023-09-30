package main

import "fmt"

func main() {

	data := map[string]int{
		"satu": 1,
		"dua":  2,
		"tiga": 3,
	}

	fmt.Println("Data awal dalam map:")
	for k, v := range data {
		fmt.Printf("Kunci: %s, Nilai: %d\n", k, v)
	}

	kunci := "dua"
	delete(data, kunci)

	fmt.Println("\nData setelah menghapus elemen dengan kunci:", kunci)
	for k, v := range data {
		fmt.Printf("Kunci: %s, Nilai: %d\n", k, v)
	}

	kunciTidakAda := "empat"
	delete(data, kunciTidakAda)

	fmt.Println("\nData setelah mencoba menghapus elemen dengan kunci yang tidak ada:", kunciTidakAda)
	for k, v := range data {
		fmt.Printf("Kunci: %s, Nilai: %d\n", k, v)
	}
}
