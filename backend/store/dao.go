package store

import (
	"errors"
	"io"
)

var (
	ErrUserNotFound   = errors.New("user not found")
	ErrDataNotFound   = errors.New("data not found")
	ErrDataNotUpdated = errors.New("data not updated")
)

type Dao interface {
	io.Closer
	Authenticate(login, password string) error
	GetProfile(login string) (Profile, error)
	GetDictionaries() (map[string][]Dictionary, error)
	SearchAppointments(patientName string) ([]AppointmentView, error)
	GetAppointment(id int64) (AppointmentView, error)
	SaveAppointment(ap *AppointmentView) error
}
