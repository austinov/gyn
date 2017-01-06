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

CREATE TABLE IF NOT EXISTS health_states (
    "id"      serial primary key,
    "name"    varchar(100) NOT NULL,
    "orderby" integer
);
COMMENT ON TABLE health_states IS 'Типы состояния пациента';
CREATE INDEX ind_health_states_id ON health_states USING btree (id);
CREATE UNIQUE INDEX uni_health_states ON health_states (lower(name));

INSERT INTO health_states (name) VALUES ('относительно удовлетворительное'), ('удовлетворительное'), ('средней тяжести'), ('тяжелое');

CREATE TABLE IF NOT EXISTS skin_states (
    "id"      serial primary key,
    "name"    varchar(100) NOT NULL,
    "orderby" integer
);
COMMENT ON TABLE skin_states IS 'Типы состояния кожных покровов';
CREATE INDEX ind_skin_states_id ON skin_states USING btree (id);
CREATE UNIQUE INDEX uni_skin_states ON skin_states (lower(name));

INSERT INTO skin_states (name) VALUES ('бледно-розовой окраски'), ('бледные');

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

INSERT INTO fetal_heartbeats (name) VALUES ('ясное'), ('ритмичное'), ('приглушенное'), ('глухое');

CREATE TABLE IF NOT EXISTS reproductive_discharges (
    "id"      serial primary key,
    "name"    varchar(100) NOT NULL,
    "orderby" integer
);
COMMENT ON TABLE reproductive_discharges IS 'Типы выделений из половых путей';
CREATE INDEX ind_reproductive_discharges_id ON reproductive_discharges USING btree (id);
CREATE UNIQUE INDEX uni_reproductive_discharges ON reproductive_discharges (lower(name));

INSERT INTO reproductive_discharges (name) VALUES ('светлые'), ('слизистые'), ('кровянистые'), ('умеренные'), ('обильные');

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

CREATE TABLE IF NOT EXISTS outer_throat_states (
    "id"      serial primary key,
    "name"    varchar(100) NOT NULL,
    "orderby" integer
);
COMMENT ON TABLE outer_throat_states IS 'Типы наружнего зева';
CREATE INDEX ind_outer_throat_states_id ON outer_throat_states USING btree (id);
CREATE UNIQUE INDEX uni_outer_throat_states ON outer_throat_states (lower(name));

INSERT INTO outer_throat_states (name) VALUES ('пропускает кончик пальца'), ('закрыт (0 баллов)');

CREATE TABLE IF NOT EXISTS fetal_bladder_previas (
    "id"      serial primary key,
    "name"    varchar(100) NOT NULL,
    "orderby" integer
);
COMMENT ON TABLE fetal_bladder_previas IS 'Типы предлежания плодный пузырья';
CREATE INDEX ind_fetal_bladder_previas_id ON fetal_bladder_previas USING btree (id);
CREATE UNIQUE INDEX uni_fetal_bladder_previas ON fetal_bladder_previas (lower(name));

