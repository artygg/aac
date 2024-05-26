package main

import (
	"context"
	"crypto/rand"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
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
	a.Router.Handle("/api/device/upload", a.deviceAuthMiddleware(http.HandlerFunc(a.uploadImage))).Methods("POST")
	a.Router.HandleFunc("/api/web/login", a.loginHandler).Methods("POST")

	a.Router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./platform/index.html")
	})

	a.Router.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./platform/login.html")
	})

	a.Router.Handle("/protected", a.userAuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./platform/protected.html")
	})))

	a.Router.HandleFunc("/logout", a.logoutHandler)
	//a.Router.Handle("/api/web/courses", a.userAuthMiddleware(http.HandlerFunc(a.getCourses))).Methods("GET")
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
func (a *App) loginHandler(w http.ResponseWriter, r *http.Request) {
	teacher := Teacher{}
	err := json.NewDecoder(r.Body).Decode(&teacher)
	log.Println(teacher)
	password := teacher.Password
	if teacher.Email == "" || teacher.Password == "" {
		http.Error(w, "Login and password are required", http.StatusBadRequest)
		return
	}

	if err := teacher.getTeacher(a.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			http.Error(w, "Invalid login", http.StatusUnauthorized)
		default:
			log.Println(err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}

	if teacher.Password != password {
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		return
	}

	// Create a session
	session, _ := a.Store.Get(r, "aas-user")
	session.Values["authenticated"] = true
	session.Values["user"] = teacher.Email
	session.Values["userID"] = teacher.Id
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}

	http.Redirect(w, r, "/protected", http.StatusFound)
}

func (a *App) logoutHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := a.Store.Get(r, "aas-user")
	session.Values["authenticated"] = false
	delete(session.Values, "user")
	delete(session.Values, "userID")
	session.Save(r, w)

	http.Redirect(w, r, "/", http.StatusFound)
}
func (a *App) checkDevice(w http.ResponseWriter, r *http.Request) {
	device, ok := r.Context().Value("device").(Device)
	if !ok {
		http.Error(w, "Device information not found", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Device MAC: %s, Device Key: %s, Room: %s", device.Mac, device.Key, device.Room)
}
func (a *App) uploadImage(w http.ResponseWriter, r *http.Request) {
	device, ok := r.Context().Value("device").(Device)
	if !ok {
		http.Error(w, "Device information not found", http.StatusInternalServerError)
		return
	}
	err := r.ParseMultipartForm(10 << 20) // 10 MB maximum file size
	if err != nil {
		fmt.Println("Error parsing form:", err)
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	// Get the file from the form data
	file, handler, err := r.FormFile("image")
	if err != nil {
		fmt.Println("Error retrieving file:", err)
		http.Error(w, "Error retrieving file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Create a new file on the server to store the uploaded image
	f, err := os.OpenFile("./uploads/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("Error creating file:", err)
		http.Error(w, "Error creating file", http.StatusInternalServerError)
		return
	}
	defer f.Close()

	// Copy the uploaded file to the new file on the server
	_, err = io.Copy(f, file)
	if err != nil {
		fmt.Println("Error copying file:", err)
		http.Error(w, "Error copying file", http.StatusInternalServerError)
		return
	}

	// Send a success response
	fmt.Fprintf(w, "File uploaded successfully: %s", handler.Filename)
	fmt.Fprintf(w, "Device MAC: %s, Device Key: %s, Room: %s", device.Mac, device.Key, device.Room)
}

func (a *App) protectedHandler(w http.ResponseWriter, r *http.Request) {

	session, _ := a.Store.Get(r, "aas-user")

	login := session.Values["user"]
	userID := session.Values["userID"]

	fmt.Fprintf(w, "Hello, %s! Your user ID is %d. This is a protected route.", login, userID)
}

func (a *App) getCourses(w http.ResponseWriter, r *http.Request) {
	session, err := a.Store.Get(r, "aas-user")
	if err != nil {
		http.Error(w, "Failed to get session", http.StatusInternalServerError)
		log.Println("Error retrieving session:", err)
		return
	}

	authenticated, ok := session.Values["authenticated"].(bool)
	if !ok || !authenticated {
		http.Error(w, "Forbidden", http.StatusForbidden)
		log.Println("User not authenticated")
		return
	}

	teacherID, ok := session.Values["userID"].(int)
	if !ok {
		http.Error(w, "Failed to get user ID from session", http.StatusInternalServerError)
		log.Println("Error retrieving user ID from session")
		return
	}

	teacher := Teacher{Id: teacherID}

	err = teacher.getCoursesByTeacher(a.DB)
	if err != nil {
		http.Error(w, "Failed to get courses", http.StatusInternalServerError)
		log.Println("Error retrieving courses:", err)
		return
	}

	log.Printf("Retrieved courses: %+v\n", teacher.Courses)

	response := map[string]interface{}{
		"user_id": teacherID,
		"courses": teacher.Courses,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode courses", http.StatusInternalServerError)
		log.Println("Error encoding courses to JSON:", err)
	}
}

func (a *App) getGroups(w http.ResponseWriter, r *http.Request) {
	session, err := a.Store.Get(r, "aas-user")
	if err != nil {
		http.Error(w, "Failed to get session", http.StatusInternalServerError)
		log.Println("Error retrieving session:", err)
		return
	}

	authenticated, ok := session.Values["authenticated"].(bool)
	if !ok || !authenticated {
		http.Error(w, "Forbidden", http.StatusForbidden)
		log.Println("User not authenticated")
		return
	}

	groupCluster := GroupCluster{}

	err = groupCluster.getGroups(a.DB)
	if err != nil {
		http.Error(w, "Failed to get courses", http.StatusInternalServerError)
		log.Println("Error retrieving courses:", err)
		return
	}

	log.Printf("Retrieved courses: %+v\n", groupCluster.Groups)

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(groupCluster.Groups); err != nil {
		http.Error(w, "Failed to encode courses", http.StatusInternalServerError)
		log.Println("Error encoding courses to JSON:", err)
	}
}

func (a *App) getAttendences(w http.ResponseWriter, r *http.Request) {
	session, err := a.Store.Get(r, "aas-user")
	if err != nil {
		http.Error(w, "Failed to get session", http.StatusInternalServerError)
		log.Println("Error retrieving session:", err)
		return
	}

	authenticated, ok := session.Values["authenticated"].(bool)
	if !ok || !authenticated {
		http.Error(w, "Forbidden", http.StatusForbidden)
		log.Println("User not authenticated")
		return
	}

	vars := mux.Vars(r)
	classIDStr := vars["class_id"]
	classID, err := strconv.Atoi(classIDStr)
	if err != nil {
		http.Error(w, "Invalid class ID", http.StatusBadRequest)
		log.Println("Invalid class ID:", err)
		return
	}

	class := Class{Id: classID}

	err = class.getAttendencesByClass(a.DB)
	if err != nil {
		http.Error(w, "Failed to get attendances", http.StatusInternalServerError)
		log.Println("Error retrieving attendances:", err)
		return
	}

	log.Printf("Retrieved attendances: %+v\n", class.Attendances)

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(class.Attendances); err != nil {
		http.Error(w, "Failed to encode attendances", http.StatusInternalServerError)
		log.Println("Error encoding attendances to JSON:", err)
	}
}
