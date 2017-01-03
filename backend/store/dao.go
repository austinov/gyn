package store

import (
	"errors"
	"io"
)

var (
	ErrUserNotFound = errors.New("user not found")
	ErrDataNotFound = errors.New("data not found")
)

type Dao interface {
	io.Closer
	Authenticate(login, password string) error
	GetProfile(login string) (Profile, error)
	GetDictionaries() (map[string][]Dictionary, error)
	GetAppointment(id int64) (Appointment, error)
}
