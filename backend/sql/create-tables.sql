CREATE TABLE IF NOT EXISTS users (
    "id"       serial primary key,
    "login"    varchar(100) NOT NULL,
    "name"     varchar(100) NOT NULL,
    "psw_hash" varchar(100) NOT NULL
);
COMMENT ON TABLE users IS 'Пользователи приложения';
CREATE INDEX ind_users_id ON users USING btree (id);
CREATE UNIQUE INDEX uni_users ON users (lower(login));
CREATE INDEX ind_users_psw ON users USING btree (login, psw_hash);

INSERT INTO users (login, name, psw_hash) VALUES ('alex', 'Алексей', '$2a$10$fpowt7ANNV0WJJ2wANxG9ONl5nzunsAAmWqN94foN411iZzpYEYhG');
INSERT INTO users (login, name, psw_hash) VALUES ('andrey', 'Андрей', '$2a$10$zUkMSQRJztTgDNHmWurGKuc9pgyYPTjTOToJpDFH5uTJBB.rsNPxK');

CREATE TABLE IF NOT EXISTS patients (
    "id"   serial primary key,
    "name" varchar(100) NOT NULL
);
COMMENT ON TABLE patients IS 'Пациенты';
CREATE INDEX ind_patients_id ON patients USING btree (id);
CREATE UNIQUE INDEX uni_patients ON patients (lower(name));

CREATE TABLE IF NOT EXISTS receipt_kind (
    "id"      serial primary key,
    "name"    varchar(100) NOT NULL,
    "orderby" integer
);
COMMENT ON TABLE receipt_kind IS 'Типы поступления пациента';
CREATE INDEX ind_receipt_kind_id ON receipt_kind USING btree (id);
CREATE UNIQUE INDEX uni_receipt_kind ON receipt_kind (lower(name));

INSERT INTO receipt_kind (name) VALUES ('самотеком'), ('по наряду БСМП'), ('по направлению ЖК'), ('МОНИИАГ');

CREATE TABLE IF NOT EXISTS health_states (
    "id"      serial primary key,
    "name"    varchar(100) NOT NULL,
    "orderby" integer
);
COMMENT ON TABLE health_states IS 'Типы состояния пациента';
CREATE INDEX ind_health_states_id ON health_states USING btree (id);
CREATE UNIQUE INDEX uni_health_states ON health_states (lower(name));

INSERT INTO health_states (name) VALUES ('удовлетворительное'), ('относительно удовлетворительное'), ('средней тяжести'), ('тяжелое');

CREATE TABLE IF NOT EXISTS skin_states (
    "id"      serial primary key,
    "name"    varchar(100) NOT NULL,
    "orderby" integer
);
COMMENT ON TABLE skin_states IS 'Типы состояния кожных покровов';
CREATE INDEX ind_skin_states_id ON skin_states USING btree (id);
CREATE UNIQUE INDEX uni_skin_states ON skin_states (lower(name));

INSERT INTO skin_states (name) VALUES ('бледно-розовой окраски'), ('бледные');

CREATE TABLE IF NOT EXISTS breath_states (
    "id"      serial primary key,
    "name"    varchar(100) NOT NULL,
    "orderby" integer
);
COMMENT ON TABLE breath_states IS 'Типы состояния дыхания';
CREATE INDEX ind_breath_states_id ON breath_states USING btree (id);
CREATE UNIQUE INDEX uni_breath_states ON breath_states (lower(name));

INSERT INTO breath_states (name) VALUES ('везиулярное'), ('с жестким оттенком'), ('жесткое');

CREATE TABLE IF NOT EXISTS rale_states (
    "id"      serial primary key,
    "name"    varchar(100) NOT NULL,
    "orderby" integer
);
COMMENT ON TABLE rale_states IS 'Типы хрипов';
CREATE INDEX ind_rale_states_id ON rale_states USING btree (id);
CREATE UNIQUE INDEX uni_rale_states ON rale_states (lower(name));

INSERT INTO rale_states (name) VALUES ('нет'), ('сухие'), ('влажные');

