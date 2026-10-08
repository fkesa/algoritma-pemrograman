package main

import "fmt"

func soal_1_lemari_baju() {
	var lebar, panjang, tinggi, volume int
	fmt.Scan(&lebar, &panjang, &tinggi)
	volume = panjang * lebar * tinggi
	fmt.Println(volume)
}

func soal_2_membuat_kopi() {
	var cangkir_kopi bool
	var air_panas, kopi, gula bool
	air_panas = true
	kopi = true
	gula = true
	cangkir_kopi = air_panas == true && kopi == true && gula == true
	fmt.Println(cangkir_kopi)
}

func soal_3_bola_terurut() {
	// bola terurut dari angka paling kecil
	var bola_a, bola_b, bola_c int
	fmt.Scan(&bola_a, &bola_b, &bola_c)
	if bola_a < bola_b {
		if bola_b < bola_c {
			fmt.Println(bola_a, bola_b, bola_c)
		} else if bola_b > bola_c {
			if bola_c < bola_a {
				fmt.Println(bola_c, bola_a, bola_b)
			} else {
				fmt.Println(bola_a, bola_c, bola_b)
			}
		}
	} else if bola_a > bola_b {
		if bola_a < bola_c {
			fmt.Println(bola_b, bola_a, bola_c)
		} else if bola_a > bola_c {
			if bola_c < bola_b {
				fmt.Println(bola_c, bola_b, bola_a)
			} else {
				fmt.Println(bola_b, bola_c, bola_a)
			}
		}
	}
}

func soal_4_bola_terberat() {
	var bola_a, bola_b, bola_c, bola_d, max_berat int
	var bola_terberat string
	fmt.Scan(&bola_a, &bola_b, &bola_c, &bola_d)
	if bola_a > bola_b {
		max_berat = bola_a
		bola_terberat = "bola_a"
	} else {
		max_berat = bola_b
		bola_terberat = "bola_b"
	}
	if bola_c > max_berat {
		max_berat = bola_c
		bola_terberat = "bola_c"
	} else if bola_d > max_berat {
		max_berat = bola_d
		bola_terberat = "bola_d"
	}
	fmt.Println(bola_terberat, "dgn berat:", max_berat)
}

func soal_5_koper() {
	var n_penumpang int
	var berat_koper_n_penumpang, total_berat, rata_rata, data_berat_koper float64
	
	fmt.Scan(&n_penumpang)
	for range n_penumpang {
		fmt.Print("Masukkan berat koper: ")
		fmt.Scan(&berat_koper_n_penumpang)
		data_berat_koper += berat_koper_n_penumpang
	}
	rata_rata = total_berat / float64(n_penumpang)
	fmt.Println(rata_rata)
}

func main() {
	soal_1_lemari_baju()
	soal_2_membuat_kopi()
	soal_3_bola_terurut()
	soal_4_bola_terberat()
	soal_5_koper()
}