INSERT INTO fetal_bladder_previas (name) VALUES ('головка'), ('ягодицы'), ('ножки');

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
    "how_receipt"               varchar(100),
    "alergo"                    varchar(100),
    "contact_infectied"         varchar(100),
    "hiv"                       varchar(100),
    "transfusion"               varchar(100),
    "dyscountry"                varchar(100),
    "smoking"                   varchar(100),
    "drugs"                     varchar(100),
    "inheritance"               varchar(100),
    "diseases"                  varchar(100),
    "gyndiseases"               varchar(100),
    "paritet"                   varchar(500),
    "pregnancy"                 varchar(500),
    "first_trimester"           varchar(500),
    "second_trimester"          varchar(500),
    "third_trimester"           varchar(500),
    "history"                   text,
    "exp_by_menstruation"       varchar(100),
    "exp_by_first_visit"        varchar(100),
    "exp_by_ultra"              varchar(100),
    "health_state_id"           integer,
    "claims"                    varchar(100),
    "head"                      varchar(100),
    "vision"                    varchar(100),
    "skin_state_id"             integer,
    "lymph"                     varchar(100),
    "breath"                    varchar(100),
    "rale"                      varchar(100),
    "tones"                     varchar(100),
    "pulse"                     varchar(100),
    "pulse_type"                varchar(100),
    "pressure"                  varchar(100),
    "tongue_clean"              boolean,
    "tongue_wet"                boolean,
    "tongue_dry"                boolean,
    "tongue_coated"             boolean,
    "tongue_uncoated"           boolean,
    "throat"                    varchar(100),
    "belly"                     varchar(100),
    "peritoneal"                varchar(100),
    "labors"                    varchar(100),
    "dysuric"                   boolean,
    "bowel"                     boolean,
    "limb_swelling"             varchar(100),
    "face_swelling"             varchar(100),
    "uteruse_state_id"          integer,
    "fetal_position_id"         integer,
    "fetal_previa_id"           integer,
    "fetal_align_id"            integer,
    "fetal_heartbeat_id"        integer,
    "fetal_pulse"               varchar(100),
    "reproductive_discharge_id" integer,
    "vdm"                       varchar(20),
    "oj"                        varchar(20),
    "dspin"                     varchar(20),
    "dcrist"                    varchar(20),
    "dtroch"                    varchar(20),
    "cext"                      varchar(20),
    "devel_organs_id"           integer,
    "genital_anomalies"         varchar(100),
    "vagina_state_id"           integer,
    "lenght_cervix"             varchar(50),
    "truncate_cervix"           varchar(50),
    "outer_throat_state_id"     integer,
    "channel_cervix"            varchar(50),
    "fetal_bladder"             varchar(50),
    "fetal_bladder_previa_id"   integer,
    "fetal_bladder_align_id"    integer,
    "arches"                    varchar(100),
    "conjugate"                 varchar(50),
    "pelvis_state_id"           integer,
    "pelvis_discharge"          varchar(100),
    "diagnosis"                 text,
    "conclusion"                text,
    "birth_plan"                text
);

