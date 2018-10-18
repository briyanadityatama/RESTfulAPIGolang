package laporanNilaiBarang

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
	router.Get("/laporanNilaiBarang/upload", GetLaporanNilaiBarang)
	router.Post("/laporanNilaiBarang/input", PostLaporanNilaiBarang)
	return router
}

func GetLaporanNilaiBarang() {
	db, err := readJSONFile("../data/LaporanNilaiBarang.json")
	if nil != err {
		log.Fatalln(err)
	}

	f, err := os.Create("../data/LaporanNilaiBarang.csv")
	if nil != err {
		log.Fatalln(err)
	}

	defer f.Close()

	w := csv.NewWriter(f)

	w.Write(types.GetHeader())
	for _, laporanNilaiBarang := range db.LaporanNilaiBarang {
		ss := laporanNilaiBarang.EncodeAsStrings()
		w.Write(ss)
	}

	w.Flush()

	err = w.Error()
	if nil != err {
		log.Fatalln(err)
	}
}

func readJSONFile(s string) (db *controllers.LaporanNilaiBarangDb, err error) {
	f, err := os.Open(s)
	if nil != err {
		return nil, err
	}
	defer f.Close()

	var dec = json.NewDecoder(f)

	db = new(controllers.LaporanNilaiBarangDb)
	dec.Decode(db)

	return
}
