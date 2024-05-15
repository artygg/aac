package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) Initialize() {
	var err error

	a.DB, err = sql.Open("mysql", "u420565238_aas:^5qJ2ZVRgEO3@tcp(109.106.246.151)/u420565238_aas")
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
	//a.Router.HandleFunc("/api/device/student", a.compareFace).Methods("PUT")
	//a.Router.HandleFunc("/products", a.).Methods("GET")
	//a.Router.HandleFunc("/product", a.createProduct).Methods("POST")
	a.Router.HandleFunc("/product", a.getUser).Methods("GET")
	//a.Router.HandleFunc("/product/{id:[0-9]+}", a.updateProduct).Methods("PUT")
	//a.Router.HandleFunc("/product/{id:[0-9]+}", a.deleteProduct).Methods("DELETE")
}

func (a *App) compareFace(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	mac := vars["mac"]
	key := vars["key"]
	if mac == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	device := Device{Mac: mac}
	if err := device.getDevice(a.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			w.WriteHeader(http.StatusNotFound)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	if key == device.Key {
		w.WriteHeader(http.StatusAccepted)
		w.Header().Set("Content-Type", "text/html")
		_, err := w.Write([]byte("<h1>Sofronie LOH</h1>"))
		if err != nil {
			log.Fatal("error writing response")
			return
		}
	}
}

func (a *App) getUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var (
		id   int
		name string
	)
	id, _ = strconv.Atoi(vars["id"])
	rows, err := a.DB.Query("select id, name from users where id = ?", 1)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, name)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	_, err = w.Write([]byte(fmt.Sprintf("ID: %d, Name: %s\n", id, name)))
	if err != nil {
		return
	}
}

//func (a *App) compareFace(w http.ResponseWriter, r *http.Request) {
//	t := time.Now()
//
//	file, handler, err := r.FormFile("image")
//	fileName := r.FormValue("file_name")
//	if err != nil {
//		panic(err)
//	}
//	defer file.Close()
//
//	f, err := os.OpenFile(handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
//	if err != nil {
//		panic(err)
//	}
//	defer f.Close()
//	_, _ = io.WriteString(w, "File "+fileName+" Uploaded successfully")
//	_, _ = io.Copy(f, file)
//
//}
