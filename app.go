package main

import (
	"bytes"
	"context"
	"crypto/rand"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"time"

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
	a.DB.SetMaxOpenConns(0)
	a.DB.SetMaxIdleConns(0)

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

	a.Router.Handle("/api/web/classes", a.userAuthMiddleware(http.HandlerFunc(a.getClassesByCourseID))).Methods("GET")
	a.Router.Handle("/api/web/courses", a.userAuthMiddleware(http.HandlerFunc(a.getCoursesByTeacherID))).Methods("GET")
	a.Router.Handle("/api/web/groups", a.userAuthMiddleware(http.HandlerFunc(a.getAllGroups))).Methods("GET")
	a.Router.Handle("/api/web/groups/by_course", a.userAuthMiddleware(http.HandlerFunc(a.getGroupsByCourseID))).Methods("GET")

	a.Router.Handle("/api/web/attendance/by_class", a.userAuthMiddleware(http.HandlerFunc(a.getAttendencesByClassID))).Methods("GET")
	a.Router.Handle("/api/web/attendance/by_course", a.userAuthMiddleware(http.HandlerFunc(a.getAttendanceByCourseID))).Methods("GET")

	a.Router.Handle("/api/web/attendance", a.userAuthMiddleware(http.HandlerFunc(a.updateAttendanceStatus))).Methods("POST")

	a.Router.Handle("/api/web/course", a.userAuthMiddleware(http.HandlerFunc(a.createCourse))).Methods("POST")
	a.Router.Handle("/api/web/teacher", http.HandlerFunc(a.registerTeacher)).Methods("POST")
	a.initializeClient()
}

func (a *App) initializeClient() {
	a.Router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./platform/index.html")
	})

	a.Router.Handle("/courses", a.userAuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./platform/courses.html")
	})))

	a.Router.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./platform/login.html")
	})

	a.Router.Handle("/protected", a.userAuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./platform/protected.html")
	})))

	a.Router.Handle("/classes", a.userAuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./platform/classes.html")
	})))

	a.Router.Handle("/attendance/edit", a.userAuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./platform/updateAttendance.html")
	})))

	a.Router.Handle("/course/create", a.userAuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./platform/newcourse.html")
	})))

	a.Router.HandleFunc("/logout", a.logoutHandler)
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

	if err := teacher.get(a.DB); err != nil {
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
	//class, err := device.getClass(a.DB)
	//if err != nil {
	//	switch err {
	//	case sql.ErrNoRows:
	//		http.Error(w, "No Active Class for this room", http.StatusServiceUnavailable)
	//	default:
	//		log.Println(err)
	//		http.Error(w, "Internal server error", http.StatusInternalServerError)
	//	}
	//	return
	//}

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		fmt.Println("Error parsing form:", err)
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	file, handler, err := r.FormFile("image")
	if err != nil {
		fmt.Println("Error retrieving file:", err)
		http.Error(w, "Error retrieving file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	f, err := os.OpenFile("./uploads/"+handler.Filename, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Println("Error creating file:", err)
		http.Error(w, "Error creating file", http.StatusInternalServerError)
		return
	}

	defer f.Close()

	_, err = io.Copy(f, file)
	if err != nil {
		log.Println("Error copying file:", err)
		http.Error(w, "Error copying file", http.StatusInternalServerError)
		return
	}

	resp, err := checkFace("./uploads/" + handler.Filename)

	var student Student
	if err := json.NewDecoder(resp.Body).Decode(&student); err != nil {
		log.Fatalf("Failed to decode JSON response: %v", err)
	}
	log.Println(student)
	if resp.StatusCode == 404 {
		http.Error(w, "Student not found", http.StatusNotFound)
		return
	}

	if err := json.NewDecoder(resp.Body).Decode(&student); err != nil {
		log.Fatalf("Failed to decode JSON response: %v", err)
	}

	fmt.Fprintf(w, "File uploaded successfully: %s", handler.Filename)
	fmt.Fprintf(w, "Device MAC: %s, Device Key: %s, Room: %s", device.Mac, device.Key, device.Room)

}

func (a *App) protectedHandler(w http.ResponseWriter, r *http.Request) {

	session, _ := a.Store.Get(r, "aas-user")

	login := session.Values["user"]
	userID := session.Values["userID"]

	fmt.Fprintf(w, "Hello, %s! Your user ID is %d. This is a protected route.", login, userID)
}

func (a *App) getCoursesByTeacherID(w http.ResponseWriter, r *http.Request) {
	session, err := a.Store.Get(r, "aas-user")
	if err != nil {
		http.Error(w, "Failed to get session", http.StatusInternalServerError)
		log.Println("Error retrieving session:", err)
		return
	}

	teacherID, ok := session.Values["userID"].(int)
	if !ok {
		http.Error(w, "Failed to get user ID", http.StatusInternalServerError)
		log.Println("Error retrieving user ID from session")
		return
	}

	teacher := Teacher{Id: teacherID}

	err = teacher.getCourses(a.DB)
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
	log.Println(response)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode courses", http.StatusInternalServerError)
		log.Println("Error encoding courses to JSON:", err)
	}
}

