package main

import "fmt"

func main() {
	type prisma struct {
		panjang        float64
		lebar          float64
		tinggi         float64
		volume         float64
		luas_permukaan float64
	}
	var pr prisma

	fmt.Scan(&pr.panjang, &pr.lebar, &pr.tinggi)
	pr.volume = pr.panjang * pr.lebar * pr.tinggi
	pr.luas_permukaan = 2*(pr.panjang*pr.lebar) + ((2*pr.panjang + 2*pr.lebar) * pr.tinggi)
	fmt.Printf("%0.1f %0.1f" , pr.volume, pr.luas_permukaan)
}
