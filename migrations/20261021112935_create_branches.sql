-- +goose Up
CREATE TABLE IF NOT EXISTS branches(
    "id" SERIAL PRIMARY KEY,
    "name" VARCHAR(50) NOT NULL,
    "status" VARCHAR(50) NOT NULL DEFAULT 'ACTIVE',
    "location" VARCHAR(200) NOT NULL,
    "created_at" TIMESTAMP DEFAULT NOW()
);

-- +goose Down
DROP TABLE IF EXISTS branches;

