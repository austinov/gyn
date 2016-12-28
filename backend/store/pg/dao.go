package pg

import (
	"database/sql"
	"errors"
	"log"

	"github.com/austinov/gyn/backend/config"
	"github.com/austinov/gyn/backend/store"
	"github.com/austinov/gyn/backend/util"
	_ "github.com/lib/pq"
)

const (
	userPasswordHash = `
	  SELECT psw_hash FROM users WHERE login = $1`

	userProfile = `
	  SELECT name FROM users WHERE login = $1`
)

var (
	UserNotFoundError    = errors.New("user not found")
	userPasswordHashStmt *sql.Stmt
	userProfileStmt      *sql.Stmt
)

type dao struct {
	db *sql.DB
}

func New(cfg config.DBConfig) store.Dao {
	db, err := sql.Open("postgres", cfg.ConnectionString)
	if err != nil {
		log.Fatal(err)
	}
	userPasswordHashStmt, err = db.Prepare(userPasswordHash)
	if err != nil {
		log.Fatal(err)
	}
	userProfileStmt, err = db.Prepare(userProfile)
	if err != nil {
		log.Fatal(err)
	}
	return &dao{
		db: db,
	}
}

func (d *dao) Close() error {
	userPasswordHashStmt.Close()
	userProfileStmt.Close()
	d.db.Close()
	return nil
}

func (d *dao) Authenticate(login, password string) error {
	var hash string
	if err := userPasswordHashStmt.QueryRow(login).Scan(&hash); err != nil {
		return err
	}
	if err := util.CompareHashAndText(hash, password); err == nil {
		return nil
	}
	return UserNotFoundError
}

func (d *dao) GetProfile(login string) (store.Profile, error) {
	var name string
	err := userProfileStmt.QueryRow(login).Scan(&name)
	return store.Profile{
		UserName: name,
	}, err
}