CREATE TABLE IF NOT EXISTS tones_states (
    "id"      serial primary key,
    "name"    varchar(100) NOT NULL,
    "orderby" integer
);
COMMENT ON TABLE tones_states IS 'Типы состояния тонов cor';
CREATE INDEX ind_tones_states_id ON tones_states USING btree (id);
CREATE UNIQUE INDEX uni_tones_states ON tones_states (lower(name));

INSERT INTO tones_states (name) VALUES ('ясные ритмичные'), ('приглушенные');

CREATE TABLE IF NOT EXISTS belly_states (
    "id"      serial primary key,
    "name"    varchar(100) NOT NULL,
    "orderby" integer
);
COMMENT ON TABLE belly_states IS 'Типы состояния живота';
CREATE INDEX ind_belly_states_id ON belly_states USING btree (id);
CREATE UNIQUE INDEX uni_belly_states ON belly_states (lower(name));

INSERT INTO belly_states (name) VALUES ('безболезненный при пальпации'), ('болезненный при пальпации');

CREATE TABLE IF NOT EXISTS uteruse_states (
    "id"      serial primary key,
    "name"    varchar(100) NOT NULL,
    "orderby" integer
);
COMMENT ON TABLE uteruse_states IS 'Типы состояния матки';
CREATE INDEX ind_uteruse_states_id ON uteruse_states USING btree (id);
CREATE UNIQUE INDEX uni_uteruse_states ON uteruse_states (lower(name));

INSERT INTO uteruse_states (name) VALUES ('невозбудима'), ('слегка возбудима'), ('возбудима'), ('в гипертонусе');

CREATE TABLE IF NOT EXISTS fetal_positions (
    "id"      serial primary key,
    "name"    varchar(100) NOT NULL,
    "orderby" integer
);
COMMENT ON TABLE fetal_positions IS 'Типы состояния положения плода';
CREATE INDEX ind_fetal_positions_id ON fetal_positions USING btree (id);
CREATE UNIQUE INDEX uni_fetal_positions ON fetal_positions (lower(name));

INSERT INTO fetal_positions (name) VALUES ('продольное'), ('косое'), ('поперечное');

CREATE TABLE IF NOT EXISTS fetal_previas (
    "id"      serial primary key,
    "name"    varchar(100) NOT NULL,
    "orderby" integer
);
COMMENT ON TABLE fetal_previas IS 'Типы предлежаний плода';
CREATE INDEX ind_fetal_previas_id ON fetal_previas USING btree (id);
CREATE UNIQUE INDEX uni_fetal_previas ON fetal_previas (lower(name));

INSERT INTO fetal_previas (name) VALUES ('головка'), ('тазовый конец'), ('смешанное'), ('ножки плода');

CREATE TABLE IF NOT EXISTS fetal_aligns (
    "id"      serial primary key,
    "name"    varchar(100) NOT NULL,
    "orderby" integer
);
COMMENT ON TABLE fetal_aligns IS 'Типы выравниваний плода';
CREATE INDEX ind_fetal_aligns_id ON fetal_aligns USING btree (id);
CREATE UNIQUE INDEX uni_fetal_aligns ON fetal_aligns (lower(name));

INSERT INTO fetal_aligns (name) VALUES ('над входом в малый таз'), ('прижата ко входу в малый таз');

CREATE TABLE IF NOT EXISTS fetal_heartbeats (
    "id"      serial primary key,
    "name"    varchar(100) NOT NULL,
    "orderby" integer
);
COMMENT ON TABLE fetal_heartbeats IS 'Типы сердцебиения плода';
CREATE INDEX ind_fetal_heartbeats_id ON fetal_heartbeats USING btree (id);
CREATE UNIQUE INDEX uni_fetal_heartbeats ON fetal_heartbeats (lower(name));

INSERT INTO fetal_heartbeats (name) VALUES ('ясное'), ('приглушенное'), ('глухое');

