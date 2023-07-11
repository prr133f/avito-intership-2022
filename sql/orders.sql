CREATE SCHEMA IF NOT EXISTS orders_schema;

CREATE TABLE IF NOT EXISTS orders_schema.order (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    date DATE DEFAULT CURRENT_DATE,
    service_id UUID
);