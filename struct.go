package main

import (
	"fmt"
	"math"
)

type Lingkaran struct {
	JariJari float64
}

func (l Lingkaran) Luas() float64 {
	return math.Pi * l.JariJari * l.JariJari
}

func (l Lingkaran) Keliling() float64 {
	return 2 * math.Pi * l.JariJari
}

type PersegiPanjang struct {
	Panjang float64
	Lebar   float64
}

func (p PersegiPanjang) Luas() float64 {
	return p.Panjang * p.Lebar
}

func (p PersegiPanjang) Keliling() float64 {
	return 2 * (p.Panjang + p.Lebar)
}

func main() {
	lingkaran := Lingkaran{JariJari: 5.0}
	persegi := PersegiPanjang{Panjang: 4.0, Lebar: 3.0}

	fmt.Println("Luas Lingkaran:", lingkaran.Luas())
	fmt.Println("Keliling Lingkaran:", lingkaran.Keliling())

	fmt.Println("Luas Persegi Panjang:", persegi.Luas())
	fmt.Println("Keliling Persegi Panjang:", persegi.Keliling())
}
