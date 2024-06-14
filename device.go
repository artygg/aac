//device.go

package main

import (
	"database/sql"
	"fmt"
	"log"
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
	qwery := fmt.Sprintf("SELECT `Id` FROM classes WHERE room = '%v' AND NOW() BETWEEN startTime AND endTime LIMIT 1", device.Room)
	class := Class{}
	err := db.QueryRow(qwery).Scan(&class.Id)
	return class, err
}

func getRooms(db *sql.DB) ([]Device, error) {
	query := fmt.Sprintf("SELECT `Room` FROM `devices`")
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

	var devices []Device

	for rows.Next() {
		var device Device
		if err := rows.Scan(&device.Room); err != nil {
			log.Println("Error scanning row:", err)
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		devices = append(devices, device)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error iterating over rows:", err)
		return nil, fmt.Errorf("error iterating over rows: %w", err)
	}

	return devices, nil
}
