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
	userPasswordHashSelect = `
	  SELECT psw_hash FROM users WHERE login = $1`

	userProfileSelect = `
	  SELECT id, name FROM users WHERE login = $1`

	dictionariesSelect = `
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

	patientAppointmentSelect = `
	  SELECT a.*, u.name AS doctor_name, p.name AS patient_name
	  FROM appointments a
        JOIN users u ON a.doctor_id = u.id
		JOIN patients p ON a.patient_id = p.id
	  WHERE a.id = $1`

	patientInsert = `
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

	appointmentsSelect = `
	  SELECT a.id, u.name AS doctor_name, p.name AS patient_name, a.date_receipt
	  FROM appointments a
        JOIN users u ON a.doctor_id = u.id
	    JOIN patients p ON a.patient_id = p.id
	  WHERE p.name LIKE $1
	  ORDER BY a.date_receipt DESC, p.name
	  LIMIT 100`

	appointmentInsert = `
      INSERT INTO appointments (
        date_receipt, doctor_id, patient_id, how_receipt, alergo, contact_infected,
        hiv, transfusion, dyscountry, smoking, drugs, inheritance, diseases, gyndiseases,
        paritet, pregnancy, first_trimester, second_trimester, third_trimester, history,
        exp_by_menstruation, exp_by_first_visit, exp_by_ultra, health_state_id, claims,
        head, vision, skin_state_id, lymph, breath, rale, tones, pulse, pulse_type,
        pressure, tongue_clean, tongue_wet, tongue_dry, tongue_coated, tongue_uncoated,
        throat, belly, peritoneal, labors, dysuric, bowel, limb_swelling, face_swelling,
        uteruse_state_id, fetal_position_id, fetal_previa_id, fetal_align_id, fetal_heartbeat_id,
        fetal_pulse, reproductive_discharge_id, vdm, oj, dspin, dcrist, dtroch, cext,
        devel_organs_id, genital_anomalies, vagina_state_id, lenght_cervix, truncate_cervix,
        outer_throat_state_id, channel_cervix, fetal_bladder, fetal_bladder_previa_id,
        fetal_bladder_align_id, arches, conjugate, pelvis_state_id, pelvis_discharge,
        diagnosis, conclusion, birth_plan)
      VALUES (
        $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19,
        $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35, $36,
        $37, $38, $39, $40, $41, $42, $43, $44, $45, $46, $47, $48, $49, $50, $51, $52, $53,
        $54, $55, $56, $57, $58, $59, $60, $61, $62, $63, $64, $65, $66, $67, $68, $69, $70,
        $71, $72, $73, $74, $75, $76, $77, $78)
  	  RETURNING id`

	appointmentUpdate = `
	  WITH rows AS (
        UPDATE appointments
          SET date_receipt = $1, doctor_id = $2, patient_id = $3, how_receipt = $4, alergo = $5,
              contact_infected = $6, hiv = $7, transfusion = $8, dyscountry = $9, smoking = $10,
              drugs = $11, inheritance = $12, diseases = $13, gyndiseases = $14, paritet = $15,
              pregnancy = $16, first_trimester = $17, second_trimester = $18, third_trimester = $19,
              history = $20, exp_by_menstruation = $21, exp_by_first_visit = $22, exp_by_ultra = $23,
              health_state_id = $24, claims = $25, head = $26, vision = $27, skin_state_id = $28,
              lymph = $29, breath = $30, rale = $31, tones = $32, pulse = $33, pulse_type = $34,
              pressure = $35, tongue_clean = $36, tongue_wet = $37, tongue_dry = $38,
              tongue_coated = $39, tongue_uncoated = $40, throat = $41, belly = $42, peritoneal = $43,
              labors = $44, dysuric = $45, bowel = $46, limb_swelling = $47, face_swelling = $48,
              uteruse_state_id = $49, fetal_position_id = $50, fetal_previa_id = $51,
              fetal_align_id = $52, fetal_heartbeat_id = $53, fetal_pulse = $54,
              reproductive_discharge_id = $55, vdm = $56, oj = $57, dspin = $58, dcrist = $59,
              dtroch = $60, cext = $61, devel_organs_id = $62, genital_anomalies = $63,
              vagina_state_id = $64, lenght_cervix = $65, truncate_cervix = $66,
              outer_throat_state_id = $67, channel_cervix = $68, fetal_bladder = $69,
              fetal_bladder_previa_id = $70, fetal_bladder_align_id = $71, arches = $72,
              conjugate = $73, pelvis_state_id = $74, pelvis_discharge = $75,
              diagnosis = $76, conclusion = $77, birth_plan = $78
        WHERE id = $79
		RETURNING 1
	  )
	  SELECT count(*) FROM rows`
)

