CREATE SCHEMA IF NOT EXISTS services_schema;

CREATE TABLE IF NOT EXISTS services_schema.service (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT,
    cost FLOAT
);