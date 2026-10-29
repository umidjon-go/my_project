-- +goose Up
CREATE TABLE IF NOT EXISTS courses(
    "id" SERIAL PRIMARY KEY,
    "user_id" INT REFERENCES users(id) ON DELETE SET NULL,
    "branch_id" INT REFERENCES branches(id) ON DELETE SET NULL,
    "title" VARCHAR(255),
    "description" TEXT,
    "price" BIGINT DEFAULT 0,
    "status" VARCHAR(55) DEFAULT 'ACTIVE',
    "created_at" TIMESTAMP DEFAULT NOW(),
    "updated_at" TIMESTAMP DEFAULT NOW()
);

-- +goose Down
DROP TABLE IF EXISTS courses;

