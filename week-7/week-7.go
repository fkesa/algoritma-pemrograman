package main

import "fmt"

func prisma() {
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

	fmt.Printf("%0.1f %0.1f", pr.volume, pr.luas_permukaan)
}

func waktu() {
	var total_detik, j, m, d int

	fmt.Scan(&total_detik)
	j = total_detik / 3600
	m = (total_detik % 3600) / 60
	d = total_detik % 60

	fmt.Println(j, m, d)
}

func vektor() {
	var i, n int
	var x, y, z, jumlah_x, jumlah_y, jumlah_z, rata_rata_x, rata_rata_y, rata_rata_z float64

	fmt.Scan(&n)
	jumlah_x = 0.0
	jumlah_y = 0.0
	jumlah_z = 0.0
	for i = 1; i <= 3; i++ {
		fmt.Scan(&x, &y, &z)
		jumlah_x = jumlah_x + x
		jumlah_y = jumlah_y + y
		jumlah_z = jumlah_z + z
	}
	rata_rata_x = jumlah_x / 3.0
	rata_rata_y = jumlah_y / 3.0
	rata_rata_z = jumlah_z / 3.0

	fmt.Println(rata_rata_x, rata_rata_y, rata_rata_z)
}

func determinan() {
	var a, b, c, d, det float64

	fmt.Scan(&a, &b, &c, &d)
	det = (a * d) - (b * c)

	fmt.Println(det)
}

func bola() {
	var i, n, massa, volume, jumlah int
	var rata_rata float64

	fmt.Scan(&n)
	jumlah = 0
	for i = 1; i <= n; i++ {
		fmt.Scan(&massa, &volume)
		jumlah = jumlah + massa*volume
	}
	rata_rata = float64(jumlah) / float64(n)
	fmt.Println(rata_rata)
}

func AND() {
	var i, n int
	var in_bool, hasil bool

	fmt.Scan(&n)
	hasil = true
	for i = 1; i <= n; i++ {
		fmt.Scan(&in_bool)
		hasil = hasil == in_bool
	}
	fmt.Println(hasil)
}

func sum_ganjil() {
	var i, x, y, mulai, jumlah int

	fmt.Scan(&x, &y)
	mulai = x + 1 - (x % 2)
	jumlah = 0
	for i = mulai; i <= y; i += 2 {
		jumlah = jumlah + i
	}

	fmt.Println(jumlah)
}