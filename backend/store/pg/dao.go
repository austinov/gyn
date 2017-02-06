package pg

import (
	"database/sql"
	"log"
	"strings"

	"github.com/austinov/gyn/backend/config"
	"github.com/austinov/gyn/backend/store"
	"github.com/austinov/gyn/backend/util"
	_ "github.com/lib/pq"
)

const (
	selectUserPasswordHashSql = `
	  SELECT psw_hash FROM users WHERE login = $1`

	selectUserProfileSql = `
	  SELECT id, name FROM users WHERE login = $1`

	selectDictionariesSql = `
	  SELECT id, name, dict FROM (
        SELECT id, name, 'pelvis_states' AS dict, orderby FROM pelvis_states
        UNION ALL
        SELECT id, name, 'fetal_bladder_aligns' AS dict, orderby FROM fetal_bladder_aligns
        UNION ALL
        SELECT id, name, 'fetal_bladder_previas' AS dict, orderby FROM fetal_bladder_previas
        UNION ALL
        SELECT id, name, 'vagina_states' AS dict, orderby FROM vagina_states
        UNION ALL
        SELECT id, name, 'devel_organs' AS dict, orderby FROM devel_organs
        UNION ALL
        SELECT id, name, 'fetal_aligns' AS dict, orderby FROM fetal_aligns
        UNION ALL
        SELECT id, name, 'fetal_heartbeats' AS dict, orderby FROM fetal_heartbeats
        UNION ALL
        SELECT id, name, 'fetal_previas' AS dict, orderby FROM fetal_previas
        UNION ALL
        SELECT id, name, 'fetal_positions' AS dict, orderby FROM fetal_positions
        UNION ALL
        SELECT id, name, 'uteruse_states' AS dict, orderby FROM uteruse_states
        UNION ALL
        SELECT id, name, 'skin_states' AS dict, orderby FROM skin_states
        UNION ALL
        SELECT id, name, 'health_states' AS dict, orderby FROM health_states
        UNION ALL
		SELECT id, name, 'receipt_kinds' AS dict, orderby FROM receipt_kinds
        UNION ALL
		SELECT id, name, 'fetal_bladder_states' AS dict, orderby FROM fetal_bladder_states
        UNION ALL
		SELECT id, name, 'breath_states' AS dict, orderby FROM breath_states
        UNION ALL
		SELECT id, name, 'rale_states' AS dict, orderby FROM rale_states
        UNION ALL
		SELECT id, name, 'tones_states' AS dict, orderby FROM tones_states
        UNION ALL
		SELECT id, name, 'belly_states' AS dict, orderby FROM belly_states
        UNION ALL
		SELECT id, name, 'heartbeat_rithms' AS dict, orderby FROM heartbeat_rithms
        UNION ALL
        SELECT id, name, 'discharge_types' AS dict, orderby FROM discharge_types
        UNION ALL
		SELECT id, name, 'discharge_states' AS dict, orderby FROM discharge_states
        UNION ALL
		SELECT id, name, 'examination_states' AS dict, orderby FROM examination_states 
        UNION ALL
		SELECT id, name, 'oprv_states' AS dict, orderby FROM oprv_states
        UNION ALL
		SELECT id, name, 'external_throat_states' AS dict, orderby FROM external_throat_states
      ) t
      ORDER BY t.dict,  t.orderby, t.id`

	insertPatientSql = `
      WITH s AS (
          SELECT id
          FROM patients
          WHERE lower(name) = $1
      ), i as (
          INSERT INTO patients (name)
          SELECT $2
          WHERE NOT EXISTS (SELECT 1 FROM s)
          RETURNING id
      )
      SELECT id FROM i
      UNION ALL
      SELECT id FROM s`

	searchAppointmentsSql = `
	  SELECT a.id, u.name AS doctor_name, p.name AS patient_name, a.date_receipt
	  FROM appointments a
        JOIN users u ON a.doctor_id = u.id
	    JOIN patients p ON a.patient_id = p.id
	  WHERE p.name LIKE $1
	  ORDER BY a.date_receipt DESC, p.name
	  LIMIT 100`
)

var (
	selectUserPasswordHashStmt *sql.Stmt
	selectUserProfileStmt      *sql.Stmt
	selectDictionariesStmt     *sql.Stmt
	insertPatientStmt          *sql.Stmt
	searchAppointmentsStmt     *sql.Stmt
)

type dao struct {
	db *sql.DB
}

