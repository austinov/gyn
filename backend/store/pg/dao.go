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

	dictionaries = `
      SELECT id, name, 'pelvis_states' AS dict FROM pelvis_states
      UNION ALL
      SELECT id, name, 'fetal_bladder_aligns' AS dict FROM fetal_bladder_aligns
      UNION ALL
      SELECT id, name, 'fetal_bladder_previas' AS dict FROM fetal_bladder_previas
      UNION ALL
      SELECT id, name, 'outer_throat_states' AS dict FROM outer_throat_states
      UNION ALL
      SELECT id, name, 'vagina_states' AS dict FROM vagina_states
      UNION ALL
      SELECT id, name, 'devel_organs' AS dict FROM devel_organs
      UNION ALL
      SELECT id, name, 'reproductive_discharges' AS dict FROM reproductive_discharges
      UNION ALL
      SELECT id, name, 'fetal_aligns' AS dict FROM fetal_aligns
      UNION ALL
      SELECT id, name, 'fetal_heartbeats' AS dict FROM fetal_heartbeats
      UNION ALL
      SELECT id, name, 'fetal_previas' AS dict FROM fetal_previas
      UNION ALL
      SELECT id, name, 'fetal_positions' AS dict FROM fetal_positions
      UNION ALL
      SELECT id, name, 'uteruse_states' AS dict FROM uteruse_states
      UNION ALL
      SELECT id, name, 'skin_states' AS dict FROM skin_states
      UNION ALL
      SELECT id, name, 'health_states' AS dict FROM health_states
      ORDER BY dict, id;`
)

var (
	UserNotFoundError    = errors.New("user not found")
	userPasswordHashStmt *sql.Stmt
	userProfileStmt      *sql.Stmt
	dictionariesStmt     *sql.Stmt
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
	dictionariesStmt, err = db.Prepare(dictionaries)
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
	dictionariesStmt.Close()
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

func (d *dao) GetDictionaries() (map[string][]store.Dictionary, error) {
	rows, err := dictionariesStmt.Query()
	if err != nil {
		return nil, err
	}
	result := make(map[string][]store.Dictionary)
	for rows.Next() {
		var (
			id         int
			name, dict string
		)
		if err := rows.Scan(&id, &name, &dict); err != nil {
			return nil, err
		}
		dicts, ok := result[dict]
		if !ok {
			dicts = make([]store.Dictionary, 0)
		}
		dicts = append(dicts, store.Dictionary{
			Id:   id,
			Name: name,
		})
		result[dict] = dicts
	}
	return result, rows.Err()
}
