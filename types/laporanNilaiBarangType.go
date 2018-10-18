package types

// Ref: https://golang.org/pkg/encoding/csv/

import (
	"encoding/xml"
	"strconv"
)

var headerNilaiBarang = []string{`sku`, `nama_item`, `jumlah`, `rata_rata_harga_beli`, `total`}

type (
	LaporanNilaiBarang struct {
		SKU                  string `json:"sku" xml:"sku"`
		Nama_Item            string `json:"nama_item" xml:"nama_item"`
		Jumlah               int    `json:"jumlah" xml:"jumlah"`
		Rata_Rata_Harga_Beli int    `json:"rata_rata_harga_beli" xml:"rata_rata_harga_beli"`
		Total                int    `json:"total" xml:"total"`
	}

	laporanNilaiBarangDb struct {
		XMLName             xml.Name             `json:"-" xml:"catatanJumlahBarangs"`
		Type                string               `json:"type,omitempty" xml:"type"`
		LaporanNilaiBarangs []LaporanNilaiBarang `json:"laporanNilaiBarangs,omitempty" xml:"laporanNilaiBarang"`
	}
)

func GetHeaderNilaiBarang() []string {
	return headerNilaiBarang
}

func (laporanNilaiBarang LaporanNilaiBarang) EncodeAsStrings() (ss []string) {
	ss = make([]string, 5)
	ss[0] = laporanNilaiBarang.SKU
	ss[1] = laporanNilaiBarang.Nama_Item
	ss[2] = strconv.Itoa(laporanNilaiBarang.Jumlah)
	ss[3] = strconv.Itoa(laporanNilaiBarang.Rata_Rata_Harga_Beli)
	ss[4] = strconv.Itoa(laporanNilaiBarang.Total)
	return
}

func (laporanNilaiBarang *LaporanNilaiBarang) FromCSV(ss []string) {
	if nil == laporanNilaiBarang {
		return
	}

	if nil == ss {
		return
	}

	laporanNilaiBarang.SKU = ss[0]
	laporanNilaiBarang.Nama_Item = ss[1]
	laporanNilaiBarang.Jumlah, _ = strconv.Atoi(ss[2])
	laporanNilaiBarang.Rata_Rata_Harga_Beli, _ = strconv.Atoi(ss[3])
	laporanNilaiBarang.Total, _ = strconv.Atoi(ss[4])
}
