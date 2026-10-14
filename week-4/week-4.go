package main

import "fmt"

type d_transaksi struct {
	nama_barang   string
	jumlah_barang int
	harga_satuan  float64
	total_harga   float64
}

type d_persegi_panjang struct {
	lebar    float64
	panjang  float64
	keliling float64
	luas     float64
}

type d_bmi struct {
	nama      string
	berat     float64
	tinggi    float64
	nilai_bmi float64
}

type d_gaji_karyawan struct {
	nama       string
	gaji_pokok float64
	tunjangan  float64
	potongan   float64
	total_gaji float64
}

type d_investasi_saham struct {
	harga_beli        float64
	harga_jual        float64
	jumlah_saham      float64
	investasi_awal    float64
	total_penjualan   float64
	keuntungan_kotor  float64
	biaya_transaksi   float64
	pajak_keuntungan  float64
	keuntungan_bersih float64
}

func transaksi() {
	var t d_transaksi
	fmt.Scan(&t.harga_satuan, &t.jumlah_barang)
	t.total_harga = t.harga_satuan * float64(t.jumlah_barang)
	fmt.Println(t.total_harga)
}

func persegi_panjang() {
	var pp d_persegi_panjang
	fmt.Scan(&pp.panjang, &pp.lebar)
	pp.luas = pp.panjang * pp.lebar
	pp.keliling = 2 * (pp.panjang + pp.lebar)
	fmt.Println(pp.luas, pp.keliling)
}

func bmi() {
	var b d_bmi
	fmt.Scan(&b.nama, &b.berat, &b.tinggi)
	b.nilai_bmi = b.berat / (b.tinggi * b.tinggi)
	fmt.Println(b.nilai_bmi)
}

func gaji_karyawan() {
	var gk d_gaji_karyawan
	fmt.Scan(&gk.nama, &gk.gaji_pokok, &gk.tunjangan, &gk.potongan)
	gk.total_gaji = gk.gaji_pokok + gk.tunjangan + gk.potongan
	fmt.Println(gk.total_gaji)
}

func investasi_saham() {
	var ivs d_investasi_saham
	fmt.Scan(&ivs.harga_beli, &ivs.harga_jual, &ivs.jumlah_saham)
	ivs.investasi_awal = ivs.harga_beli * ivs.jumlah_saham
	ivs.total_penjualan = ivs.harga_jual * ivs.jumlah_saham
	ivs.keuntungan_kotor = ivs.total_penjualan - ivs.investasi_awal
	ivs.biaya_transaksi = 2.0 / 1000.0 * ivs.total_penjualan
	// ivs.pajak_keuntungan = math.Max(0, 10.0 / 100.0 * ivs.keuntungan_kotor) // Teknik lain
	if ivs.keuntungan_kotor > 0 {
        ivs.pajak_keuntungan = 10.0 / 100.0 * ivs.keuntungan_kotor
    } else {
        ivs.pajak_keuntungan = 0
    }
	ivs.keuntungan_bersih = ivs.keuntungan_kotor - ivs.biaya_transaksi - ivs.pajak_keuntungan
	fmt.Println(ivs.keuntungan_bersih)
}

func main() {
	investasi_saham()
}
