package main

import "fmt"

func main() {

	data := make(map[string]int)

	data["satu"] = 1
	data["dua"] = 2
	data["tiga"] = 3

	panjangMap := len(data)

	fmt.Printf("Jumlah elemen dalam map: %d\n", panjangMap)

	data["empat"] = 4
	data["lima"] = 5

	panjangMapBaru := len(data)

	fmt.Printf("Jumlah elemen dalam map setelah penambahan: %d\n", panjangMapBaru)
}
