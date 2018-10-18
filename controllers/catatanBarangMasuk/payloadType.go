package catatanBarangMasuk

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

var headerBarangMasuk = []string{`waktu`, `sku`, `nama_barang`, `jumlah_pemesanan`, `jumlah_diterima`, `harga_beli`, `total`, `nomoe_kwitansi`, `catatan`}

type (
	CatatanBarangMasuk struct {
		Waktu            time.Time `json:"waktu" xml:"waktu,attr"`
		SKU              string    `json:"sku" xml:"sku"`
		Nama_Barang      string    `json:"nama_barang" xml:"nama_barang"`
		Jumlah_Pemesanan int       `json:"jumlah_pemesanan" xml:"jumlah_pemesanan"`
		Jumlah_Diterima  int       `json:"jumlah_diterima" xml:"jumlah_diterima"`
		Harga_Beli       int       `json:"harga_beli" xml:"harga_beli"`
		Total            int       `json:"total" xml:"total"`
		Nomor_Kwitansi   string    `json:"nomor_kwitansi" xml:"nomor_kwitansi"`
		Catatan          string    `json:"catatan" xml:"catatan"`
	}

	catatanBarangMasukDb struct {
		XMLName            xml.Name             `json:"-" xml:"catatanBarangMasuk"`
		Type               string               `json:"type,omitempty" xml:"type"`
		CatatanBarangMasuk []CatatanBarangMasuk `json:"catatanBarangMasuk,omitempty" xml:"catatanBarangMasuk"`
	}
)

func GetHeaderBarangMasuk() []string {
	return headerBarangMasuk
}

func (catatanBarangMasuk CatatanBarangMasuk) EncodeAsStrings() (ss []string) {
	ss = make([]string, 9)
	ss[0] = time.Time.String(catatanBarangMasuk.Waktu)
	ss[1] = catatanBarangMasuk.SKU
	ss[2] = catatanBarangMasuk.Nama_Barang
	ss[3] = strconv.Itoa(catatanBarangMasuk.Jumlah_Pemesanan)
	ss[4] = strconv.Itoa(catatanBarangMasuk.Jumlah_Diterima)
	ss[5] = strconv.Itoa(catatanBarangMasuk.Harga_Beli)
	ss[6] = strconv.Itoa(catatanBarangMasuk.Total)
	ss[7] = catatanBarangMasuk.Nomor_Kwitansi
	ss[8] = catatanBarangMasuk.Catatan
	return
}

func (catatanBarangMasuk *CatatanBarangMasuk) FromCSV(ss []string) {
	if nil == catatanBarangMasuk {
		return
	}

	if nil == ss {
		return
	}

	catatanBarangMasuk.Waktu, _ = time.Parse(time.RFC3339, ss[0])
	catatanBarangMasuk.SKU = ss[1]
	catatanBarangMasuk.Nama_Barang = ss[2]
	catatanBarangMasuk.Jumlah_Pemesanan, _ = strconv.Atoi(ss[3])
	catatanBarangMasuk.Jumlah_Diterima, _ = strconv.Atoi(ss[4])
	catatanBarangMasuk.Harga_Beli, _ = strconv.Atoi(ss[5])
	catatanBarangMasuk.Total, _ = strconv.Atoi(ss[6])
	catatanBarangMasuk.Nomor_Kwitansi = ss[7]
	catatanBarangMasuk.Catatan = ss[8]
}
