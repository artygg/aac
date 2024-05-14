package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) Initialize(user, password, dbname string) {
	var err error

	a.DB, err = sql.Open("mysql", "user:password@/dbname")
	if err != nil {
		log.Fatal(err)
	}

	a.DB.SetConnMaxLifetime(0)
	a.DB.SetMaxOpenConns(10)
	a.DB.SetMaxIdleConns(10)

	a.Router = mux.NewRouter()
}

func (a *App) Run(addr string) {
	a.initializeRoutes()
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/api/device/student", a.compareFace).Methods("PUT")
	//a.Router.HandleFunc("/products", a.).Methods("GET")
	//a.Router.HandleFunc("/product", a.createProduct).Methods("POST")
	a.Router.HandleFunc("/product", a.getDevice).Methods("GET")
	//a.Router.HandleFunc("/product/{id:[0-9]+}", a.updateProduct).Methods("PUT")
	//a.Router.HandleFunc("/product/{id:[0-9]+}", a.deleteProduct).Methods("DELETE")
}

func (a *App) getDevice(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	d := Device{Id: id}
	if err := d.getDevice(a.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			w.WriteHeader(http.StatusNotFound)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	//w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Type", "text/html")
	_, err := w.Write([]byte("<h1>Sofronie LOH</h1>"))
	if err != nil {
		log.Fatal("error writing response")
		return
	}
}