CREATE TABLE IF NOT EXISTS heartbeat_rithms (
    "id"      serial primary key,
    "name"    varchar(100) NOT NULL,
    "orderby" integer
);
COMMENT ON TABLE heartbeat_rithms IS 'Типы ритмичности сердцебиения';
CREATE INDEX ind_heartbeat_rithms_id ON heartbeat_rithms USING btree (id);
CREATE UNIQUE INDEX uni_heartbeat_rithms ON heartbeat_rithms (lower(name));

INSERT INTO heartbeat_rithms (name) VALUES ('ритмичные'), ('аритмичные');

CREATE TABLE IF NOT EXISTS reproductive_discharges (
    "id"      serial primary key,
    "name"    varchar(100) NOT NULL,
    "orderby" integer
);
COMMENT ON TABLE reproductive_discharges IS 'Типы выделений из половых путей';
CREATE INDEX ind_reproductive_discharges_id ON reproductive_discharges USING btree (id);
CREATE UNIQUE INDEX uni_reproductive_discharges ON reproductive_discharges (lower(name));

INSERT INTO reproductive_discharges (name) VALUES ('светлые'), ('слизистые'), ('кровяные');

CREATE TABLE IF NOT EXISTS discharge_states (
    "id"      serial primary key,
    "name"    varchar(100) NOT NULL,
    "orderby" integer
);
COMMENT ON TABLE discharge_states IS 'Состояние выделений';
CREATE INDEX ind_discharge_states_id ON discharge_states USING btree (id);
CREATE UNIQUE INDEX uni_discharge_states ON discharge_states (lower(name));

INSERT INTO discharge_states (name) VALUES ('умеренные'), ('обильные'), ('скудные');

CREATE TABLE IF NOT EXISTS devel_organs (
    "id"      serial primary key,
    "name"    varchar(100) NOT NULL,
    "orderby" integer
);
COMMENT ON TABLE devel_organs IS 'Типы развития наружних половых органов';
CREATE INDEX ind_devel_organs_id ON devel_organs USING btree (id);
CREATE UNIQUE INDEX uni_devel_organs ON devel_organs (lower(name));

INSERT INTO devel_organs (name) VALUES ('правильно, аномалий нет'), ('аномалии');

CREATE TABLE IF NOT EXISTS vagina_states (
    "id"      serial primary key,
    "name"    varchar(100) NOT NULL,
    "orderby" integer
);
COMMENT ON TABLE vagina_states IS 'Типы состояний влагалища';
CREATE INDEX ind_vagina_states_id ON vagina_states USING btree (id);
CREATE UNIQUE INDEX uni_vagina_states ON vagina_states (lower(name));

INSERT INTO vagina_states (name) VALUES ('рожавшей'), ('не рожавшей');

CREATE TABLE IF NOT EXISTS fetal_bladder_previas (
    "id"      serial primary key,
    "name"    varchar(100) NOT NULL,
    "orderby" integer
);
COMMENT ON TABLE fetal_bladder_previas IS 'Типы предлежания плодный пузырь';
CREATE INDEX ind_fetal_bladder_previas_id ON fetal_bladder_previas USING btree (id);
CREATE UNIQUE INDEX uni_fetal_bladder_previas ON fetal_bladder_previas (lower(name));

INSERT INTO fetal_bladder_previas (name) VALUES ('головка'), ('ягодицы'), ('ножки'), ('тазовый конец');

CREATE TABLE IF NOT EXISTS fetal_bladder_states (
    "id"      serial primary key,
    "name"    varchar(100) NOT NULL,
    "orderby" integer
);
COMMENT ON TABLE fetal_bladder_states IS 'Состояние плодного пузыря';
CREATE INDEX ind_fetal_bladder_states_id ON fetal_bladder_states USING btree (id);
CREATE UNIQUE INDEX uni_fetal_bladder_states ON fetal_bladder_states (lower(name));

INSERT INTO fetal_bladder_states (name) VALUES ('цел'), ('оболочка не определяющаяся'), ('оболочка определяющаяся');

