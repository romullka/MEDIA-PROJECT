CREATE TABLE products (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT,
    specs JSONB,
    weight FLOAT,
    barcode TEXT
);