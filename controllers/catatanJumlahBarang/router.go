package catatanJumlahBarang

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
	router.Get("/catatanJumlahBarang/upload", GetCatatanJumlahBarang)
	router.Post("/catatanJumlahBarang/input", PostCatatanJumlahBarang)
	return router
}

func GetCatatanJumlahBarang() {
	db, err := readJSONFile("../data/CatatanJumlahBarang.json")
	if nil != err {
		log.Fatalln(err)
	}

	f, err := os.Create("../data/CatatanJumlahBarang.csv")
	if nil != err {
		log.Fatalln(err)
	}

	defer f.Close()

	w := csv.NewWriter(f)

	w.Write(types.GetHeader())
	for _, catatanJumlahBarang := range db.CatatanJumlahBarang {
		ss := catatanJumlahBarang.EncodeAsStrings()
		w.Write(ss)
	}

	w.Flush()

	err = w.Error()
	if nil != err {
		log.Fatalln(err)
	}
}

func readJSONFile(s string) (db *controllers.CatatanJumlahBarangDb, err error) {
	f, err := os.Open(s)
	if nil != err {
		return nil, err
	}
	defer f.Close()

	var dec = json.NewDecoder(f)

	db = new(controllers.CatatanJumlahBarangDb)
	dec.Decode(db)

	return
}
