-- +goose Up
CREATE TABLE IF NOT EXISTS groups(
    "id" SERIAL PRIMARY KEY,
    "course_id" INT REFERENCES courses(id) ON DELETE CASCADE,
    "user_id" INT REFERENCES users(id) ON DELETE SET NULL,
    "name" VARCHAR(150),
    "status" VARCHAR(50),
    "days" VARCHAR(30),
    "start_time" VARCHAR(100),
    "end_time" VARCHAR(100),
    "room" VARCHAR(100),
    "start_date" VARCHAR(100),
    "end_date" VARCHAR(100),
    "created_at" TIMESTAMP DEFAULT NOW(),
    "updated_at" TIMESTAMP DEFAULT NOW()
);
-- +goose Down
DROP TABLE IF EXISTS groups;