COMMENT ON TABLE appointments IS 'Осмотры пациентов';
COMMENT ON COLUMN appointments.date_receipt IS 'Дата приёма, Unix timestamp in seconds';
COMMENT ON COLUMN appointments.doctor_id IS 'Дежурный врач акушер-гинеколог';
COMMENT ON COLUMN appointments.patient_id IS 'Пациент';
COMMENT ON COLUMN appointments.how_receipt IS 'Поступила';
COMMENT ON COLUMN appointments.alergo IS 'Аллергоанамнез';
COMMENT ON COLUMN appointments.contact_infectied IS 'Контакт с инфекционными больными';
COMMENT ON COLUMN appointments.hiv IS 'Малярия, туберкулез, гепатиты, ВИЧ';
COMMENT ON COLUMN appointments.transfusion IS 'Гематрансфузия';
COMMENT ON COLUMN appointments.dyscountry IS 'Пребывание в неблагополучных странах в течении 3-х лет';
COMMENT ON COLUMN appointments.smoking IS 'Табакокурение';
COMMENT ON COLUMN appointments.drugs IS 'Прием наркотических, психотропных препаратов';
COMMENT ON COLUMN appointments.inheritance IS 'Наследственность';
COMMENT ON COLUMN appointments.diseases IS 'Перенесенные заболевания';
COMMENT ON COLUMN appointments.gyndiseases IS 'Гинекологические заболевания';
COMMENT ON COLUMN appointments.paritet IS 'Паритет';
COMMENT ON COLUMN appointments.pregnancy IS 'Течение беременности';
COMMENT ON COLUMN appointments.first_trimester IS 'I триместр';
COMMENT ON COLUMN appointments.second_trimester IS 'II триместр';
COMMENT ON COLUMN appointments.third_trimester IS 'III триместр';
COMMENT ON COLUMN appointments.history IS 'Из анамнеза';
COMMENT ON COLUMN appointments.exp_by_menstruation IS 'Сроки беременности по менструации';
COMMENT ON COLUMN appointments.exp_by_first_visit IS 'Сроки беременности по 1 явке';
COMMENT ON COLUMN appointments.exp_by_ultra IS 'Сроки беременности по УЗИ';
COMMENT ON COLUMN appointments.health_state_id IS 'Состояние';
COMMENT ON COLUMN appointments.claims IS 'Жалобы';
COMMENT ON COLUMN appointments.head IS 'Голова';
COMMENT ON COLUMN appointments.vision IS 'Зрение';
COMMENT ON COLUMN appointments.skin_state_id IS 'Кожные покровы';
COMMENT ON COLUMN appointments.lymph IS 'Лимфоузлы';
COMMENT ON COLUMN appointments.breath IS 'Дыхание';
COMMENT ON COLUMN appointments.rale IS 'Хрипы';
COMMENT ON COLUMN appointments.tones IS 'Тоны cor';
COMMENT ON COLUMN appointments.pulse IS 'Пульс';
COMMENT ON COLUMN appointments.pulse_type IS 'Ритмичность пульса';
COMMENT ON COLUMN appointments.pressure IS 'АД';
COMMENT ON COLUMN appointments.tongue_clean IS 'Язык чистый';
COMMENT ON COLUMN appointments.tongue_wet IS 'Язык влажный';
COMMENT ON COLUMN appointments.tongue_dry IS 'Язык сухой';
COMMENT ON COLUMN appointments.tongue_coated IS 'Язык обложен';
COMMENT ON COLUMN appointments.tongue_uncoated IS 'Язык не обложен';
COMMENT ON COLUMN appointments.throat IS 'Осмотр зева';
COMMENT ON COLUMN appointments.belly IS 'Живот';
COMMENT ON COLUMN appointments.peritoneal IS 'Перитонеальные симптомы';
COMMENT ON COLUMN appointments.labors IS 'Родовая деятельность';
COMMENT ON COLUMN appointments.dysuric IS 'Дизурические явления (нет/есть)';
COMMENT ON COLUMN appointments.bowel IS 'Стул (не регулярный/регулярный)';
COMMENT ON COLUMN appointments.limb_swelling IS 'Отеки';
COMMENT ON COLUMN appointments.face_swelling IS 'Отеки лица';
COMMENT ON COLUMN appointments.uteruse_state_id IS 'Матка';
COMMENT ON COLUMN appointments.fetal_position_id IS 'Положение плода';
COMMENT ON COLUMN appointments.fetal_previa_id IS 'Предлежит плода';
COMMENT ON COLUMN appointments.fetal_align_id IS 'Выравнивание плода';
COMMENT ON COLUMN appointments.fetal_heartbeat_id IS 'Сердцебиение плода';
COMMENT ON COLUMN appointments.fetal_pulse IS 'Пульс плода';
COMMENT ON COLUMN appointments.reproductive_discharge_id IS 'Выделения из половых путей';
COMMENT ON COLUMN appointments.vdm IS 'ВДМ';
COMMENT ON COLUMN appointments.oj IS 'ОЖ';
COMMENT ON COLUMN appointments.dspin IS 'D.spin';
COMMENT ON COLUMN appointments.dcrist IS 'D.crist';
COMMENT ON COLUMN appointments.dtroch IS 'D.troch';
COMMENT ON COLUMN appointments.cext IS 'C.ext';
COMMENT ON COLUMN appointments.devel_organs_id IS 'Наружние половые органы развиты';
COMMENT ON COLUMN appointments.genital_anomalies IS 'Аномалии';
COMMENT ON COLUMN appointments.vagina_state_id IS 'Влагалище';
COMMENT ON COLUMN appointments.lenght_cervix IS 'Шейка матки длиной';
COMMENT ON COLUMN appointments.truncate_cervix IS 'Шейка матки укорочена';
COMMENT ON COLUMN appointments.outer_throat_state_id IS 'Наружний зев';
COMMENT ON COLUMN appointments.channel_cervix IS 'Цервикальный канал';
COMMENT ON COLUMN appointments.fetal_bladder IS 'Плодный пузырь';
COMMENT ON COLUMN appointments.fetal_bladder_previa_id IS 'Плодный пузырь предлежит';
COMMENT ON COLUMN appointments.fetal_bladder_align_id IS 'Плодный пузырь прижат';
COMMENT ON COLUMN appointments.arches IS 'Своды';
COMMENT ON COLUMN appointments.conjugate IS 'Диагональная коньюгата';
COMMENT ON COLUMN appointments.pelvis_state_id IS 'Костный таз';
COMMENT ON COLUMN appointments.pelvis_discharge IS 'Выделения (костный таз)';
COMMENT ON COLUMN appointments.diagnosis IS 'Диагноз';
COMMENT ON COLUMN appointments.conclusion IS 'Заключение';
COMMENT ON COLUMN appointments.birth_plan IS 'План родов';

