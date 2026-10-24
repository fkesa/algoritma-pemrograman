package main

import "fmt"

func cetak_teks() {
	var i, n int
	var teks string
	fmt.Scan(&n, &teks)
	for i = 1; i <= n; i++ {
		fmt.Println(teks)
	}
}

func kue() {
	var i, n, bil, jumlah int
	fmt.Scan(&n)
	for i = 1; i <= n; i++ {
		fmt.Scan(&bil)
		jumlah = bil * 10
		fmt.Println(jumlah)
	}
}

func perkalian() {
	var i, n, bil, hasil int
	fmt.Scan(&n)
	for i = 1; i <= n; i++ {
		fmt.Scan(&bil)
		hasil = bil * i
		fmt.Println(hasil)
	}
}

func olahraga() {
	var i, j, n, langkah, jumlah int
	fmt.Scan(&n)
	for i = 1; i <= n; i++ {
		jumlah = 0
		for j = 1; j <= 7; j++ {
			fmt.Scan(&langkah)
			jumlah = jumlah + langkah
		}
		fmt.Println(jumlah)
	}
}

func sandi() {
	var i, n, bil, b1, b4, jumlah int
	fmt.Scan(&n)
	jumlah = 0
	for i = 1; i <= n; i++ {
		fmt.Scan(&bil)
		b1 = bil / 1000
		b4 = bil % 10
		jumlah = jumlah + b1 + b4
		fmt.Println(jumlah)
	}

}

func bidang_persegi() {
	var i, n, sisi, luas, keliling int
	fmt.Scan(&n)
	for i = 1; i <= n; i++ {
		fmt.Scan(&sisi)
		luas = sisi * sisi
		keliling = 4 * sisi
		fmt.Println(luas, keliling)
	}
}

func koding() {
	var i, n, bil, jumlah, rata_rata int
	fmt.Scan(&n)
	jumlah = 0
	for i = 1; i <= n; i++ {
		fmt.Scan(&bil)
		jumlah = jumlah + bil
	}
	rata_rata = jumlah / n
	fmt.Println(rata_rata)
}

func faktorial() {
	var i, n, faktorial int
	fmt.Scan(&n)
	faktorial = 1
	for i = n; i >= 1; i-- {
		faktorial = faktorial * i
	}
	fmt.Println(faktorial)
}

func faktor_bilangan() {
	var i, n int
	var hasil bool
	fmt.Scan(&n)
	for i = 1; i <= n; i++ {
		hasil = n%i == 0
		fmt.Println(hasil)
	}

}
