package main

import (
	"database/sql"
	"errors"
	"fmt"
)

type Device struct {
	Mac  string `json:"mac"`
	Key  string `json:"key"`
	Room string `json:"room"`
	//CurrentLesson Lesson `json:"current_lesson"`
}

func (device *Device) getDevice(db *sql.DB) error {
	qwery := fmt.Sprintf("SELECT `Key`, `Room` FROM `device`  where MAC LIKE '%v' LIMIT 1", device.Mac)
	return db.QueryRow(qwery).Scan(&device.Key, &device.Room)
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
