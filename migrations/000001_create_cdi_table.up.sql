-------------------------------------------------------------------------
-- V0__extension_and_update_trigger
-------------------------------------------------------------------------

CREATE
EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE
OR REPLACE FUNCTION trigger_set_updated_at() RETURNS TRIGGER AS
$$
BEGIN
    NEW.updated_at
= NOW();
RETURN NEW;
END;
$$
LANGUAGE plpgsql;

-------------------------------------------------------------------------
-- V1__initial_schema
-------------------------------------------------------------------------

DO
$$
BEGIN

END;
$$
LANGUAGE plpgsql;

CREATE TABLE cdb
(
    id              UUID           NOT NULL PRIMARY KEY,
    user_id         TEXT           NOT NULL,
    amount          NUMERIC(15, 2) NOT NULL DEFAULT 0,
    percentage      NUMERIC(15, 2) NOT NULL,
    investment_type TEXT           NOT NULL
);


CREATE TABLE cdb_order
(
    id         UUID           NOT NULL PRIMARY KEY,
    user_id    TEXT           NOT NULL,
    amount     NUMERIC(15, 2) NOT NULL DEFAULT 0,
    type TEXT           NOT NULL,
    cdb_id     UUID           NOT NULL
);

ALTER TABLE cdb_order
    ADD CONSTRAINT cdb_order_cdb_id_fkey FOREIGN KEY (cdb_id) REFERENCES cdb (id);
