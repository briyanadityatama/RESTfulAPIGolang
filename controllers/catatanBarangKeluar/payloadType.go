package catatanBarangKeluar

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

var headerBarangKeluar = []string{`waktu`, `sku`, `nama_barang`, `jumlah_keluar`, `harga_jual`, `total`, `catatan`}

type (
	CatatanBarangKeluar struct {
		Waktu         time.Time `json:"waktu" xml:"waktu,attr"`
		SKU           string    `json:"sku" xml:"sku"`
		Nama_Barang   string    `json:"nama_barang" xml:"nama_barang"`
		Jumlah_Keluar int       `json:"jumlah_keluar" xml:"jumlah_keluar"`
		Harga_Jual    int       `json:"harga_jual" xml:"harga_jual"`
		Total         int       `json:"total" xml:"total"`
		Catatan       string    `json:"catatan" xml:"catatan"`
	}

	CatatanBarangKeluarDb struct {
		XMLName              xml.Name              `json:"-" xml:"catatanBarangKeluars"`
		Type                 string                `json:"type,omitempty" xml:"type"`
		CatatanBarangKeluars []CatatanBarangKeluar `json:"catatanBarangKeluars,omitempty" xml:"catatanBarangKeluar"`
	}
)

func GetHeaderBarangKeluar() []string {
	return headerBarangKeluar
}

func (catatanBarangKeluar CatatanBarangKeluar) EncodeAsStrings() (ss []string) {
	ss = make([]string, 7)
	ss[0] = time.Time.String(catatanBarangKeluar.Waktu)
	ss[1] = catatanBarangKeluar.SKU
	ss[2] = catatanBarangKeluar.Nama_Barang
	ss[3] = strconv.Itoa(catatanBarangKeluar.Jumlah_Keluar)
	ss[4] = strconv.Itoa(catatanBarangKeluar.Harga_Jual)
	ss[5] = strconv.Itoa(catatanBarangKeluar.Total)
	ss[6] = catatanBarangKeluar.Catatan
	return
}

func (catatanBarangKeluar *CatatanBarangKeluar) FromCSV(ss []string) {
	if nil == catatanBarangKeluar {
		return
	}

	if nil == ss {
		return
	}

	catatanBarangKeluar.Waktu, _ = time.Parse(time.RFC3339, ss[0])
	catatanBarangKeluar.SKU = ss[1]
	catatanBarangKeluar.Nama_Barang = ss[2]
	catatanBarangKeluar.Jumlah_Keluar, _ = strconv.Atoi(ss[3])
	catatanBarangKeluar.Harga_Jual, _ = strconv.Atoi(ss[4])
	catatanBarangKeluar.Total, _ = strconv.Atoi(ss[5])
	catatanBarangKeluar.Catatan = ss[6]
}
