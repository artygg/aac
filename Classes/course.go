package Classes

import (
	"database/sql"
	"errors"
	"fmt"
)

type Course struct {
	ID   string `json:"mac"`
	Key  string `json:"key"`
	Room string `json:"room"`
	//CurrentLesson Lesson `json:"current_lesson"`
}

func (course *Course) getCourse(db *sql.DB) error {
	qwery := fmt.Sprintf("SELECT `Key`, `Room` FROM `device` where ID LIKE '%v' LIMIT 1", course.ID)
	return db.QueryRow(qwery).Scan(&course.Key, &course.Room)
}

func (course *Course) getCourses(db *sql.DB, id int) error {
	qwery := fmt.Sprintf("SELECT `Key`, `Room` FROM `device` where ID LIKE '%v' LIMIT 1", course.ID)
	return db.QueryRow(qwery).Scan(&course.Key, &course.Room)
}
func (device *Device) updateDevice(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (device *Device) deleteDevice(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (device *Device) createDevice(db *sql.DB) error {
	return errors.New("Not implemented")
}
