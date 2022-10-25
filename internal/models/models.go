package models

import (
	"time"
)

// user model for the users table
type User struct {
	ID          int
	FirstName   string
	LastName    string
	Email       string
	Password    string
	AccessLevel int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// room model
type Room struct {
	ID        int
	RoomName  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// restriction model
type Restriction struct {
	ID              int
	RestrictionName string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

// model for reservations table
type Reservation struct {
	ID        int
	CreatedAt time.Time
	UpdatedAt time.Time
	FirstName string
	LastName  string
	Email     string
	Phone     string
	StartDate time.Time
	EndDate   time.Time
	RoomID    int
	Room      Room
}

// model for room restrictions table
type RoomRestriction struct {
	ID            int
	CreatedAt     time.Time
	UpdatedAt     time.Time
	StartDate     time.Time
	EndDate       time.Time
	RoomID        int
	Room          Room
	ReservationID int
	Reservation   Reservation
	RestrictionID int
	Restriction   Restriction
}

// type for an email message
type MailData struct {
	To      string
	From    string
	Subject string
	Content string
}
