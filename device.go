//device.go

package main

import (
	"database/sql"
	"fmt"
)

type Device struct {
	Mac  string `json:"mac"`
	Key  string `json:"key"`
	Room string `json:"room"`
	//CurrentLesson Lesson `json:"current_lesson"`
}

func (device *Device) getDevice(db *sql.DB) error {
	qwery := fmt.Sprintf("SELECT `Key`, `Room` FROM `devices` where MAC LIKE '%v' LIMIT 1", device.Mac)
	return db.QueryRow(qwery).Scan(&device.Key, &device.Room)
}

func (device *Device) getClass(db *sql.DB) (Class, error) {
	qwery := fmt.Sprintf("SELECT * FROM classes WHERE room = '%v' AND NOW() BETWEEN startTime AND endTime LIMIT 1", device.Room)
	class := Class{}
	err := db.QueryRow(qwery).Scan(&class)
	return class, err
}
