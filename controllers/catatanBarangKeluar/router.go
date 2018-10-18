package catatanBarangKeluar

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
	router.Get("/catatanBarangKeluar/upload", GetCatatanBarangKeluar)
	router.Post("/catatanBarangKeluar/input", PostUserRegistrationHandler)
	return router
}

func GetCatatanBarangKeluar() {
	db, err := readJSONFile("../data/catatanBarangKeluar.json")
	if nil != err {
		log.Fatalln(err)
	}

	f, err := os.Create("../data/catatanBarangKeluar.csv")
	if nil != err {
		log.Fatalln(err)
	}

	defer f.Close()

	w := csv.NewWriter(f)

	w.Write(types.GetHeader())
	for _, user := range db.Users {
		ss := user.EncodeAsStrings()
		w.Write(ss)
	}

	w.Flush()

	err = w.Error()
	if nil != err {
		log.Fatalln(err)
	}
}

func readJSONFile(s string) (db *controllers.CatatanBarangKeluarDb, err error) {
	f, err := os.Open(s)
	if nil != err {
		return nil, err
	}
	defer f.Close()

	var dec = json.NewDecoder(f)

	db = new(controllers.CatatanBarangKeluarDb)
	dec.Decode(db)

	return
}