var (
	userPasswordHashSelectStmt   *sql.Stmt
	userProfileSelectStmt        *sql.Stmt
	dictionariesSelectStmt       *sql.Stmt
	patientAppointmentSelectStmt *sql.Stmt
	patientInsertStmt            *sql.Stmt
	appointmentsSelectStmt       *sql.Stmt
	appointmentInsertStmt        *sql.Stmt
	appointmentUpdateStmt        *sql.Stmt
)

type dao struct {
	db *sql.DB
}

func New(cfg config.DBConfig) store.Dao {
	db, err := sql.Open("postgres", cfg.ConnectionString)
	if err != nil {
		log.Fatal(err)
	}
	userPasswordHashSelectStmt, err = db.Prepare(userPasswordHashSelect)
	if err != nil {
		log.Fatal(err)
	}
	userProfileSelectStmt, err = db.Prepare(userProfileSelect)
	if err != nil {
		log.Fatal(err)
	}
	dictionariesSelectStmt, err = db.Prepare(dictionariesSelect)
	if err != nil {
		log.Fatal(err)
	}
	patientAppointmentSelectStmt, err = db.Prepare(patientAppointmentSelect)
	if err != nil {
		log.Fatal(err)
	}
	patientInsertStmt, err = db.Prepare(patientInsert)
	if err != nil {
		log.Fatal(err)
	}
	appointmentInsertStmt, err = db.Prepare(appointmentInsert)
	if err != nil {
		log.Fatal(err)
	}
	appointmentUpdateStmt, err = db.Prepare(appointmentUpdate)
	if err != nil {
		log.Fatal(err)
	}
	appointmentsSelectStmt, err = db.Prepare(appointmentsSelect)
	if err != nil {
		log.Fatal(err)
	}
	return &dao{
		db: db,
	}
}

func (d *dao) Close() error {
	userPasswordHashSelectStmt.Close()
	userProfileSelectStmt.Close()
	dictionariesSelectStmt.Close()
	patientAppointmentSelectStmt.Close()
	patientInsertStmt.Close()
	appointmentsSelectStmt.Close()
	appointmentInsertStmt.Close()
	appointmentUpdateStmt.Close()
	d.db.Close()
	return nil
}

func (d *dao) Authenticate(login, password string) error {
	var hash string
	login = strings.ToLower(login)
	if err := userPasswordHashSelectStmt.QueryRow(login).Scan(&hash); err != nil {
		return err
	}
	if err := util.CompareHashAndText(hash, password); err == nil {
		return nil
	}
	return store.ErrUserNotFound
}

func (d *dao) GetProfile(login string) (store.Profile, error) {
	var (
		id   int64
		name string
	)
	err := userProfileSelectStmt.QueryRow(login).Scan(&id, &name)
	return store.Profile{
		Id:       id,
		UserName: name,
	}, err
}

