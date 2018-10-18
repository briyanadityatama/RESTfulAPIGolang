package main

import (
	"fmt"
	"net/http"

	"github.com/briyanadityatama/RESTfulAPIGolang/controllers/catatanBarangKeluar"
	"github.com/briyanadityatama/RESTfulAPIGolang/controllers/catatanBarangMasuk"
	"github.com/briyanadityatama/RESTfulAPIGolang/controllers/catatanJumlahBarang"
	"github.com/briyanadityatama/RESTfulAPIGolang/controllers/laporanNilaiBarang"
	"github.com/briyanadityatama/RESTfulAPIGolang/controllers/laporanPenjualan"
	"github.com/go-chi/chi"
)

// APIServer ...
type APIServer struct {
	Router *chi.Mux
}

// Initialize , initialize api server
func (app *APIServer) Initialize() {
	app.Router = chi.NewRouter()
	catatanBarangKeluarRouter := catatanBarangKeluar.Routers()
	catatanBarangMasukRouter := catatanBarangMasuk.Routers()
	catatanJumlahBarangRouter := catatanJumlahBarang.Routers()
	laporanNilaiBarangRouter := laporanNilaiBarang.Routers()
	laporanPenjualanRouter := laporanPenjualan.Routers()

	app.Router.Mount("/catatanBarangKeluar", catatanBarangKeluarRouter)
	app.Router.Mount("/catatanBarangMasuk", catatanBarangMasukRouter)
	app.Router.Mount("/catatanJumlahBarang", catatanJumlahBarangRouter)
	app.Router.Mount("/laporanNilaiBarang", laporanNilaiBarangRouter)
	app.Router.Mount("/laporanPenjualan", laporanPenjualanRouter)
}

// Run , run api server
func (app *APIServer) Run(addr string) {
	err := http.ListenAndServe(addr, app.Router)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Printf("Server runs on port %s\n", addr)
}

func main() {
	apiServer := APIServer{}
	apiServer.Initialize()
	apiServer.Run(":8080")
}
