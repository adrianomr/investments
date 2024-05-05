
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

CREATE TABLE investments (
    investment_id UUID NOT NULL,
    investor_name TEXT,
    investment_amount NUMERIC(15, 2),
    investment_type TEXT NOT NULL
);
