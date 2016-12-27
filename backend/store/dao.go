package store

import "io"

type Dao interface {
	io.Closer
	Authenticate(login, password string) error
	GetProfile(login string) (Profile, error)
}
