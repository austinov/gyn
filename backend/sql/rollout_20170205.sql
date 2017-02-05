-- rollout 05.02.2017

CREATE TABLE IF NOT EXISTS external_throat_states (
    "id"      serial primary key,
    "name"    varchar(100) NOT NULL,
    "orderby" integer
);
COMMENT ON TABLE oprv_states IS 'Состояние наружнего зева';
CREATE INDEX ind_external_throat_states_id ON external_throat_states USING btree (id);
CREATE UNIQUE INDEX uni_external_throat_states ON external_throat_states (lower(name));

INSERT INTO external_throat_states (name) VALUES ('закрыт'), ('пропускает кончик пальца');


ALTER TABLE appointments ADD COLUMN ultra_in_reception boolean;
COMMENT ON COLUMN appointments.ultra_in_reception IS 'УЗИ в приемном отделении';

ALTER TABLE appointments ADD COLUMN doppler_in_reception boolean;
COMMENT ON COLUMN appointments.doppler_in_reception IS 'Допплерометрия в приемном отделении';

--
ALTER TABLE appointments ADD COLUMN cervix_back boolean;
COMMENT ON COLUMN appointments.cervix_back IS 'Шейка матки отклонена кзади (0 баллов)';

ALTER TABLE appointments ADD COLUMN cervix_front boolean;
COMMENT ON COLUMN appointments.cervix_front IS 'Шейка матки отклонена кпереди (1 балл)';

ALTER TABLE appointments ADD COLUMN cervix_center boolean;
COMMENT ON COLUMN appointments.cervix_center IS 'Шейка матки центрирована (2 балла)';

ALTER TABLE appointments ADD COLUMN cervix_tight boolean;
COMMENT ON COLUMN appointments.cervix_tight IS 'Шейка матки плотная (0 балла)';

ALTER TABLE appointments ADD COLUMN cervix_middle_soft boolean;
COMMENT ON COLUMN appointments.cervix_middle_soft IS 'Шейка матки умеренно размягчена (1 балла)';

ALTER TABLE appointments ADD COLUMN cervix_soft boolean;
COMMENT ON COLUMN appointments.cervix_soft IS 'Шейка матки мягкая (2 балла)';

ALTER TABLE appointments ADD COLUMN cervix_length boolean;
COMMENT ON COLUMN appointments.cervix_length IS 'Шейка матки длиной (см)';

ALTER TABLE appointments ADD COLUMN cervix_channel boolean;
COMMENT ON COLUMN appointments.cervix_channel IS 'Цервикальный канал проходим';

ALTER TABLE appointments ADD COLUMN external_throat_state_id integer;
COMMENT ON COLUMN appointments.external_throat_state_id IS 'Состояние наружнего зева';
CREATE INDEX ind_appointments_external_throat_state_id ON appointments USING btree (external_throat_state_id);
ALTER TABLE appointments ADD CONSTRAINT fk_appointments_external_throat_state FOREIGN KEY (external_throat_state_id) REFERENCES external_throat_states (id);

CREATE INDEX ind_appointments_receipt_kind_id ON appointments USING btree (receipt_kind_id);
CREATE INDEX ind_appointments_pelvis_state_id ON appointments USING btree (pelvis_state_id);
CREATE INDEX ind_appointments_fetal_bladder_state_id ON appointments USING btree (fetal_bladder_state_id);
CREATE INDEX ind_appointments_fetal_bladder_align_id ON appointments USING btree (fetal_bladder_align_id);
CREATE INDEX ind_appointments_fetal_bladder_previa_id ON appointments USING btree (fetal_bladder_previa_id);
CREATE INDEX ind_appointments_vagina_state_id ON appointments USING btree (vagina_state_id);
CREATE INDEX ind_appointments_devel_organ_id ON appointments USING btree (devel_organs_id);
CREATE INDEX ind_appointments_reproductive_discharge_type_id ON appointments USING btree (reproductive_discharge_type_id);
CREATE INDEX ind_appointments_reproductive_discharge_state_id ON appointments USING btree (reproductive_discharge_state_id);
CREATE INDEX ind_appointments_fetal_align_id ON appointments USING btree (fetal_align_id);
CREATE INDEX ind_appointments_fetal_heartbeat_id ON appointments USING btree (fetal_heartbeat_id);
CREATE INDEX ind_appointments_fetal_previa_id ON appointments USING btree (fetal_previa_id);
CREATE INDEX ind_appointments_fetal_position_id ON appointments USING btree (fetal_position_id);
CREATE INDEX ind_appointments_uteruse_state_id ON appointments USING btree (uteruse_state_id);
CREATE INDEX ind_appointments_skin_state_id ON appointments USING btree (skin_state_id);
CREATE INDEX ind_appointments_health_state_id ON appointments USING btree (health_state_id);
CREATE INDEX ind_appointments_breath_state_id ON appointments USING btree (breath_state_id);
CREATE INDEX ind_appointments_rale_state_id ON appointments USING btree (rale_state_id);
CREATE INDEX ind_appointments_tones_state_id ON appointments USING btree (tones_state_id);
CREATE INDEX ind_appointments_belly_state_id ON appointments USING btree (belly_state_id);
CREATE INDEX ind_appointments_epigastrium_state_id ON appointments USING btree (epigastrium_state_id);
CREATE INDEX ind_appointments_scar_state_id ON appointments USING btree (scar_state_id);
CREATE INDEX ind_appointments_heartbeat_rithm_id ON appointments USING btree (heartbeat_rithm_id);
CREATE INDEX ind_appointments_pelvis_discharge_type_id ON appointments USING btree (pelvis_discharge_type_id);
CREATE INDEX ind_appointments_infection_markers_state_id ON appointments USING btree (infection_markers_state_id);
CREATE INDEX ind_appointments_tromboflebia_state_id ON appointments USING btree (tromboflebia_state_id);
CREATE INDEX ind_appointments_oprv_state_id ON appointments USING btree (oprv_state_id);

