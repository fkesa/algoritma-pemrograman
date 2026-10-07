package main

import "fmt"

func bilangan_ganjil(angka int) bool {
	if angka%2 == 1 {
		return true
	} else {
		return false
	}
}

func cumlaude(semester int, nilai_eprt int) bool {
	if semester <= 8 && nilai_eprt > 500 {
		return true
	} else {
		return false
	}
}

func tiga_digit_terurut_ke_kecil(angka int) bool {
	// 3 digit angkanya itu terurut dari besar ke kecil
	var digit_1, digit_2, digit_3 int
	digit_1 = angka / 100
	digit_2 = (angka / 10) % 10
	digit_3 = angka % 10
	if digit_1 > digit_2 && digit_2 > digit_3 {
		return true
	} else {
		return false
	}
}

func tiga_digit_terurut(angka int) bool {
	// 3 digit angkanya bisa terurut dari kecil/besar
	var digit_1, digit_2, digit_3 int
	digit_1 = angka / 100
	digit_2 = (angka / 10) % 10
	digit_3 = angka % 10

	if digit_1 > digit_2 && digit_2 > digit_3 {
		return true
	} else if digit_1 < digit_2 && digit_2 < digit_3 {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println(bilangan_ganjil(7))
	fmt.Println(bilangan_ganjil(10))
	fmt.Println(cumlaude(7, 520))
	fmt.Println(cumlaude(10, 573))
	fmt.Println(tiga_digit_terurut_ke_kecil(530))
	fmt.Println(tiga_digit_terurut_ke_kecil(555))
	fmt.Println(tiga_digit_terurut(149))
	fmt.Println(tiga_digit_terurut(555))
	fmt.Println(tiga_digit_terurut(961))
	fmt.Println(tiga_digit_terurut(183))
}
