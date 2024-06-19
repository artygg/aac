//class.go

package main

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"
)

type Class struct {
	Id          int          `json:"id"`
	CourseId    int          `json:"course_id"`
	Room        string       `json:"room"`
	StartTime   string       `json:"starttime"`
	EndTime     string       `json:"endtime"`
	Attendances []Attendance `json:"attendances"`
}

func (class *Class) getAttendences(db *sql.DB) error {
	qwery := fmt.Sprintf("SELECT a.Status, a.Time, s.Id, s.FirstName, s.LastName, s.Email FROM attendances a INNER JOIN classes c ON a.ClassId = c.Id INNER JOIN students s ON a.StudentId = s.Id WHERE a.ClassID = '%v' ", class.Id)

	rows, err := db.Query(qwery)
	if err != nil {
		return err
	}
	for rows.Next() {
		var attendance Attendance
		err := rows.Scan(&attendance.Status, &attendance.Time, &attendance.Student.Id, &attendance.Student.FirstName, &attendance.Student.LastName, &attendance.Student.Email)
		if err != nil {
			return err
		}
		class.Attendances = append(class.Attendances, attendance)
	}
	return nil
}

func getClassesByCourseID(db *sql.DB, courseID int) ([]Class, error) {
	query := fmt.Sprintf("SELECT Id, CourseID, Room, StartTime, EndTime FROM classes WHERE CourseID = %v", courseID)
	rows, err := db.Query(query)
	if err != nil {
		log.Println("Error executing query:", err)
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer func() {
		if err := rows.Close(); err != nil {
			log.Println("Error closing rows:", err)
		}
	}()

	var classes []Class

	for rows.Next() {
		var class Class
		if err := rows.Scan(&class.Id, &class.CourseId, &class.Room, &class.StartTime, &class.EndTime); err != nil {
			log.Println("Error scanning row:", err)
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		classes = append(classes, class)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error iterating over rows:", err)
		return nil, fmt.Errorf("error iterating over rows: %w", err)
	}

	return classes, nil
}

func createClass(db *sql.DB, courseID int, startTime time.Time, endTime time.Time, room string, groups []string) error {
	tx, err := db.Begin()
	if err != nil {
		log.Println("Error starting transaction:", err)
		return err
	}

	query := `INSERT INTO classes (CourseID, StartTime, EndTime, Room) VALUES (?, ?, ?, ?)`
	res, err := tx.Exec(query, courseID, startTime, endTime, room)
	if err != nil {
		tx.Rollback()
		log.Println("Error inserting class:", err)
		return err
	}

	classID, err := res.LastInsertId()
	if err != nil {
		tx.Rollback()
		log.Println("Error getting last insert ID:", err)
		return err
	}

	for _, groupID := range groups {
		log.Println("*Group list: <", groups, ">")
		query := "INSERT INTO `classes-groups-bridge` (Classid, GroupID) VALUES (?, ?)"
		if _, err := tx.Exec(query, classID, groupID); err != nil {
			tx.Rollback()
			log.Println("Error inserting into classes_groups_bridge:", err)
			return err
		}
	}

	for i := range groups {
		log.Print("Group : <", groups[i], ">")
	}
	// Retrieve students belonging to the groups associated with the class
	students, err := getStudentsInGroups(db, groups)
	log.Println("Students: ", students)
	if err != nil {
		tx.Rollback()
		log.Println("Error retrieving students:", err)
		return err
	}

	// Insert attendance records for each student and the newly created class
	for _, student := range students {
		query := "INSERT INTO attendances (ClassID, StudentID, Time, Status) VALUES (?, ?, ?, ?)"
		_, err := tx.Exec(query, classID, student.Id, time.Now(), 0)
		if err != nil {
			tx.Rollback()
			log.Println("Error inserting attendance record:", err)
			return err
		}
	}

	if err := tx.Commit(); err != nil {
		log.Println("Error committing transaction:", err)
		return err
	}

	return nil
}

func getStudentsInGroups(db *sql.DB, groups []string) ([]Student, error) {
	var students []Student
	log.Println("Before student appending")
	// Construct a query to retrieve students belonging to the specified groups
	query := fmt.Sprintf("SELECT s.ID, s.FirstName, s.LastName, s.Email FROM students s INNER JOIN `groups assign` ga ON s.ID = ga.StudentID INNER JOIN `groups` g ON ga.GroupID = g.ID WHERE g.ID IN ('%s')", strings.Join(groups, "', '"))

	rows, err := db.Query(query)
	if err != nil {
		log.Println("Error executing getStudentsInGroups() query:", err)
		return nil, err
	}
	defer rows.Close()
	log.Println("Row: ", rows)
	for rows.Next() {
		var student Student
		log.Println("Student with id ", student.Id, " is being added")
		if err := rows.Scan(&student.Id, &student.FirstName, &student.LastName, &student.Email); err != nil {
			log.Println("Error scanning row:", err)
			return nil, err
		}

		students = append(students, student)
	}
	log.Println("After student appending")
	if err := rows.Err(); err != nil {
		log.Println("Error iterating over rows:", err)
		return nil, err
	}

	return students, nil
}

func endClassPrematurely(db *sql.DB, classID int) error {
	query := `UPDATE classes SET EndTime = ? WHERE Id = ?`
	_, err := db.Exec(query, time.Now(), classID)
	if err != nil {
		log.Println("Error updating class end time:", err)
		return err
	}
	return nil
}
