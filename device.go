package main

import (
	"database/sql"
	"errors"
)

type Device struct {
	Mac  string `json:"id"`
	Key  string `json:"key"`
	Room string `json:"room"`
	//CurrentLesson Lesson `json:"current_lesson"`
}

func (device *Device) getDevice(db *sql.DB) error {
	return db.QueryRow("select Key, Room from users where Mac = ? LIMIT 1", device.Mac).Scan(&device.Key, &device.Room)
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