CREATE INDEX ind_appointments_id ON appointments USING btree (id);
CREATE INDEX ind_appointments_date ON appointments USING btree (date_receipt);
CREATE INDEX ind_appointments_doctor_id ON appointments USING btree (doctor_id);
CREATE INDEX ind_appointments_patient_id ON appointments USING btree (patient_id);
ALTER TABLE appointments ADD CONSTRAINT fk_appointments_doctor FOREIGN KEY (doctor_id) REFERENCES users (id);
ALTER TABLE appointments ADD CONSTRAINT fk_appointments_patient FOREIGN KEY (patient_id) REFERENCES patients (id);
ALTER TABLE appointments ADD CONSTRAINT fk_appointments_pelvis_states FOREIGN KEY (pelvis_state_id) REFERENCES pelvis_states (id);
ALTER TABLE appointments ADD CONSTRAINT fk_appointments_fetal_bladder_aligns FOREIGN KEY (fetal_bladder_align_id) REFERENCES fetal_bladder_aligns (id);
ALTER TABLE appointments ADD CONSTRAINT fk_appointments_fetal_bladder_previas FOREIGN KEY (fetal_bladder_previa_id) REFERENCES fetal_bladder_previas (id);
ALTER TABLE appointments ADD CONSTRAINT fk_appointments_outer_throat_states FOREIGN KEY (outer_throat_state_id) REFERENCES outer_throat_states (id);
ALTER TABLE appointments ADD CONSTRAINT fk_appointments_vagina_states FOREIGN KEY (vagina_state_id) REFERENCES vagina_states (id);
ALTER TABLE appointments ADD CONSTRAINT fk_appointments_devel_organs FOREIGN KEY (devel_organs_id) REFERENCES devel_organs (id);
ALTER TABLE appointments ADD CONSTRAINT fk_appointments_reproductive_discharges FOREIGN KEY (reproductive_discharge_id) REFERENCES reproductive_discharges (id);
ALTER TABLE appointments ADD CONSTRAINT fk_appointments_fetal_aligns FOREIGN KEY (fetal_align_id) REFERENCES fetal_aligns (id);
ALTER TABLE appointments ADD CONSTRAINT fk_appointments_fetal_heartbeats FOREIGN KEY (fetal_heartbeat_id) REFERENCES fetal_heartbeats (id);
ALTER TABLE appointments ADD CONSTRAINT fk_appointments_fetal_previas FOREIGN KEY (fetal_previa_id) REFERENCES fetal_previas (id);
ALTER TABLE appointments ADD CONSTRAINT fk_appointments_fetal_positions FOREIGN KEY (fetal_position_id) REFERENCES fetal_positions (id);
ALTER TABLE appointments ADD CONSTRAINT fk_appointments_uteruse_states FOREIGN KEY (uteruse_state_id) REFERENCES uteruse_states (id);
ALTER TABLE appointments ADD CONSTRAINT fk_appointments_skin_states FOREIGN KEY (skin_state_id) REFERENCES skin_states (id);
ALTER TABLE appointments ADD CONSTRAINT fk_appointments_health_states FOREIGN KEY (health_state_id) REFERENCES health_states (id);