DROP VIEW vw_appointments;
CREATE OR REPLACE VIEW vw_appointments AS
SELECT a.*,
       u.name AS doctor_name,
       p.name AS patient_name,
       rk.name AS receipt_kind_name,
       ps.name AS pelvis_state_name,
       fbs.name AS fetal_bladder_state_name,
       fba.name AS fetal_bladder_align_name,
       fbp.name AS fetal_bladder_previa_name,
       vs.name AS vagina_state_name,
       dor.name AS devel_organs_name,
       rdt.name AS reproductive_discharge_type_name,
       rds.name AS reproductive_discharge_state_name,
       fa.name AS fetal_align_name,
       fh.name AS fetal_heartbeat_name,
       fpr.name AS fetal_previa_name,
       fpp.name AS fetal_position_name,
       us.name AS uteruse_state_name,
       ss.name AS skin_state_name,
       hs.name AS health_state_name,
       brs.name AS breath_state_name,
       rs.name AS rale_state_name,
       ts.name AS tones_state_name,
       bes.name AS belly_state_name,
       es.name AS epigastrium_state_name,
       scs.name AS scar_state_name,
       hbr.name AS heartbeat_rithm_name,
       pdt.name AS pelvis_discharge_type_name,
       pds.name AS pelvis_discharge_state_name,
       imes.name AS infection_markers_state_name,
       tes.name AS tromboflebia_state_name,
       ops.name AS oprv_state_name,
       ets.name AS external_throat_state_name
FROM appointments a
  JOIN users u ON a.doctor_id = u.id
  JOIN patients p ON a.patient_id = p.id
  LEFT JOIN receipt_kinds rk ON a.receipt_kind_id = rk.id
  LEFT JOIN pelvis_states ps ON a.pelvis_state_id = ps.id
  LEFT JOIN fetal_bladder_states fbs ON a.fetal_bladder_state_id = fbs.id
  LEFT JOIN fetal_bladder_aligns fba ON a.fetal_bladder_align_id = fba.id
  LEFT JOIN fetal_bladder_previas fbp ON a.fetal_bladder_previa_id = fbp.id
  LEFT JOIN vagina_states vs ON a.vagina_state_id = vs.id
  LEFT JOIN devel_organs dor ON a.devel_organs_id = dor.id
  LEFT JOIN discharge_types rdt ON a.reproductive_discharge_type_id = rdt.id
  LEFT JOIN discharge_states rds ON a.reproductive_discharge_state_id = rds.id
  LEFT JOIN fetal_aligns fa ON a.fetal_align_id = fa.id
  LEFT JOIN fetal_heartbeats fh ON a.fetal_heartbeat_id = fh.id
  LEFT JOIN fetal_previas fpr ON a.fetal_previa_id = fpr.id
  LEFT JOIN fetal_positions fpp ON a.fetal_position_id = fpp.id
  LEFT JOIN uteruse_states us ON a.uteruse_state_id = us.id
  LEFT JOIN skin_states ss ON a.skin_state_id = ss.id
  LEFT JOIN health_states hs ON a.health_state_id = hs.id
  LEFT JOIN breath_states brs ON a.breath_state_id = brs.id
  LEFT JOIN rale_states rs ON a.rale_state_id = rs.id
  LEFT JOIN tones_states ts ON a.tones_state_id = ts.id
  LEFT JOIN belly_states bes ON a.belly_state_id = bes.id
  LEFT JOIN belly_states es ON a.epigastrium_state_id = es.id
  LEFT JOIN belly_states scs ON a.scar_state_id = scs.id
  LEFT JOIN heartbeat_rithms hbr ON a.heartbeat_rithm_id = hbr.id
  LEFT JOIN discharge_types pdt ON a.pelvis_discharge_type_id = pdt.id
  LEFT JOIN discharge_states pds ON a.pelvis_discharge_state_id = pds.id
  LEFT JOIN examination_states imes ON a.infection_markers_state_id = imes.id
  LEFT JOIN examination_states tes ON a.tromboflebia_state_id = tes.id
  LEFT JOIN oprv_states ops ON a.oprv_state_id = ops.id
  LEFT JOIN external_throat_states ets ON a.external_throat_state_id = ets.id;
