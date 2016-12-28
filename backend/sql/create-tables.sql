CREATE TABLE IF NOT EXISTS users (
    "id"    serial primary key,
    "login" varchar(100) NOT NULL,
    "name"  varchar(100) NOT NULL,
    "psw_hash" varchar(100) NOT NULL
);

CREATE INDEX ind_users_id ON users USING btree (id);
CREATE UNIQUE INDEX uni_users ON users (lower(login));
CREATE INDEX ind_users_psw ON users USING btree (login, psw_hash);