func (d *dao) GetDictionaries() (map[string][]store.Dictionary, error) {
	rows, err := dictionariesSelectStmt.Query()
	if err != nil {
		return nil, err
	}
	result := make(map[string][]store.Dictionary)
	for rows.Next() {
		var (
			id         int64
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

func (d *dao) SearchAppointments(patientName string) ([]store.Appointment, error) {
	rows, err := appointmentsSelectStmt.Query("%" + patientName + "%")
	if err != nil {
		return nil, err
	}
	result := make([]store.Appointment, 0)
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
		result = append(result, store.Appointment{
			Id:          id,
			DoctorName:  doctorName,
			PatientName: patientName,
			DateReceipt: dateReceipt,
		})
	}
	return result, rows.Err()
}

func (d *dao) GetAppointment(id int64) (store.Appointment, error) {
	ap := store.Appointment{}
	err := patientAppointmentSelectStmt.QueryRow(id).Scan(
		&ap.Id,
		&ap.DateReceipt,
		&ap.DoctorId,
		&ap.PatientId,
		&ap.HowReceipt,
		&ap.Alergo,
		&ap.ContactInfected,
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

func (d *dao) SaveAppointment(ap *store.Appointment) error {
	tx, err := d.db.Begin()
	if err != nil {
		return err
	}
	if err := tx.Stmt(patientInsertStmt).QueryRow(strings.ToLower(ap.PatientName), ap.PatientName).Scan(&ap.PatientId); err != nil {
		tx.Rollback()
		return err
	}
	if ap.Id == 0 {
		err = d.insertAppointment(tx, ap)
	} else {
		err = d.updateAppointment(tx, ap)
	}
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

func (d *dao) insertAppointment(tx *sql.Tx, ap *store.Appointment) error {
	row := tx.Stmt(appointmentInsertStmt).QueryRow(
		ap.DateReceipt,
		ap.DoctorId,
		ap.PatientId,
		ap.HowReceipt,
		ap.Alergo,
		ap.ContactInfected,
		ap.Hiv,
		ap.Transfusion,
		ap.Dyscountry,
		ap.Smoking,
		ap.Drugs,
		ap.Inheritance,
		ap.Diseases,
		ap.Gyndiseases,
		ap.Paritet,
		ap.Pregnancy,
		ap.FirstTrimester,
		ap.SecondTrimester,
		ap.ThirdTrimester,
		ap.History,
		ap.ExpByMenstruation,
		ap.ExpByFirstVisit,
		ap.ExpByUltra,
		ap.HealthStateId,
		ap.Claims,
		ap.Head,
		ap.Vision,
		ap.SkinStateId,
		ap.Lymph,
		ap.Breath,
		ap.Rale,
		ap.Tones,
		ap.Pulse,
		ap.PulseType,
		ap.Pressure,
		ap.TongueClean,
		ap.TongueWet,
		ap.TongueDry,
		ap.TongueCoated,
		ap.TongueUncoated,
		ap.Throat,
		ap.Belly,
		ap.Peritoneal,
		ap.Labors,
		ap.Dysuric,
		ap.Bowel,
		ap.LimbSwelling,
		ap.FaceSwelling,
		ap.UteruseStateId,
		ap.FetalPositionId,
		ap.FetalPreviaId,
		ap.FetalAlignId,
		ap.FetalHeartbeatId,
		ap.FetalPulse,
		ap.ReproductiveDischargeId,
		ap.Vdm,
		ap.Oj,
		ap.Dspin,
		ap.Dcrist,
		ap.Dtroch,
		ap.Cext,
		ap.DevelOrgansId,
		ap.GenitalAnomalies,
		ap.VaginaStateId,
		ap.LenghtCervix,
		ap.TruncateCervix,
		ap.OuterThroatStateId,
		ap.ChannelCervix,
		ap.FetalBladder,
		ap.FetalBladderPreviaId,
		ap.FetalBladderAlignId,
		ap.Arches,
		ap.Conjugate,
		ap.PelvisStateId,
		ap.PelvisDischarge,
		ap.Diagnosis,
		ap.Conclusion,
		ap.BirthPlan)
	return row.Scan(&ap.Id)
}

func (d *dao) updateAppointment(tx *sql.Tx, ap *store.Appointment) error {
	row := tx.Stmt(appointmentUpdateStmt).QueryRow(
		ap.DateReceipt,
		ap.DoctorId,
		ap.PatientId,
		ap.HowReceipt,
		ap.Alergo,
		ap.ContactInfected,
		ap.Hiv,
		ap.Transfusion,
		ap.Dyscountry,
		ap.Smoking,
		ap.Drugs,
		ap.Inheritance,
		ap.Diseases,
		ap.Gyndiseases,
		ap.Paritet,
		ap.Pregnancy,
		ap.FirstTrimester,
		ap.SecondTrimester,
		ap.ThirdTrimester,
		ap.History,
		ap.ExpByMenstruation,
		ap.ExpByFirstVisit,
		ap.ExpByUltra,
		ap.HealthStateId,
		ap.Claims,
		ap.Head,
		ap.Vision,
		ap.SkinStateId,
		ap.Lymph,
		ap.Breath,
		ap.Rale,
		ap.Tones,
		ap.Pulse,
		ap.PulseType,
		ap.Pressure,
		ap.TongueClean,
		ap.TongueWet,
		ap.TongueDry,
		ap.TongueCoated,
		ap.TongueUncoated,
		ap.Throat,
		ap.Belly,
		ap.Peritoneal,
		ap.Labors,
		ap.Dysuric,
		ap.Bowel,
		ap.LimbSwelling,
		ap.FaceSwelling,
		ap.UteruseStateId,
		ap.FetalPositionId,
		ap.FetalPreviaId,
		ap.FetalAlignId,
		ap.FetalHeartbeatId,
		ap.FetalPulse,
		ap.ReproductiveDischargeId,
		ap.Vdm,
		ap.Oj,
		ap.Dspin,
		ap.Dcrist,
		ap.Dtroch,
		ap.Cext,
		ap.DevelOrgansId,
		ap.GenitalAnomalies,
		ap.VaginaStateId,
		ap.LenghtCervix,
		ap.TruncateCervix,
		ap.OuterThroatStateId,
		ap.ChannelCervix,
		ap.FetalBladder,
		ap.FetalBladderPreviaId,
		ap.FetalBladderAlignId,
		ap.Arches,
		ap.Conjugate,
		ap.PelvisStateId,
		ap.PelvisDischarge,
		ap.Diagnosis,
		ap.Conclusion,
		ap.BirthPlan,
		ap.Id)
	var cnt int
	if err := row.Scan(&cnt); err != nil {
		return err
	}
	if cnt != 1 {
		return store.ErrDataNotUpdated
	}
	return nil
}