CREATE TABLE IF NOT EXISTS fetal_bladder_aligns (
    "id"      serial primary key,
    "name"    varchar(100) NOT NULL,
    "orderby" integer
);
COMMENT ON TABLE fetal_bladder_aligns IS 'Типы прижатия плодного пузыря';
CREATE INDEX ind_fetal_bladder_aligns_id ON fetal_bladder_aligns USING btree (id);
CREATE UNIQUE INDEX uni_fetal_bladder_aligns ON fetal_bladder_aligns (lower(name));

INSERT INTO fetal_bladder_aligns (name) VALUES ('прижата'), ('подвижна'), ('над входом в малый таз');

CREATE TABLE IF NOT EXISTS pelvis_states (
    "id"      serial primary key,
    "name"    varchar(100) NOT NULL,
    "orderby" integer
);
COMMENT ON TABLE pelvis_states IS 'Типы состояний костного таза';
CREATE INDEX ind_pelvis_states_id ON pelvis_states USING btree (id);
CREATE UNIQUE INDEX uni_pelvis_states ON pelvis_states (lower(name));

INSERT INTO pelvis_states (name) VALUES ('без экзостозов'), ('экзостозы');

CREATE TABLE IF NOT EXISTS appointments (
    "id"                        serial primary key,
    "date_receipt"              bigint,
    "doctor_id"                 integer,
    "patient_id"                integer,
    "receipt_kind_id"           integer,
    "receipt_diagnosis"         varchar(500),
    "alergo"                    varchar(500),
    "contact_infected"          varchar(500),
    "hiv"                       varchar(500),
    "transfusion"               varchar(500),
    "dyscountry"                varchar(500),
    "smoking"                   varchar(500),
    "drugs"                     varchar(500),
    "inheritance"               varchar(500),
    "gyndiseases"               varchar(100),
    "paritet"                   varchar(500),
    "paritet_b"                 varchar(500),
    "paritet_p"                 varchar(500),
    "paritet_a"                 varchar(500),
    "paritet_sv"                varchar(500),
    "paritet_nb"                varchar(500),
    "paritet_eb"                varchar(500),
    "pregnancy"                 text,
    "infection_markers"         boolean,
    "infection_markers_desc"    varchar(100),
    "tromboflebia"              boolean,
    "tromboflebia_desc"         varchar(100),
    "first_trimester"           text,
    "second_trimester"          text,
    "third_trimester"           text,
    "history"                   text,
    "oprv"                      varchar(100),
    "oprv_homo"                 boolean,
    "exp_by_menstruation"       varchar(100),
    "exp_by_first_visit"        varchar(100),
    "exp_by_ultra_first"        varchar(100),
    "exp_by_ultra_second"       varchar(100),
    "exp_by_ultra_third"        varchar(100),
    "health_state_id"           integer,
    "claims"                    varchar(100),
    "head"                      varchar(100),
    "vision"                    varchar(100),
    "skin_state_id"             integer,
    "lymph"                     varchar(100),
    "breath_state_id"           integer,
    "rale_state_id"             integer,
    "tones_state_id"            integer,
    "pulse"                     varchar(100),
    "pulse_type"                varchar(100),
    "pressure"                  varchar(100),
    "tongue_clean"              boolean,
    "tongue_wet"                boolean,
    "tongue_dry"                boolean,
    "tongue_coated"             boolean,
    "tongue_uncoated"           boolean,
    "throat"                    varchar(100),
    "belly_period"              varchar(100),
    "belly_state_id"            integer,
    "epigastrium_state_id"      integer,
    "scar_state_id"             integer,
    "peritoneal"                varchar(100),
    "labors"                    varchar(100),
    "dysuric"                   boolean,
    "bowel"                     boolean,
    "limb_swelling"             varchar(100),
    "uteruse_state_id"          integer,
    "fetal_position_id"         integer,
    "fetal_previa_id"           integer,
    "fetal_align_id"            integer,
    "fetal_heartbeat_id"        integer,
    "heartbeat_rithm_id"        integer,
    "fetal_pulse"               varchar(100),
    "reproductive_discharge_id" integer,
    "discharge_state_id"        integer,
    "vdm"                       varchar(20),
    "oj"                        varchar(20),
    "dspin"                     varchar(20),
    "dcrist"                    varchar(20),
    "dtroch"                    varchar(20),
    "cext"                      varchar(20),
    "devel_organs_id"           integer,
    "genital_anomalies"         varchar(100),
    "vagina_state_id"           integer,
    "bishop"                    text,
    "fetal_bladder_state_id"    integer,
    "fetal_bladder_previa_id"   integer,
    "fetal_bladder_align_id"    integer,
    "arches"                    varchar(100),
    "conjugate"                 varchar(50),
    "pelvis_state_id"           integer,
    "pelvis_exostosis"          varchar(100),
    "pelvis_discharge"          varchar(100),
    "diagnosis"                 text,
    "conclusion"                text,
    "birth_plan_exist"          boolean,
    "birth_plan"                text
);

