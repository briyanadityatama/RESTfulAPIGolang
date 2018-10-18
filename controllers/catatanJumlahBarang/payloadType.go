package catatanJumlahBarang

// Ref: https://golang.org/pkg/encoding/csv/

import (
	"encoding/xml"
	"strconv"
)

var headerJumlahBarang = []string{`sku`, `nama_item`, `jumlah_sekarang`}

type (
	CatatanJumlahBarang struct {
		SKU             string `json:"sku" xml:"sku"`
		Nama_Item       string `json:"nama_item" xml:"nama_item"`
		Jumlah_Sekarang int    `json:"jumlah_sekarang" xml:"jumlah_sekarang"`
	}

	catatanJumlahBarangDb struct {
		XMLName              xml.Name              `json:"-" xml:"catatanJumlahBarangs"`
		Type                 string                `json:"type,omitempty" xml:"type"`
		CatatanJumlahBarangs []CatatanJumlahBarang `json:"catatanJumlahBarangs,omitempty" xml:"catatanJumlahBarang"`
	}
)

func GetHeaderJumlahBarang() []string {
	return headerJumlahBarang
}

func (catatanJumlahBarang CatatanJumlahBarang) EncodeAsStrings() (ss []string) {
	ss = make([]string, 3)
	ss[0] = catatanJumlahBarang.SKU
	ss[1] = catatanJumlahBarang.Nama_Item
	ss[2] = strconv.Itoa(catatanJumlahBarang.Jumlah_Sekarang)
	return
}

func (catatanJumlahBarang *CatatanJumlahBarang) FromCSV(ss []string) {
	if nil == catatanJumlahBarang {
		return
	}

	if nil == ss {
		return
	}

	catatanJumlahBarang.SKU = ss[0]
	catatanJumlahBarang.Nama_Item = ss[1]
	catatanJumlahBarang.Jumlah_Sekarang, _ = strconv.Atoi(ss[2])
}
