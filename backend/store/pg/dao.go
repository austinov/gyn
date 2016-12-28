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
	userExist = `
	  SELECT count(id) FROM users WHERE login = $1 AND psw_hash = $2`
)

var (
	UserNotFoundError = errors.New("user not found")
	userExistStmt     *sql.Stmt
)

type dao struct {
	db *sql.DB
}

func New(cfg config.DBConfig) store.Dao {
	db, err := sql.Open("postgres", cfg.ConnectionString)
	if err != nil {
		log.Fatal(err)
	}
	userExistStmt, err = db.Prepare(userExist)
	if err != nil {
		log.Fatal(err)
	}
	return &dao{
		db: db,
	}
}

func (d *dao) Close() error {
	userExistStmt.Close()
	d.db.Close()
	return nil
}

func (d *dao) Authenticate(login, password string) error {
	// TODO
	var cnt int
	if err := userExistStmt.QueryRow(login, util.Hash(password)).Scan(&cnt); err != nil {
		return err
	}
	if cnt == 1 {
		return nil
	}
	//if login == "q@aa.zz" && password == "123" {
	//	return nil
	//}
	return UserNotFoundError
}

func (d *dao) GetProfile(login string) (store.Profile, error) {
	// TODO
	return store.Profile{
		UserName: "Алексей Лукаев",
	}, nil
}