func (a *App) getAllGroups(w http.ResponseWriter, r *http.Request) {
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

func (a *App) getAttendencesByClassID(w http.ResponseWriter, r *http.Request) {
	classID, err := strconv.Atoi(r.URL.Query().Get("classID"))
	if err != nil {
		http.Error(w, "Invalid classID", http.StatusBadRequest)
	}
	class := Class{Id: classID}

	err = class.getAttendences(a.DB)
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

func (a *App) getClassesByCourseID(w http.ResponseWriter, r *http.Request) {
	courseID, err := strconv.Atoi(r.URL.Query().Get("courseID"))
	if err != nil {
		http.Error(w, "Invalid courseID", http.StatusBadRequest)
	}
	classes, err := getClassesByCourseID(a.DB, courseID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(classes); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (a *App) getGroupsByCourseID(w http.ResponseWriter, r *http.Request) {
	courseID, err := strconv.Atoi(r.URL.Query().Get("courseID"))
	if err != nil {
		http.Error(w, "Invalid courseID", http.StatusBadRequest)
		return
	}

	groups, err := getGroupsByCourseID(a.DB, courseID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(groups); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (a *App) getAttendanceByCourseID(w http.ResponseWriter, r *http.Request) {
	courseID, err := strconv.Atoi(r.URL.Query().Get("courseID"))
	if err != nil {
		http.Error(w, "Invalid courseID", http.StatusBadRequest)
	}

	attendance, err := getAttendanceByCourse(a.DB, courseID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(attendance); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (a *App) updateAttendanceStatus(w http.ResponseWriter, r *http.Request) {
	var attendance = Attendance{}

	if err := json.NewDecoder(r.Body).Decode(&attendance); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		log.Println("Error decoding input:", err)
		return
	}

	err := attendance.update(a.DB)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("Error updating attendance status:", err)
		return
	}

	w.WriteHeader(200)
}

func (a *App) createCourse(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name      string   `json:"name"`
		Year      int      `json:"year"`
		StartDate string   `json:"start_date"`
		EndDate   string   `json:"end_date"`
		Groups    []string `json:"groups"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		log.Println("Error decoding input:", err)
		return
	}

	startDate, err := time.Parse("2006-01-02", input.StartDate)
	if err != nil {
		http.Error(w, "Invalid start date format", http.StatusBadRequest)
		log.Println("Error parsing start date:", err)
		return
	}
	endDate, err := time.Parse("2006-01-02", input.EndDate)
	if err != nil {
		http.Error(w, "Invalid end date format", http.StatusBadRequest)
		log.Println("Error parsing end date:", err)
		return
	}

	session, err := a.Store.Get(r, "aas-user")
	if err != nil {
		http.Error(w, "Failed to get session", http.StatusInternalServerError)
		log.Println("Error retrieving session:", err)
		return
	}

	teacherID, _ := session.Values["userID"].(int)
	err = createCourse(a.DB, input.Name, input.Year, startDate, endDate, teacherID, input.Groups)
	if err != nil {
		http.Error(w, "Failed to create course", http.StatusInternalServerError)
		log.Println("Error creating course:", err)
		return
	}
	w.WriteHeader(200)
}

func (a *App) registerTeacher(w http.ResponseWriter, r *http.Request) {
	var teacher = Teacher{}

	if err := json.NewDecoder(r.Body).Decode(&teacher); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		log.Println("Error decoding input:", err)
		return
	}

	emailPattern := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@nhlstenden\.com$`)
	if !emailPattern.MatchString(teacher.Email) {
		http.Error(w, "Invalid email format", http.StatusBadRequest)
		log.Println("Invalid email format:", teacher.Email)
		return
	}

	err := teacher.register(a.DB)
	if err != nil {
		http.Error(w, "Failed to register teacher", http.StatusInternalServerError)
		log.Println("Error registering teacher:", err)
		return
	}

	w.WriteHeader(200)
}

func checkFace(filename string) (*http.Response, error) {
	// Define the target URL
	targetURL := "http://172.20.10.13:5001/logic"

	// Open the specified file
	file, err := os.Open(filename)
	if err != nil {
		log.Println("Error opening file:", err)
		return nil, err
	}
	defer file.Close()

	// Create a buffer to store the multipart form data
	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)

	// Create a form file field for the image
	fileField, err := writer.CreateFormFile("image", filename)
	if err != nil {
		log.Println("Error creating form file field:", err)
		return nil, err
	}

	// Copy the image data to the form file field
	_, err = io.Copy(fileField, file)
	if err != nil {
		log.Println("Error copying image data:", err)
		return nil, err
	}

	// Close the multipart writer to finalize the form
	err = writer.Close()
	if err != nil {
		log.Println("Error closing multipart writer:", err)
		return nil, err
	}

	// Create a new HTTP request with the multipart form data
	req, err := http.NewRequest(http.MethodPost, targetURL, &requestBody)
	if err != nil {
		log.Println("Error creating HTTP request:", err)
		return nil, err
	}

	// Set the content type for the request
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Create an HTTP client and send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error sending request:", err)
		return nil, err
	}

	defer resp.Body.Close()

	// Read the response body
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response body:", err)
		return nil, err
	}

	// Print the response body for debugging purposes
	fmt.Println("Response body:", string(bodyBytes))

	// Decode the JSON response (optional)
	// var result map[string]interface{}
	// if err := json.Unmarshal(bodyBytes, &result); err != nil {
	// 	log.Println("Failed to decode JSON response:", err)
	// 	return nil, err
	// }

	// Log the decoded response for debugging purposes (optional)
	// log.Printf("Received response: %+v", result)

	return resp, nil
}
