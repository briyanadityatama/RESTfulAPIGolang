package laporanPenjualan

import (
	"encoding/csv"
	"encoding/json"
	"net/http"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/go-chi/chi"
)

// Routers , return router constructed for auth purpose
func Routers() http.Handler {
	router := chi.NewRouter()
	router.Get("/laporanPenjualan/upload", GetLaporanPenjualan)
	router.Post("/laporanPenjualan/input", PostLaporanPenjualan)
	return router
}

func GetLaporanPenjualan() {
	db, err := readJSONFile("../data/LaporanPenjualan.json")
	if nil != err {
		log.Fatalln(err)
	}

	f, err := os.Create("../data/LaporanPenjualan.csv")
	if nil != err {
		log.Fatalln(err)
	}

	defer f.Close()

	w := csv.NewWriter(f)

	w.Write(types.GetHeader())
	for _, laporanPenjualan := range db.LaporanPenjualan {
		ss := laporanPenjualan.EncodeAsStrings()
		w.Write(ss)
	}

	w.Flush()

	err = w.Error()
	if nil != err {
		log.Fatalln(err)
	}
}

func readJSONFile(s string) (db *controllers.LaporanPenjualanDb, err error) {
	f, err := os.Open(s)
	if nil != err {
		return nil, err
	}
	defer f.Close()

	var dec = json.NewDecoder(f)

	db = new(controllers.LaporanPenjualanDb)
	dec.Decode(db)

	return
}
