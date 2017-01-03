package pg

import (
	"database/sql"
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
	  SELECT id, name, dict FROM (
        SELECT id, name, 'pelvis_states' AS dict, orderby FROM pelvis_states
        UNION ALL
        SELECT id, name, 'fetal_bladder_aligns' AS dict, orderby FROM fetal_bladder_aligns
        UNION ALL
        SELECT id, name, 'fetal_bladder_previas' AS dict, orderby FROM fetal_bladder_previas
        UNION ALL
        SELECT id, name, 'outer_throat_states' AS dict, orderby FROM outer_throat_states
        UNION ALL
        SELECT id, name, 'vagina_states' AS dict, orderby FROM vagina_states
        UNION ALL
        SELECT id, name, 'devel_organs' AS dict, orderby FROM devel_organs
        UNION ALL
        SELECT id, name, 'reproductive_discharges' AS dict, orderby FROM reproductive_discharges
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
      ) t
      ORDER BY t.dict,  t.orderby, t.id`

	patientAppointment = `
	  SELECT a.*, u.name AS doctor_name, p.name AS patient_name
	  FROM appointments a
        JOIN users u ON a.doctor_id = u.id
		JOIN patients p ON a.patient_id = p.id
	  WHERE a.id = $1`
)

var (
	userPasswordHashStmt   *sql.Stmt
	userProfileStmt        *sql.Stmt
	dictionariesStmt       *sql.Stmt
	patientAppointmentStmt *sql.Stmt
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
	patientAppointmentStmt, err = db.Prepare(patientAppointment)
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
	patientAppointmentStmt.Close()
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
	return store.ErrUserNotFound
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

func (d *dao) GetAppointment(id int64) (store.Appointment, error) {
	ap := store.Appointment{}
	err := patientAppointmentStmt.QueryRow(id).Scan(
		&ap.Id,
		&ap.Date,
		&ap.DoctorId,
		&ap.PatientId,
		&ap.HowReceipt,
		&ap.Alergo,
		&ap.ContactInfectied,
		&ap.Hiv,
		&ap.Transfusion,
		&ap.Dyscountry,
		&ap.Smoking,
		&ap.Drugs,
		&ap.Inheritance,
		&ap.Diseases,
		&ap.Gyndiseases,
		&ap.Paritet,
		&ap.Pregnancy,
		&ap.FirstTrimester,
		&ap.SecondTrimester,
		&ap.ThirdTrimester,
		&ap.History,
		&ap.ExpByMenstruation,
		&ap.ExpByFirstVisit,
		&ap.ExpByUltra,
		&ap.HealthStateId,
		&ap.Claims,
		&ap.Head,
		&ap.Vision,
		&ap.SkinStateId,
		&ap.Lymph,
		&ap.Breath,
		&ap.Rale,
		&ap.Tones,
		&ap.Pulse,
		&ap.PulseType,
		&ap.Pressure,
		&ap.TongueClean,
		&ap.TongueWet,
		&ap.TongueDry,
		&ap.TongueCoated,
		&ap.TongueUncoated,
		&ap.Throat,
		&ap.Belly,
		&ap.Peritoneal,
		&ap.Labors,
		&ap.Dysuric,
		&ap.Bowel,
		&ap.LimbSwelling,
		&ap.FaceSwelling,
		&ap.UteruseStateId,
		&ap.FetalPositionId,
		&ap.FetalPreviaId,
		&ap.FetalAlignId,
		&ap.FetalHeartbeatId,
		&ap.FetalPulse,
		&ap.ReproductiveDischargeId,
		&ap.Vdm,
		&ap.Oj,
		&ap.Dspin,
		&ap.Dcrist,
		&ap.Dtroch,
		&ap.Cext,
		&ap.DevelOrgansId,
		&ap.GenitalAnomalies,
		&ap.VaginaStateId,
		&ap.LenghtCervix,
		&ap.TruncateCervix,
		&ap.OuterThroatStateId,
		&ap.ChannelCervix,
		&ap.FetalBladder,
		&ap.FetalBladderPreviaId,
		&ap.FetalBladderAlignId,
		&ap.Arches,
		&ap.Conjugate,
		&ap.PelvisStateId,
		&ap.PelvisDischarge,
		&ap.Diagnosis,
		&ap.Conclusion,
		&ap.BirthPlan,
		&ap.DoctorName,
		&ap.PatientName)
	if err != nil && err == sql.ErrNoRows {
		err = store.ErrDataNotFound
	}
	return ap, err
}
