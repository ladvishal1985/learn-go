package models

import (
	"fmt"
	"time"

	"events.com/api/db"
)

type Event struct {
	ID          int64
	Name        string
	Description string
	Location    string
	DataTime    time.Time
	UserId      int
}

var events = []Event{}

func (e Event) Save() error {
	query := `INSERT INTO events(name, description,location, dateTime, user_Id)
	VALUES (?, ?, ?, ?, ?)
	`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		fmt.Print(err)
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DataTime, e.UserId)
	if err != nil {
		fmt.Print(err)
		return err
	}

	id, err := result.LastInsertId()
	e.ID = id
	return err
}

func GetAllEvents() ([]Event, error) {
	query := `
		SELECT * FROM events
	`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var events []Event
	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DataTime, &event.UserId)
		if err != nil {
			return nil, err
		}
	}
	return events, nil
}

func GetEventById(id int64) (*Event, error) {
	query := `
		SELECT * FROM events WHERE id = ?
	`
	row := db.DB.QueryRow(query, id)
	// if err != nil {
	// 	return nil, err;
	// }
	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DataTime, &event.UserId)
	if err != nil {
		return nil, err
	}
	return &event, nil
}

func (event Event) UpdateEvent() error {
	query := `
		UPDATE  events SET 
		SET name = ?, description = ?, location = ?, dateTime = ?
		WHERE id = ?
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(event.Name, event.Description, event.Location, event.DataTime, event.ID)
	return err
}
