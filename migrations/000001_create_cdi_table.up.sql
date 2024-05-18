
-------------------------------------------------------------------------
-- V0__extension_and_update_trigger
-------------------------------------------------------------------------

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE OR REPLACE FUNCTION trigger_set_updated_at() RETURNS TRIGGER AS
$$
BEGIN
    NEW.updated_at = NOW();
RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-------------------------------------------------------------------------
-- V1__initial_schema
-------------------------------------------------------------------------

DO $$
BEGIN

END;
$$ LANGUAGE plpgsql;

CREATE TABLE cdb (
    id UUID NOT NULL,
    user_id TEXT,
    amount NUMERIC(15, 2),
    cdi_percentage NUMERIC(15, 2),
    investment_type TEXT NOT NULL
);

CREATE TABLE cdb_order (
    id UUID NOT NULL,
    user_id TEXT,
    amount NUMERIC(15, 2),
    order_type TEXT NOT NULL,
    cdb_id UUID NOT NULL
);