func New(cfg config.DBConfig) store.Dao {
	db, err := sql.Open("postgres", cfg.ConnectionString)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf("failed to ping database: %s", err)
	}
	selectUserPasswordHashStmt, err = db.Prepare(selectUserPasswordHashSql)
	if err != nil {
		log.Fatal(err)
	}
	selectUserProfileStmt, err = db.Prepare(selectUserProfileSql)
	if err != nil {
		log.Fatal(err)
	}
	selectDictionariesStmt, err = db.Prepare(selectDictionariesSql)
	if err != nil {
		log.Fatal(err)
	}
	selectAppointmentViewByIdStmt, err = db.Prepare(selectAppointmentViewByIdSql)
	if err != nil {
		log.Fatal(err)
	}
	insertPatientStmt, err = db.Prepare(insertPatientSql)
	if err != nil {
		log.Fatal(err)
	}
	insertAppointmentStmt, err = db.Prepare(insertAppointmentSql)
	if err != nil {
		log.Fatal(err)
	}
	updateAppointmentStmt, err = db.Prepare(updateAppointmentSql)
	if err != nil {
		log.Fatal(err)
	}
	searchAppointmentsStmt, err = db.Prepare(searchAppointmentsSql)
	if err != nil {
		log.Fatal(err)
	}
	return &dao{
		db: db,
	}
}

func (d *dao) Close() error {
	selectUserPasswordHashStmt.Close()
	selectUserProfileStmt.Close()
	selectDictionariesStmt.Close()
	insertPatientStmt.Close()
	searchAppointmentsStmt.Close()
	selectAppointmentViewByIdStmt.Close()
	insertAppointmentStmt.Close()
	updateAppointmentStmt.Close()
	d.db.Close()
	return nil
}

func (d *dao) Authenticate(login, password string) error {
	var hash string
	login = strings.ToLower(login)
	if err := selectUserPasswordHashStmt.QueryRow(login).Scan(&hash); err != nil {
		return err
	}
	if err := util.CompareHashAndText(hash, password); err == nil {
		return nil
	}
	return store.ErrUserNotFound
}

func (d *dao) GetProfile(login string) (store.Profile, error) {
	var (
		id   int32
		name string
	)
	err := selectUserProfileStmt.QueryRow(login).Scan(&id, &name)
	return store.Profile{
		Id:       id,
		UserName: name,
	}, err
}

func (d *dao) GetDictionaries() (map[string][]store.Dictionary, error) {
	rows, err := selectDictionariesStmt.Query()
	if err != nil {
		return nil, err
	}
	result := make(map[string][]store.Dictionary)
	for rows.Next() {
		var (
			id         int32
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

func (d *dao) SearchAppointments(patientName string) ([]store.AppointmentView, error) {
	rows, err := searchAppointmentsStmt.Query("%" + patientName + "%")
	if err != nil {
		return nil, err
	}
	result := make([]store.AppointmentView, 0)
	for rows.Next() {
		var (
			id, dateReceipt         int64
			patientName, doctorName string
		)
		if err := rows.Scan(&id, &doctorName, &patientName, &dateReceipt); err != nil {
			if err == sql.ErrNoRows {
				return result, nil
			}
			return nil, err
		}
		result = append(result, store.AppointmentView{
			Appointment: store.Appointment{
				Id:          id,
				DateReceipt: dateReceipt,
			},
			DoctorName:  doctorName,
			PatientName: patientName,
		})
	}
	return result, rows.Err()
}

func (d *dao) GetAppointment(id int64) (store.AppointmentView, error) {
	ap, err := getAppointmentViewById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			err = store.ErrDataNotFound
		}
	}
	return ap, err
}

func (d *dao) SaveAppointment(apv store.AppointmentView) (id int64, err error) {
	tx, err := d.db.Begin()
	if err != nil {
		return id, err
	}
	if err := tx.Stmt(insertPatientStmt).QueryRow(strings.ToLower(apv.PatientName), apv.PatientName).Scan(&apv.PatientId); err != nil {
		tx.Rollback()
		return id, err
	}
	ap := apv.Appointment
	if ap.Id == 0 {
		err = insertAppointment(tx, &ap)
	} else {
		cnt, nerr := updateAppointment(tx, &ap)
		if cnt != 1 {
			return ap.Id, store.ErrDataNotUpdated
		}
		err = nerr
	}
	if err != nil {
		tx.Rollback()
		return ap.Id, err
	}
	return ap.Id, tx.Commit()
}
