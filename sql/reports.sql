CREATE SCHEMA IF NOT EXISTS reports_schema;

CREATE TABLE IF NOT EXISTS reports_schema.report (
    date DATE,
    service_id UUID,
    user_id UUID,
);