COMMENT ON TABLE appointments IS 'Осмотры пациентов';
COMMENT ON COLUMN appointments.date_receipt IS 'Дата приёма, Unix timestamp in seconds';
COMMENT ON COLUMN appointments.doctor_id IS 'Дежурный врач акушер-гинеколог';
COMMENT ON COLUMN appointments.patient_id IS 'Пациент';
COMMENT ON COLUMN appointments.receipt_kind_id IS 'Тип поступиления пациента'; -- +
COMMENT ON COLUMN appointments.receipt_diagnosis IS 'Диагноз при поступилении'; -- +
COMMENT ON COLUMN appointments.alergo IS 'Аллергоанамнез';
COMMENT ON COLUMN appointments.contact_infected IS 'Контакт с инфекционными больными';
COMMENT ON COLUMN appointments.hiv IS 'Малярия, туберкулез, гепатиты, ВИЧ';
COMMENT ON COLUMN appointments.transfusion IS 'Гематрансфузия';
COMMENT ON COLUMN appointments.dyscountry IS 'Пребывание в неблагополучных странах в течении 3-х лет';
COMMENT ON COLUMN appointments.smoking IS 'Табакокурение';
COMMENT ON COLUMN appointments.drugs IS 'Прием наркотических, психотропных препаратов';
COMMENT ON COLUMN appointments.inheritance IS 'Наследственность';
COMMENT ON COLUMN appointments.gyndiseases IS 'Гинекологические заболевания';
COMMENT ON COLUMN appointments.paritet IS 'Паритет';
COMMENT ON COLUMN appointments.paritet_b IS 'Паритет - кол-во беременностей'; -- +
COMMENT ON COLUMN appointments.paritet_p IS 'Паритет - кол-во родов'; -- +
COMMENT ON COLUMN appointments.paritet_a IS 'Паритет - кол-во абортов'; -- +
COMMENT ON COLUMN appointments.paritet_sv IS 'Паритет - кол-во самопроизвольных выкидышей'; -- +
COMMENT ON COLUMN appointments.paritet_nb IS 'Паритет - кол-во неразвивающихся беременностей'; -- +
COMMENT ON COLUMN appointments.paritet_eb IS 'Паритет - кол-во эктопических беременностей'; -- +
COMMENT ON COLUMN appointments.pregnancy IS 'Течение беременности';
COMMENT ON COLUMN appointments.infection_markers IS 'Обследование на инфекционные маркеры'; -- +
COMMENT ON COLUMN appointments.infection_markers_desc IS 'Текст обследования на инфекционные маркеры'; -- +
COMMENT ON COLUMN appointments.tromboflebia IS 'Обследование на наследственную тромбофлебию'; -- +
COMMENT ON COLUMN appointments.tromboflebia_desc IS 'Описание обследования на наследственную тромбофлебию'; -- +
COMMENT ON COLUMN appointments.first_trimester IS 'I триместр';
COMMENT ON COLUMN appointments.second_trimester IS 'II триместр';
COMMENT ON COLUMN appointments.third_trimester IS 'III триместр';
COMMENT ON COLUMN appointments.history IS 'Из анамнеза';
COMMENT ON COLUMN appointments.oprv IS 'ОПРВ'; -- +
COMMENT ON COLUMN appointments.oprv_homo IS 'ОПРВ (не)равномерно'; -- +
COMMENT ON COLUMN appointments.exp_by_menstruation IS 'Сроки беременности по менструации';
COMMENT ON COLUMN appointments.exp_by_first_visit IS 'Сроки беременности по 1 явке';
COMMENT ON COLUMN appointments.exp_by_ultra_first IS 'Сроки беременности по первому УЗИ'; -- +
COMMENT ON COLUMN appointments.exp_by_ultra_second IS 'Сроки беременности по второму УЗИ'; -- +
COMMENT ON COLUMN appointments.exp_by_ultra_third IS 'Сроки беременности по третьему УЗИ'; -- +
COMMENT ON COLUMN appointments.health_state_id IS 'Состояние';
COMMENT ON COLUMN appointments.claims IS 'Жалобы';
COMMENT ON COLUMN appointments.head IS 'Голова';
COMMENT ON COLUMN appointments.vision IS 'Зрение';
COMMENT ON COLUMN appointments.skin_state_id IS 'Кожные покровы';
COMMENT ON COLUMN appointments.lymph IS 'Лимфоузлы';
COMMENT ON COLUMN appointments.breath_state_id IS 'Состояние дыхания'; -- +
COMMENT ON COLUMN appointments.rale_state_id IS 'Типы хрипов'; -- +
COMMENT ON COLUMN appointments.tones_state_id IS 'Типы тонов cor'; -- +
COMMENT ON COLUMN appointments.pulse IS 'Пульс';
COMMENT ON COLUMN appointments.pulse_type IS 'Ритмичность пульса';
COMMENT ON COLUMN appointments.pressure IS 'АД';
COMMENT ON COLUMN appointments.tongue_clean IS 'Язык чистый';
COMMENT ON COLUMN appointments.tongue_wet IS 'Язык влажный';
COMMENT ON COLUMN appointments.tongue_dry IS 'Язык сухой';
COMMENT ON COLUMN appointments.tongue_coated IS 'Язык обложен';
COMMENT ON COLUMN appointments.tongue_uncoated IS 'Язык не обложен';
COMMENT ON COLUMN appointments.throat IS 'Осмотр зева';
COMMENT ON COLUMN appointments.belly_period IS 'Живот соответствует периоду'; -- +
COMMENT ON COLUMN appointments.belly_state_id IS 'Состояние живота'; -- +
COMMENT ON COLUMN appointments.epigastrium_state_id IS 'Состояние область эпигастрия'; -- +
COMMENT ON COLUMN appointments.scar_state_id IS 'Состояние области послеоперационного рубца'; -- +
COMMENT ON COLUMN appointments.peritoneal IS 'Перитонеальные симптомы';
COMMENT ON COLUMN appointments.labors IS 'Родовая деятельность';
COMMENT ON COLUMN appointments.dysuric IS 'Дизурические явления (нет/есть)';
COMMENT ON COLUMN appointments.bowel IS 'Стул (не регулярный/регулярный)';
COMMENT ON COLUMN appointments.limb_swelling IS 'Отеки';
COMMENT ON COLUMN appointments.uteruse_state_id IS 'Матка';
COMMENT ON COLUMN appointments.fetal_position_id IS 'Положение плода';
COMMENT ON COLUMN appointments.fetal_previa_id IS 'Предлежит плода';
COMMENT ON COLUMN appointments.fetal_align_id IS 'Выравнивание плода';
COMMENT ON COLUMN appointments.fetal_heartbeat_id IS 'Сердцебиение плода';
COMMENT ON COLUMN appointments.heartbeat_rithm_id IS 'Ритмичность сердцебиения'; -- +
COMMENT ON COLUMN appointments.fetal_pulse IS 'Пульс плода';
COMMENT ON COLUMN appointments.reproductive_discharge_id IS 'Выделения из половых путей';
COMMENT ON COLUMN appointments.discharge_state_id IS 'Состояние выделений'; -- +
COMMENT ON COLUMN appointments.vdm IS 'ВДМ';
COMMENT ON COLUMN appointments.oj IS 'ОЖ';
COMMENT ON COLUMN appointments.dspin IS 'D.spin';
COMMENT ON COLUMN appointments.dcrist IS 'D.crist';
COMMENT ON COLUMN appointments.dtroch IS 'D.troch';
COMMENT ON COLUMN appointments.cext IS 'C.ext';
COMMENT ON COLUMN appointments.devel_organs_id IS 'Наружние половые органы развиты';
COMMENT ON COLUMN appointments.genital_anomalies IS 'Аномалии';
COMMENT ON COLUMN appointments.vagina_state_id IS 'Влагалище';
COMMENT ON COLUMN appointments.bishop IS 'Оценка по Бишопу'; -- +
COMMENT ON COLUMN appointments.fetal_bladder_state_id IS 'Состояния плодного пузыря'; -- +
COMMENT ON COLUMN appointments.fetal_bladder_previa_id IS 'Плодный пузырь предлежит';
COMMENT ON COLUMN appointments.fetal_bladder_align_id IS 'Плодный пузырь прижат';
COMMENT ON COLUMN appointments.arches IS 'Своды';
COMMENT ON COLUMN appointments.conjugate IS 'Диагональная коньюгата';
COMMENT ON COLUMN appointments.pelvis_state_id IS 'Костный таз';
COMMENT ON COLUMN appointments.pelvis_discharge IS 'Выделения (костный таз)';
COMMENT ON COLUMN appointments.diagnosis IS 'Диагноз';
COMMENT ON COLUMN appointments.conclusion IS 'Заключение';
COMMENT ON COLUMN appointments.birth_plan_exist IS 'Использование плана родов'; -- +
COMMENT ON COLUMN appointments.birth_plan IS 'План родов';

