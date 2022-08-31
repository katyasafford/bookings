package repository

import (
	"time"

	"github.com/katyasafford/bookings/internal/models"
)

type DatabaseRepo interface {
	AllUsers() bool

	InsertReservation(res models.Reservation) (int, error)
	InsertRoomRestriction(r models.RoomRestriction) error
	SearchAvailabilityByDatesByRoomID(start, end time.Time, roomID int) (bool, error)
	SearchAvailalbilityForAllRooms(start, end time.Time) ([]models.Room, error)
}
