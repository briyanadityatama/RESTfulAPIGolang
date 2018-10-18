package catatanBarangMasuk

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
	router.Get("/catatanBarangMasuk/upload", GetCatatanBarangMasuk)
	router.Post("/catatanBarangMasuk/input", PostCatatanBarangMasuk)
	return router
}

func GetCatatanBarangKeluar() {
	db, err := readJSONFile("../data/CatatanBarangMasuk.json")
	if nil != err {
		log.Fatalln(err)
	}

	f, err := os.Create("../data/CatatanBarangMasuk.csv")
	if nil != err {
		log.Fatalln(err)
	}

	defer f.Close()

	w := csv.NewWriter(f)

	w.Write(types.GetHeader())
	for _, catatanBarangMasuk := range db.CatatanBarangMasuk {
		ss := catatanBarangMasuk.EncodeAsStrings()
		w.Write(ss)
	}

	w.Flush()

	err = w.Error()
	if nil != err {
		log.Fatalln(err)
	}
}

func readJSONFile(s string) (db *controllers.CatatanBarangMasukDb, err error) {
	f, err := os.Open(s)
	if nil != err {
		return nil, err
	}
	defer f.Close()

	var dec = json.NewDecoder(f)

	db = new(controllers.CatatanBarangMasukDb)
	dec.Decode(db)

	return
}
