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
		SELECT id, name, 'discharge_states' AS dict, orderby FROM discharge_states
      ) t
      ORDER BY t.dict,  t.orderby, t.id`

	patientAppointmentSelect = `
	  SELECT
        id, date_receipt, doctor_id, patient_id, receipt_kind_id, receipt_diagnosis, alergo,
		contact_infected, hiv, transfusion, dyscountry, smoking, drugs, inheritance, gyndiseases,
        paritet, paritet_b, paritet_p, paritet_a, paritet_sv, paritet_nb, paritet_eb,
		infection_markers, infection_markers_desc, tromboflebia, tromboflebia_desc,
		first_trimester, second_trimester, third_trimester, history, oprv, oprv_homo,
        exp_by_menstruation, exp_by_first_visit, exp_by_ultra_first, exp_by_ultra_second,
		exp_by_ultra_third, health_state_id, claims, head, vision, skin_state_id, lymph, 
		breath_state_id, rale_state_id, tones_state_id, pulse, pulse_type, pressure,
        tongue_clean, tongue_wet, tongue_dry, tongue_coated, tongue_uncoated, throat,
		belly_period, belly_state_id, epigastrium_state_use, epigastrium_state_id, 
		scar_state_use, scar_state_id, peritoneal, labors, dysuric, bowel, limb_swelling,
		uteruse_state_id, fetal_position_id, fetal_previa_id, fetal_align_id, fetal_heartbeat_id,
		heartbeat_rithm_id, fetal_pulse, reproductive_discharge_id, discharge_state_id,
		vdm, oj, dspin, dcrist, dtroch, cext, devel_organs_id, genital_anomalies, vagina_state_id,
		bishop, fetal_bladder_state_id, fetal_bladder_previa_id, fetal_bladder_align_id, arches,
        conjugate, pelvis_state_id, pelvis_exostosis, pelvis_discharge,
		diagnosis, conclusion, birth_plan_use, birth_plan,
        doctor_name, patient_name, receipt_kind_name, pelvis_state_name,
		fetal_bladder_state_name, fetal_bladder_align_name, fetal_bladder_previa_name,
		vagina_state_name, devel_organs_name, reproductive_discharge_name, discharge_state_name,
		fetal_align_name, fetal_heartbeat_name, fetal_previa_name,
        fetal_position_name, uteruse_state_name, skin_state_name, health_state_name,
		breath_state_name, rale_state_name, tones_state_name, belly_state_name,
		epigastrium_state_name, scar_state_name, heartbeat_rithm_name
	  FROM vw_appointments a
	  WHERE id = $1`

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
        date_receipt, doctor_id, patient_id, receipt_kind_id, receipt_diagnosis, alergo,
		contact_infected, hiv, transfusion, dyscountry, smoking, drugs, inheritance, gyndiseases,
        paritet, paritet_b, paritet_p, paritet_a, paritet_sv, paritet_nb, paritet_eb,
		infection_markers, infection_markers_desc, tromboflebia, tromboflebia_desc,
		first_trimester, second_trimester, third_trimester, history, oprv, oprv_homo,
        exp_by_menstruation, exp_by_first_visit, exp_by_ultra_first, exp_by_ultra_second,
		exp_by_ultra_third, health_state_id, claims, head, vision, skin_state_id, lymph,
		breath_state_id, rale_state_id, tones_state_id, pulse, pulse_type,
        pressure, tongue_clean, tongue_wet, tongue_dry, tongue_coated, tongue_uncoated,
        throat, belly_period, belly_state_id, epigastrium_state_use, epigastrium_state_id,
		scar_state_use, scar_state_id, peritoneal, labors, dysuric, bowel, limb_swelling,
        uteruse_state_id, fetal_position_id, fetal_previa_id, fetal_align_id, fetal_heartbeat_id,
		heartbeat_rithm_id, fetal_pulse, reproductive_discharge_id, discharge_state_id,
		vdm, oj, dspin, dcrist, dtroch, cext, devel_organs_id, genital_anomalies, vagina_state_id,
		bishop, fetal_bladder_state_id, fetal_bladder_previa_id, fetal_bladder_align_id,
		arches, conjugate, pelvis_state_id, pelvis_exostosis, pelvis_discharge,	diagnosis,
		conclusion, birth_plan_use, birth_plan)
      VALUES (
        $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19,
        $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35, $36,
        $37, $38, $39, $40, $41, $42, $43, $44, $45, $46, $47, $48, $49, $50, $51, $52, $53,
        $54, $55, $56, $57, $58, $59, $60, $61, $62, $63, $64, $65, $66, $67, $68, $69, $70,
        $71, $72, $73, $74, $75, $76, $77, $78, $79, $80, $81, $82, $83, $84, $85, $86, $87,
	    $88, $89, $90, $91, $92, $93, $94, $95, $96)
  	  RETURNING id`

	appointmentUpdate = `
	  WITH rows AS (
        UPDATE appointments
          SET date_receipt = $1, doctor_id = $2, patient_id = $3, receipt_kind_id = $4,
		      receipt_diagnosis = $5, alergo = $6, contact_infected = $7, hiv = $8,
			  transfusion = $9, dyscountry = $10, smoking = $11, drugs = $12, inheritance = $13,
			  gyndiseases = $14, paritet = $15, paritet_b = $16, paritet_p = $17, paritet_a = $18,
			  paritet_sv = $19, paritet_nb = $20, paritet_eb = $21, infection_markers = $22,
			  infection_markers_desc = $23, tromboflebia = $24, tromboflebia_desc = $25,
			  first_trimester = $26, second_trimester = $27, third_trimester = $28,
              history = $29, oprv = $30, oprv_homo = $31, exp_by_menstruation = $32,
			  exp_by_first_visit = $33, exp_by_ultra_first = $34, exp_by_ultra_second = $35,
			  exp_by_ultra_third = $36, health_state_id = $37, claims = $38, head = $39, vision = $40,
			  skin_state_id = $41, lymph = $42, breath_state_id = $43, rale_state_id = $44,
			  tones_state_id = $45, pulse = $46, pulse_type = $47, pressure = $48, tongue_clean = $49,
			  tongue_wet = $50, tongue_dry = $51, tongue_coated = $52, tongue_uncoated = $53,
			  throat = $54, belly_period = $55, belly_state_id = $56, epigastrium_state_use = $57,
			  epigastrium_state_id = $58,
			  scar_state_use = $59, scar_state_id = $60, peritoneal = $61, labors = $62, dysuric = $63,
			  bowel = $64, limb_swelling = $65, uteruse_state_id = $66, fetal_position_id = $67,
			  fetal_previa_id = $68, fetal_align_id = $69, fetal_heartbeat_id = $70, heartbeat_rithm_id = $71,
			  fetal_pulse = $72, reproductive_discharge_id = $73, discharge_state_id = $74, vdm = $75,
			  oj = $76, dspin = $77, dcrist = $78, dtroch = $79, cext = $80, devel_organs_id = $81,
			  genital_anomalies = $82, vagina_state_id = $83, bishop = $84, fetal_bladder_state_id = $85,
              fetal_bladder_previa_id = $86, fetal_bladder_align_id = $87, arches = $88, conjugate = $89,
			  pelvis_state_id = $90, pelvis_exostosis = $91, pelvis_discharge = $92, diagnosis = $93,
			  conclusion = $94, birth_plan_use = $95, birth_plan = $96
        WHERE id = $97
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
		&ap.ReceiptKindId,
		&ap.ReceiptDiagnosis,
		&ap.Alergo,
		&ap.ContactInfected,
		&ap.Hiv,
		&ap.Transfusion,
		&ap.Dyscountry,
		&ap.Smoking,
		&ap.Drugs,
		&ap.Inheritance,
		&ap.Gyndiseases,
		&ap.Paritet,
		&ap.ParitetB,
		&ap.ParitetP,
		&ap.ParitetA,
		&ap.ParitetSV,
		&ap.ParitetNB,
		&ap.ParitetEB,
		&ap.InfectionMarkers,
		&ap.InfectionMarkersDesc,
		&ap.Tromboflebia,
		&ap.TromboflebiaDesc,
		&ap.FirstTrimester,
		&ap.SecondTrimester,
		&ap.ThirdTrimester,
		&ap.History,
		&ap.Oprv,
		&ap.OprvHomo,
		&ap.ExpByMenstruation,
		&ap.ExpByFirstVisit,
		&ap.ExpByUltraFirst,
		&ap.ExpByUltraSecond,
		&ap.ExpByUltraThird,
		&ap.HealthStateId,
		&ap.Claims,
		&ap.Head,
		&ap.Vision,
		&ap.SkinStateId,
		&ap.Lymph,
		&ap.BreathStateId,
		&ap.RaleStateId,
		&ap.TonesStateId,
		&ap.Pulse,
		&ap.PulseType,
		&ap.Pressure,
		&ap.TongueClean,
		&ap.TongueWet,
		&ap.TongueDry,
		&ap.TongueCoated,
		&ap.TongueUncoated,
		&ap.Throat,
		&ap.BellyPeriod,
		&ap.BellyStateId,
		&ap.EpigastriumStateUse,
		&ap.EpigastriumStateId,
		&ap.ScarStateUse,
		&ap.ScarStateId,
		&ap.Peritoneal,
		&ap.Labors,
		&ap.Dysuric,
		&ap.Bowel,
		&ap.LimbSwelling,
		&ap.UteruseStateId,
		&ap.FetalPositionId,
		&ap.FetalPreviaId,
		&ap.FetalAlignId,
		&ap.FetalHeartbeatId,
		&ap.HeartbeatRithmId,
		&ap.FetalPulse,
		&ap.ReproductiveDischargeId,
		&ap.DischargeStateId,
		&ap.Vdm,
		&ap.Oj,
		&ap.Dspin,
		&ap.Dcrist,
		&ap.Dtroch,
		&ap.Cext,
		&ap.DevelOrgansId,
		&ap.GenitalAnomalies,
		&ap.VaginaStateId,
		&ap.Bishop,
		&ap.FetalBladderStateId,
		&ap.FetalBladderPreviaId,
		&ap.FetalBladderAlignId,
		&ap.Arches,
		&ap.Conjugate,
		&ap.PelvisStateId,
		&ap.PelvisDischarge,
		&ap.Diagnosis,
		&ap.Conclusion,
		&ap.BirthPlanUse,
		&ap.BirthPlan,
		&ap.DoctorName,
		&ap.PatientName,
		&ap.ReceiptKindName,
		&ap.PelvisStateName,
		&ap.FetalBladderStateName,
		&ap.FetalBladderAlignName,
		&ap.FetalBladderPreviaName,
		&ap.VaginaStateName,
		&ap.DevelOrgansName,
		&ap.ReproductiveDischargeName,
		&ap.DischargeStateName,
		&ap.FetalAlignName,
		&ap.FetalHeartbeatName,
		&ap.FetalPreviaName,
		&ap.FetalPositionName,
		&ap.UteruseStateName,
		&ap.SkinStateName,
		&ap.HealthStateName,
		&ap.BreathStateName,
		&ap.RaleStateName,
		&ap.TonesStateName,
		&ap.BellyStateName,
		&ap.EpigastriumStateName,
		&ap.ScarStateName,
		&ap.HeartbeatRithmName)
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
		ap.ReceiptKindId,
		ap.ReceiptDiagnosis,
		ap.Alergo,
		ap.ContactInfected,
		ap.Hiv,
		ap.Transfusion,
		ap.Dyscountry,
		ap.Smoking,
		ap.Drugs,
		ap.Inheritance,
		ap.Gyndiseases,
		ap.Paritet,
		ap.ParitetB,
		ap.ParitetP,
		ap.ParitetA,
		ap.ParitetSV,
		ap.ParitetNB,
		ap.ParitetEB,
		ap.InfectionMarkers,
		ap.InfectionMarkersDesc,
		ap.Tromboflebia,
		ap.TromboflebiaDesc,
		ap.FirstTrimester,
		ap.SecondTrimester,
		ap.ThirdTrimester,
		ap.History,
		ap.Oprv,
		ap.OprvHomo,
		ap.ExpByMenstruation,
		ap.ExpByFirstVisit,
		ap.ExpByUltraFirst,
		ap.ExpByUltraSecond,
		ap.ExpByUltraThird,
		ap.HealthStateId,
		ap.Claims,
		ap.Head,
		ap.Vision,
		ap.SkinStateId,
		ap.Lymph,
		ap.BreathStateId,
		ap.RaleStateId,
		ap.TonesStateId,
		ap.Pulse,
		ap.PulseType,
		ap.Pressure,
		ap.TongueClean,
		ap.TongueWet,
		ap.TongueDry,
		ap.TongueCoated,
		ap.TongueUncoated,
		ap.Throat,
		ap.BellyPeriod,
		ap.BellyStateId,
		ap.EpigastriumStateUse,
		ap.EpigastriumStateId,
		ap.ScarStateUse,
		ap.ScarStateId,
		ap.Peritoneal,
		ap.Labors,
		ap.Dysuric,
		ap.Bowel,
		ap.LimbSwelling,
		ap.UteruseStateId,
		ap.FetalPositionId,
		ap.FetalPreviaId,
		ap.FetalAlignId,
		ap.FetalHeartbeatId,
		ap.HeartbeatRithmId,
		ap.FetalPulse,
		ap.ReproductiveDischargeId,
		ap.DischargeStateId,
		ap.Vdm,
		ap.Oj,
		ap.Dspin,
		ap.Dcrist,
		ap.Dtroch,
		ap.Cext,
		ap.DevelOrgansId,
		ap.GenitalAnomalies,
		ap.VaginaStateId,
		ap.Bishop,
		ap.FetalBladderStateId,
		ap.FetalBladderPreviaId,
		ap.FetalBladderAlignId,
		ap.Arches,
		ap.Conjugate,
		ap.PelvisStateId,
		ap.PelvisExostosis,
		ap.PelvisDischarge,
		ap.Diagnosis,
		ap.Conclusion,
		ap.BirthPlanUse,
		ap.BirthPlan)
	return row.Scan(&ap.Id)
}

func (d *dao) updateAppointment(tx *sql.Tx, ap *store.Appointment) error {
	row := tx.Stmt(appointmentUpdateStmt).QueryRow(
		ap.DateReceipt,
		ap.DoctorId,
		ap.PatientId,
		ap.ReceiptKindId,
		ap.ReceiptDiagnosis,
		ap.Alergo,
		ap.ContactInfected,
		ap.Hiv,
		ap.Transfusion,
		ap.Dyscountry,
		ap.Smoking,
		ap.Drugs,
		ap.Inheritance,
		ap.Gyndiseases,
		ap.Paritet,
		ap.ParitetB,
		ap.ParitetP,
		ap.ParitetA,
		ap.ParitetSV,
		ap.ParitetNB,
		ap.ParitetEB,
		ap.InfectionMarkers,
		ap.InfectionMarkersDesc,
		ap.Tromboflebia,
		ap.TromboflebiaDesc,
		ap.FirstTrimester,
		ap.SecondTrimester,
		ap.ThirdTrimester,
		ap.History,
		ap.Oprv,
		ap.OprvHomo,
		ap.ExpByMenstruation,
		ap.ExpByFirstVisit,
		ap.ExpByUltraFirst,
		ap.ExpByUltraSecond,
		ap.ExpByUltraThird,
		ap.HealthStateId,
		ap.Claims,
		ap.Head,
		ap.Vision,
		ap.SkinStateId,
		ap.Lymph,
		ap.BreathStateId,
		ap.RaleStateId,
		ap.TonesStateId,
		ap.Pulse,
		ap.PulseType,
		ap.Pressure,
		ap.TongueClean,
		ap.TongueWet,
		ap.TongueDry,
		ap.TongueCoated,
		ap.TongueUncoated,
		ap.Throat,
		ap.BellyPeriod,
		ap.BellyStateId,
		ap.EpigastriumStateUse,
		ap.EpigastriumStateId,
		ap.ScarStateUse,
		ap.ScarStateId,
		ap.Peritoneal,
		ap.Labors,
		ap.Dysuric,
		ap.Bowel,
		ap.LimbSwelling,
		ap.UteruseStateId,
		ap.FetalPositionId,
		ap.FetalPreviaId,
		ap.FetalAlignId,
		ap.FetalHeartbeatId,
		ap.HeartbeatRithmId,
		ap.FetalPulse,
		ap.ReproductiveDischargeId,
		ap.DischargeStateId,
		ap.Vdm,
		ap.Oj,
		ap.Dspin,
		ap.Dcrist,
		ap.Dtroch,
		ap.Cext,
		ap.DevelOrgansId,
		ap.GenitalAnomalies,
		ap.VaginaStateId,
		ap.Bishop,
		ap.FetalBladderStateId,
		ap.FetalBladderPreviaId,
		ap.FetalBladderAlignId,
		ap.Arches,
		ap.Conjugate,
		ap.PelvisStateId,
		ap.PelvisExostosis,
		ap.PelvisDischarge,
		ap.Diagnosis,
		ap.Conclusion,
		ap.BirthPlanUse,
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
