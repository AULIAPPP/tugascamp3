package main

import "fmt"

func main() {
	data := map[string]int{
		"satu": 1,
		"dua":  2,
		"tiga": 3,
	}

	kunci := "dua"
	nilai, ada := data[kunci]

	if ada {
		fmt.Printf("Nilai dari kunci '%s' adalah %d\n", kunci, nilai)
	} else {
		fmt.Printf("Kunci '%s' tidak ditemukan dalam map\n", kunci)
	}

	kunciTidakAda := "empat"
	nilaiTidakAda, adaTidakAda := data[kunciTidakAda]

	if adaTidakAda {
		fmt.Printf("Nilai dari kunci '%s' adalah %d\n", kunciTidakAda, nilaiTidakAda)
	} else {
		fmt.Printf("Kunci '%s' tidak ditemukan dalam map\n", kunciTidakAda)
	}
}
