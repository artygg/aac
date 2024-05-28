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
	"time"
	"regexp"

	"github.com/go-sql-driver/mysql"
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

	a.Router.HandleFunc("/api/web/classes/{courseID}", a.userAuthMiddleware(http.HandlerFunc(a.getClassesByCourseID))).Methods("GET")
	a.Router.HandleFunc("/api/web/groups/{courseID}", a.userAuthMiddleware(http.HandlerFunc(a.getGroupsByCourseID))).Methods("GET")
	a.Router.HandleFunc("/api/web/groups/{courseID}", a.userAuthMiddleware(http.HandlerFunc(a.getAttendanceByCourseID))).Methods("GET")
	a.Router.HandleFunc("/api/web/attendance/update", a.userAuthMiddleware(http.HandlerFunc(a.updateAttendanceStatus))).Methods("POST")

	a.Router.HandleFunc("/api/web/course/create", a.userAuthMiddleware(http.HandlerFunc(a.createCourse))).Methods("POST")
	a.Router.HandleFunc("/api/web/teacher/register", http.HandlerFunc(a.registerTeacher)).Methods("POST")
	a.Router.HandleFunc("/api/web/teacher/register", a.userAuthMiddleware(http.HandlerFunc(a.createCourse))).Methods("POST")


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

func (a *App) getCoursesByTeacherID(w http.ResponseWriter, r *http.Request) {
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
	courseID := r.URL.Query().Get("courseID")

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
	courseID := r.URL.Query().Get("courseID")

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
	courseID := r.URL.Query().Get("courseID")

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
    var input struct {
        StudentID int `json:"student_id"`
        Status    int `json:"status"`
    }

    if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        log.Println("Error decoding input:", err)
        return
    }

    if input.Status < 0 || input.Status > 3 {
        http.Error(w, "Invalid status value", http.StatusBadRequest)
        log.Println("Invalid status value:", input.Status)
        return
    }

    success, err := updateAttendance(a.DB, input.StudentID, input.Status)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        log.Println("Error updating attendance status:", err)
        return
    }

    if !success {
        http.Error(w, "No attendance record updated", http.StatusNotFound)
        log.Println("No attendance record updated for StudentID:", input.StudentID)
        return
    }

    response := map[string]string{"message": "Attendance status updated successfully"}
    w.Header().Set("Content-Type", "application/json")
    if err := json.NewEncoder(w).Encode(response); err != nil {
        http.Error(w, "Failed to encode response", http.StatusInternalServerError)
        log.Println("Error encoding response:", err)
    }
}

func (a *App) createCourse(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name      string `json:"name"`
		Year      int    `json:"year"`
		StartDate string `json:"start_date"`
		EndDate   string `json:"end_date"`
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

	teacherID := getTeacherIDFromSession(r) 

	success, err := createCourse(a.DB, input.Name, input.Year, startDate, endDate, teacherID, input.Groups)
	if err != nil {
		http.Error(w, "Failed to create course", http.StatusInternalServerError)
		log.Println("Error creating course:", err)
		return
	}

	response := map[string]string{
		"message": "Course created successfully",
		"success": fmt.Sprintf("%t", success),
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		log.Println("Error encoding response:", err)
	}
}


func (a *App) registerTeacher(w http.ResponseWriter, r *http.Request) {
    var input struct {
        Email           string `json:"email"`
        FirstName       string `json:"first_name"`
        LastName        string `json:"last_name"`
        Password        string `json:"password"`
        RegistrationDate string `json:"registration_date"`
    }

    if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        log.Println("Error decoding input:", err)
        return
    }

    emailPattern := regexp.MustCompile(`^[^@]+@[^@]+\.[^@]{2,}$`)
    if !emailPattern.MatchString(input.Email) {
        http.Error(w, "Invalid email format", http.StatusBadRequest)
        log.Println("Invalid email format:", input.Email)
        return
    }

    if input.RegistrationDate == "" {
        input.RegistrationDate = time.Now().Format("2006-01-02")
    }

    success, err := registerTeacher(a.DB, input.Email, input.FirstName, input.LastName, input.Password, input.RegistrationDate)
    if err != nil {
        http.Error(w, "Failed to register teacher", http.StatusInternalServerError)
        log.Println("Error registering teacher:", err)
        return
    }

    response := map[string]string{
        "message": "Teacher registered successfully",
        "success": fmt.Sprintf("%t", success),
    }
    w.Header().Set("Content-Type", "application/json")
    if err := json.NewEncoder(w).Encode(response); err != nil {
        http.Error(w, "Failed to encode response", http.StatusInternalServerError)
        log.Println("Error encoding response:", err)
    }
}