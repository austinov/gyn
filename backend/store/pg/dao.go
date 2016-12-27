package pg

import (
	"database/sql"
	"errors"
	"log"

	"github.com/austinov/gyn/backend/config"
	"github.com/austinov/gyn/backend/store"
	_ "github.com/lib/pq"
)

var (
	UserNotFoundError = errors.New("user not found")
)

type dao struct {
	db *sql.DB
}

func New(cfg config.DBConfig) store.Dao {
	db, err := sql.Open("postgres", cfg.ConnectionString)
	if err != nil {
		log.Fatal(err)
	}
	return &dao{
		db: db,
	}
}

func (d *dao) Close() error {
	d.db.Close()
	return nil
}

func (d *dao) Authenticate(login, password string) error {
	// TODO
	if login == "q@aa.zz" && password == "123" {
		return nil
	}
	return UserNotFoundError
}

func (d *dao) GetProfile(login string) (store.Profile, error) {
	// TODO
	return store.Profile{
		UserName: "Алексей Лукаев",
	}, nil
}
