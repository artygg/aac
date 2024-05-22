package main

import (
	"context"
	"crypto/rand"
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"log"
	"net/http"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
	Store  *sessions.CookieStore
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

	err = a.generateKey()
	if err != nil {
		a.Store = sessions.NewCookieStore([]byte("Stan0dard0101Coo6kie0101Sto7reByAAS"))
	}
}

func (a *App) generateKey() error {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		return err
	}
	a.Store = sessions.NewCookieStore(key)
	return nil
}

func (a *App) Run(addr string) {
	a.initializeRoutes()
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a *App) initializeRoutes() {
	a.Router.Handle("/api/device/check", a.deviceAuthMiddleware(http.HandlerFunc(a.checkDevice))).Methods("GET")

	//a.Router.HandleFunc("/products", a.).Methods("GET")
	//a.Router.HandleFunc("/product", a.createProduct).Methods("POST")
	//a.Router.HandleFunc("/product", a.getUser).Methods("GET")
	a.Router.Handle("/api/web/courses", a.userAuthMiddleware(http.HandlerFunc(a.getCourses))).Methods("GET")
	//a.Router.HandleFunc("/product/{id:[0-9]+}", a.updateProduct).Methods("PUT")
	//a.Router.HandleFunc("/product/{id:[0-9]+}", a.deleteProduct).Methods("DELETE")
}
func (a *App) deviceAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		device := Device{}
		device.Mac = r.Header.Get("X-MAC-ADDRESS")
		key := r.Header.Get("X-API-KEY")

		if key == "" {
			http.Error(w, "API key is missing", http.StatusUnauthorized)
			return
		}

		if err := device.getDevice(a.DB); err != nil {
			switch err {
			case sql.ErrNoRows:
				http.Error(w, "Device wasn't found", http.StatusNotFound)
			default:
				log.Fatalln(err)
				w.WriteHeader(http.StatusInternalServerError)
			}
			return
		}

		if device.Key != key {
			http.Error(w, "Invalid API key", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "device", device)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
func (a *App) userAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, _ := a.Store.Get(r, "aas-user")
		authenticated, ok := session.Values["authenticated"].(bool)

		if !ok || !authenticated {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}
func (a *App) checkDevice(w http.ResponseWriter, r *http.Request) {
	device, ok := r.Context().Value("device").(Device)
	if !ok {
		http.Error(w, "Device information not found", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Device MAC: %s, Device Key: %s, Room: %s", device.Mac, device.Key, device.Room)
}
func (a *App) compareFace(w http.ResponseWriter, r *http.Request) {
	device := Device{}
	err := json.NewDecoder(r.Body).Decode(&device)
	log.Println(device)
	key := device.Key
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if device.Mac == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if device.Key == "" {
		w.WriteHeader(http.StatusNetworkAuthenticationRequired)
		return
	}
	if err := device.getDevice(a.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			w.WriteHeader(http.StatusNotFound)
		default:

			log.Fatalln(err)
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	log.Println(key)
	log.Println(device.Key)
	if key == device.Key {
		w.WriteHeader(http.StatusAccepted)
		w.Header().Set("Content-Type", "text/html")
		_, err := w.Write([]byte("<h1>Sofronie LOH</h1>"))
		if err != nil {
			log.Fatal("error writing response")
			return
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

func (a *App) getCourses(w http.ResponseWriter, r *http.Request) {
	courses := []Course{}
	teacher.getCourse(a.DB)
}