CREATE INDEX ind_appointments_id ON appointments USING btree (id);
CREATE INDEX ind_appointments_date ON appointments USING btree (date_receipt);
CREATE INDEX ind_appointments_doctor_id ON appointments USING btree (doctor_id);
CREATE INDEX ind_appointments_patient_id ON appointments USING btree (patient_id);
ALTER TABLE appointments ADD CONSTRAINT fk_appointments_doctor FOREIGN KEY (doctor_id) REFERENCES users (id);
ALTER TABLE appointments ADD CONSTRAINT fk_appointments_patient FOREIGN KEY (patient_id) REFERENCES patients (id);
ALTER TABLE appointments ADD CONSTRAINT fk_appointments_receipt_kind FOREIGN KEY (receipt_kind_id) REFERENCES receipt_kind (id);
ALTER TABLE appointments ADD CONSTRAINT fk_appointments_pelvis_states FOREIGN KEY (pelvis_state_id) REFERENCES pelvis_states (id);
ALTER TABLE appointments ADD CONSTRAINT fk_appointments_fetal_bladder_states FOREIGN KEY (fetal_bladder_state_id) REFERENCES fetal_bladder_states (id);
ALTER TABLE appointments ADD CONSTRAINT fk_appointments_fetal_bladder_aligns FOREIGN KEY (fetal_bladder_align_id) REFERENCES fetal_bladder_aligns (id);
ALTER TABLE appointments ADD CONSTRAINT fk_appointments_fetal_bladder_previas FOREIGN KEY (fetal_bladder_previa_id) REFERENCES fetal_bladder_previas (id);
ALTER TABLE appointments ADD CONSTRAINT fk_appointments_vagina_states FOREIGN KEY (vagina_state_id) REFERENCES vagina_states (id);
ALTER TABLE appointments ADD CONSTRAINT fk_appointments_devel_organs FOREIGN KEY (devel_organs_id) REFERENCES devel_organs (id);
ALTER TABLE appointments ADD CONSTRAINT fk_appointments_reproductive_discharges FOREIGN KEY (reproductive_discharge_id) REFERENCES reproductive_discharges (id);
ALTER TABLE appointments ADD CONSTRAINT fk_appointments_discharge_states FOREIGN KEY (discharge_state_id) REFERENCES discharge_states (id);
ALTER TABLE appointments ADD CONSTRAINT fk_appointments_fetal_aligns FOREIGN KEY (fetal_align_id) REFERENCES fetal_aligns (id);
ALTER TABLE appointments ADD CONSTRAINT fk_appointments_fetal_heartbeats FOREIGN KEY (fetal_heartbeat_id) REFERENCES fetal_heartbeats (id);
ALTER TABLE appointments ADD CONSTRAINT fk_appointments_fetal_previas FOREIGN KEY (fetal_previa_id) REFERENCES fetal_previas (id);
ALTER TABLE appointments ADD CONSTRAINT fk_appointments_fetal_positions FOREIGN KEY (fetal_position_id) REFERENCES fetal_positions (id);
ALTER TABLE appointments ADD CONSTRAINT fk_appointments_uteruse_states FOREIGN KEY (uteruse_state_id) REFERENCES uteruse_states (id);
ALTER TABLE appointments ADD CONSTRAINT fk_appointments_skin_states FOREIGN KEY (skin_state_id) REFERENCES skin_states (id);
ALTER TABLE appointments ADD CONSTRAINT fk_appointments_health_states FOREIGN KEY (health_state_id) REFERENCES health_states (id);
ALTER TABLE appointments ADD CONSTRAINT fk_appointments_breath_states FOREIGN KEY (breath_state_id) REFERENCES breath_states (id);
ALTER TABLE appointments ADD CONSTRAINT fk_appointments_rale_states FOREIGN KEY (rale_state_id) REFERENCES rale_states (id);
ALTER TABLE appointments ADD CONSTRAINT fk_appointments_tones_states FOREIGN KEY (tones_state_id) REFERENCES tones_states (id);
ALTER TABLE appointments ADD CONSTRAINT fk_appointments_belly_states FOREIGN KEY (belly_state_id) REFERENCES belly_states (id);
ALTER TABLE appointments ADD CONSTRAINT fk_appointments_epigastrium_states FOREIGN KEY (epigastrium_state_id) REFERENCES belly_states (id);
ALTER TABLE appointments ADD CONSTRAINT fk_appointments_scar_states FOREIGN KEY (scar_state_id) REFERENCES belly_states (id);
ALTER TABLE appointments ADD CONSTRAINT fk_appointments_heartbeat_rithms FOREIGN KEY (heartbeat_rithm_id) REFERENCES heartbeat_rithms (id);

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
       rd.name AS reproductive_discharge_name,
       ds.name AS discharge_state_name,
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
       scs.name AS scar_state_name
FROM appointments a
  JOIN users u ON a.doctor_id = u.id
  JOIN patients p ON a.patient_id = p.id
  LEFT JOIN receipt_kind rk ON a.receipt_kind_id = rk.id
  LEFT JOIN pelvis_states ps ON a.pelvis_state_id = ps.id
  LEFT JOIN fetal_bladder_states fbs ON a.fetal_bladder_state_id = fbs.id
  LEFT JOIN fetal_bladder_aligns fba ON a.fetal_bladder_align_id = fba.id
  LEFT JOIN fetal_bladder_previas fbp ON a.fetal_bladder_previa_id = fbp.id
  LEFT JOIN vagina_states vs ON a.vagina_state_id = vs.id
  LEFT JOIN devel_organs dor ON a.devel_organs_id = dor.id
  LEFT JOIN reproductive_discharges rd ON a.reproductive_discharge_id = rd.id
  LEFT JOIN discharge_states ds ON a.discharge_state_id = ds.id
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
  LEFT JOIN belly_states scs ON a.scar_state_id = scs.id;