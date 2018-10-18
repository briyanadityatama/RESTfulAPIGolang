package types

// Ref: https://golang.org/pkg/encoding/csv/

import (
	"encoding/xml"
	"strconv"
	"time"
)

// const dateFormat = "2006-01-02 15:04:05"

// type MyTime struct {
// 	time.Time
// }

// func (m *MyTime) UnmarshalJSON(p []byte) error {
// 	t, err := time.Parse(dateFormat, strings.Replace(
// 		string(p),
// 		"\"",
// 		"",
// 		-1,
// 	))

// 	if err != nil {
// 		return err
// 	}

// 	m.Time = t

// 	return nil
// }

var headerPenjualan = []string{`id_pesanan`, `Waktu`, `sku`, `nama_barang`, `jumlah`, `harga_jual`, `total`, `harga_beli`, `laba`}

type (
	LaporanPenjualan struct {
		ID_Pesanan  int       `json:"id_pesanan" xml:"id_pesanan,attr"`
		Waktu       time.Time `json:"waktu" xml:"waktu,attr"`
		SKU         string    `json:"sku" xml:"sku"`
		Nama_Barang string    `json:"nama_barang" xml:"nama_barang"`
		Jumlah      int       `json:"jumlah" xml:"jumlah"`
		Harga_Jual  int       `json:"harga_jual" xml:"harga_jual"`
		Total       int       `json:"total" xml:"total"`
		Harga_Beli  int       `json:"harga_beli" xml:"harga_beli"`
		Laba        int       `json:"laba" xml:"laba"`
	}

	laporanPenjualanDb struct {
		XMLName           xml.Name           `json:"-" xml:"laporanPenjualans"`
		Type              string             `json:"type,omitempty" xml:"type"`
		LaporanPenjualans []LaporanPenjualan `json:"laporanPenjualans,omitempty" xml:"laporanPenjualan"`
	}
)

func GetHeaderPenjualan() []string {
	return headerPenjualan
}

func (laporanPenjualan LaporanPenjualan) EncodeAsStrings() (ss []string) {
	ss = make([]string, 9)
	ss[0] = strconv.Itoa(laporanPenjualan.ID_Pesanan)
	ss[1] = time.Time.String(laporanPenjualan.Waktu)
	ss[2] = laporanPenjualan.SKU
	ss[3] = laporanPenjualan.Nama_Barang
	ss[4] = strconv.Itoa(laporanPenjualan.Jumlah)
	ss[5] = strconv.Itoa(laporanPenjualan.Harga_Jual)
	ss[6] = strconv.Itoa(laporanPenjualan.Total)
	ss[7] = strconv.Itoa(laporanPenjualan.Harga_Beli)
	ss[8] = strconv.Itoa(laporanPenjualan.Laba)
	return
}

func (laporanPenjualan *LaporanPenjualan) FromCSV(ss []string) {
	if nil == laporanPenjualan {
		return
	}

	if nil == ss {
		return
	}

	laporanPenjualan.ID_Pesanan, _ = strconv.Atoi(ss[0])
	laporanPenjualan.Waktu, _ = time.Parse(time.RFC3339, ss[1])
	laporanPenjualan.SKU = ss[2]
	laporanPenjualan.Nama_Barang = ss[3]
	laporanPenjualan.Jumlah, _ = strconv.Atoi(ss[4])
	laporanPenjualan.Harga_Jual, _ = strconv.Atoi(ss[5])
	laporanPenjualan.Total, _ = strconv.Atoi(ss[6])
	laporanPenjualan.Harga_Beli, _ = strconv.Atoi(ss[7])
	laporanPenjualan.Laba, _ = strconv.Atoi(ss[8])
}
