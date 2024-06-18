package main

import (
	"context"
	"crypto/rand"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
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
	Aws    string
}

func (a *App) Initialize() {
	var err error

	a.DB, err = sql.Open("mysql", "")
	if err != nil {
		log.Fatal(err)
	}
	a.DB.SetConnMaxLifetime(0)
	a.DB.SetMaxOpenConns(0)
	a.DB.SetMaxIdleConns(0)

	a.Router = mux.NewRouter()
	a.Aws = "192.168.100.12"
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
	a.Router.Handle("/api/device/authorize", a.deviceAuthMiddleware(http.HandlerFunc(a.checkDevice)))
	a.Router.Handle("/api/device/attendance", a.deviceAuthMiddleware(http.HandlerFunc(a.putAttendance))).Methods("POST")

	a.Router.HandleFunc("/api/web/login", a.loginHandler).Methods("POST")

	a.Router.Handle("/api/web/classes", a.userAuthMiddleware(http.HandlerFunc(a.getClassesByCourseID))).Methods("GET")
	a.Router.Handle("/api/web/courses", a.userAuthMiddleware(http.HandlerFunc(a.getCoursesByTeacherID))).Methods("GET")
	a.Router.Handle("/api/web/groups", a.userAuthMiddleware(http.HandlerFunc(a.getAllGroups))).Methods("GET")
	a.Router.Handle("/api/web/groups/by_course", a.userAuthMiddleware(http.HandlerFunc(a.getGroupsByCourseID))).Methods("GET")
	a.Router.Handle("/api/web/rooms", a.userAuthMiddleware(http.HandlerFunc(a.getRooms))).Methods("GET")

	a.Router.Handle("/api/web/attendance/by_class", a.userAuthMiddleware(http.HandlerFunc(a.getAttendencesByClassID))).Methods("GET")
	a.Router.Handle("/api/web/attendance/by_course", a.userAuthMiddleware(http.HandlerFunc(a.getAttendanceByCourseID))).Methods("GET")

	a.Router.Handle("/api/web/attendance", a.userAuthMiddleware(http.HandlerFunc(a.updateAttendanceStatus))).Methods("POST")
	a.Router.Handle("/api/web/class/end", a.userAuthMiddleware(http.HandlerFunc(a.endClassPrematurely))).Methods("POST")
	a.Router.Handle("/api/web/class", a.userAuthMiddleware(http.HandlerFunc(a.createClass))).Methods("POST")
	a.Router.Handle("/api/web/course", a.userAuthMiddleware(http.HandlerFunc(a.createCourse))).Methods("POST")
	a.Router.Handle("/api/web/teacher", http.HandlerFunc(a.registerTeacher)).Methods("POST")
	a.initializeClient()
}

func (a *App) initializeClient() {

	a.Router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./platform/static"))))

	a.Router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./platform/index.html")
	})

	a.Router.Handle("/courses", a.userAuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./platform/courses-page.html")
	})))

	a.Router.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./platform/signin.html")
	})

	a.Router.HandleFunc("/registration", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./platform/register.html")
	})

	a.Router.Handle("/protected", a.userAuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./platform/protected.html")
	})))

	a.Router.Handle("/classes", a.userAuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./platform/classes-page.html")
	})))

	a.Router.Handle("/course/create", a.userAuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./platform/start-new-course.html")
	})))

	a.Router.Handle("/class/create", a.userAuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./platform/start-new-class.html")
	})))

	a.Router.Handle("/attendance/by_course", a.userAuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./platform/course-attendance-statistics.html")
	})))

	a.Router.Handle("/attendance/by_class", a.userAuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./platform/class-attendance.html")
	})))

	a.Router.HandleFunc("/logout", a.logoutHandler)
}

func (a *App) deviceAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		device := Device{}
		device.Mac = r.Header.Get("X-MAC-ADDRESS")
		key := r.Header.Get("X-API-KEY")
		log.Println(r.RemoteAddr)
		log.Println(device, key)
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


	if !checkPasswordHash(password, teacher.Password) {
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
	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	if host == a.Aws {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}
}

func (a *App) putAttendance(w http.ResponseWriter, r *http.Request) {
	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	if host == a.Aws {
		device, ok := r.Context().Value("device").(Device)
		if !ok {
			log.Println("Device not dound!")
			http.Error(w, "Device information not found", http.StatusInternalServerError)
			return
		}
		class, err := device.getClass(a.DB)
		if err != nil {
			switch err {
			case sql.ErrNoRows:
				log.Println("No active class for the current device")
				http.Error(w, "No Active Class for this room", http.StatusServiceUnavailable)
			default:
				log.Println(err)
				http.Error(w, "Internal server error", http.StatusInternalServerError)
			}
			return
		}
		var student Student
		if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
			log.Fatalf("Failed to decode JSON response: %v", err)
		}
		var attendance = Attendance{ClassID: class.Id, Student: student, Status: "1"}
		err = attendance.update(a.DB)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Println("Error updating attendance status:", err)
			return
		}
		w.WriteHeader(200)
	}
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

func (a *App) createClass(w http.ResponseWriter, r *http.Request) {
	var input struct {
		CourseID  int      `json:"course_id"`
		StartTime string   `json:"start_time"`
		EndTime   string   `json:"end_time"`
		Room      string   `json:"room"`
		Groups    []string `json:"groups"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		log.Println("Error decoding input:", err)
		return
	}

	startTime, err := time.Parse("2006-01-02 15:04", input.StartTime)
	if err != nil {
		http.Error(w, "Invalid start time format", http.StatusBadRequest)
		log.Println("Error parsing start time:", err)
		return
	}
	endTime, err := time.Parse("2006-01-02 15:04", input.EndTime)
	if err != nil {
		http.Error(w, "Invalid end time format", http.StatusBadRequest)
		log.Println("Error parsing end time:", err)
		return
	}

	err = createClass(a.DB, input.CourseID, startTime, endTime, input.Room, input.Groups)
	if err != nil {
		http.Error(w, "Failed to create class", http.StatusInternalServerError)
		log.Println("Error creating class:", err)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (a *App) endClassPrematurely(w http.ResponseWriter, r *http.Request) {
	var input struct {
		ClassID int `json:"class_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		log.Println("Error decoding input:", err)
		return
	}

	err := endClassPrematurely(a.DB, input.ClassID)
	if err != nil {
		http.Error(w, "Failed to end class prematurely", http.StatusInternalServerError)
		log.Println("Error ending class prematurely:", err)
		return
	}

	w.WriteHeader(http.StatusOK)
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

func (a *App) getRooms(w http.ResponseWriter, r *http.Request) {
	groups, err := getRooms(a.DB)

	if err != nil {
		http.Error(w, "Failed to get courses", http.StatusInternalServerError)
		log.Println("Error retrieving courses:", err)
		return
	}

	log.Printf("Retrieved courses: %+v\n", groups)

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(groups); err != nil {
		http.Error(w, "Failed to encode courses", http.StatusInternalServerError)
		log.Println("Error encoding courses to JSON:", err)
	}
}
