package main

import "fmt"

// bukan latihan soal
func mencari_posisi() {
	var posisi_awal, kecepatan, selang_waktu, posisi_akhir int
	fmt.Scan(&posisi_awal, &kecepatan, &selang_waktu)
	posisi_akhir = posisi_awal + (kecepatan * selang_waktu)
	fmt.Println(posisi_akhir)
}

func temperatur() {
	var suhu_celcius, suhu_reamur, suhu_fahrenheit, suhu_kelvin float64
	fmt.Scan(&suhu_celcius)
	suhu_reamur = (suhu_celcius * 4.0 / 5.0)
	suhu_fahrenheit = (suhu_celcius * 9.0 / 5.0) + 32
	suhu_kelvin = (suhu_celcius + 273.15)
	fmt.Printf("%.2f\n", suhu_reamur)
	fmt.Printf("%.2f\n", suhu_fahrenheit)
	fmt.Printf("%.2f\n", suhu_kelvin)
}

func belajar_character() {
	var alfabet byte
	fmt.Scanf("%c", &alfabet)
	fmt.Println(alfabet)
}

// latihan soal
func persegi_panjang() {
	var p, l int
	var luas, keliling int
	fmt.Scan(&p, &l)
	luas = p * l
	keliling = 2*p + 2*l
	fmt.Println(luas, keliling)
}

func lingkaran() {
	var jari_jari, luas, keliling float64
	fmt.Scan(&jari_jari)
	luas = 22 / 7 * jari_jari * jari_jari
	keliling = 22 / 7 * 2 * jari_jari
	fmt.Println(luas, keliling)
}

func fungsi_xy() {
	var x, y int
	var hasil float64
	fmt.Scan(&x, &y)
	hasil = 1.0/(3.0*float64(x)*3.0*float64(x)+10.0) + 10.0*float64(y) + 7.0
	fmt.Println(hasil)
}

func digit() {
	var angka, d1, d2, d3 int
	fmt.Scan(&angka)
	d1 = angka / 100
	d2 = (angka / 10) % 10
	d3 = angka % 10
	fmt.Println(d1, d2, d3)
}

func toko() {
	var barang1, barang2, barang3 int
	var af_barang1, af_barang2, af_barang3 float64
	fmt.Scan(&barang1, &barang2, &barang3)
	af_barang1 = float64(barang1) + (float64(barang1) * 5.0 / 100.0)
	af_barang2 = float64(barang2) + (float64(barang2) * 5.0 / 100.0)
	af_barang3 = float64(barang3) + (float64(barang3) * 5.0 / 100.0)
	fmt.Println(af_barang1, af_barang2, af_barang3)
}

func main() {
	toko()
}
