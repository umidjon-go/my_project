-- +goose Up
CREATE TABLE IF NOT EXISTS lessons(
    "id" SERIAL PRIMARY KEY,
    "course_id" INT REFERENCES courses(id) ON DELETE CASCADE,
    "user_id" INT REFERENCES students(id) ON DELETE SET NULL,
    "group_id" INT REFERENCES groups(id) ON DELETE CASCADE,
    "title" VARCHAR(100),
    "description" TEXT,
    "status" VARCHAR(55) DEFAULT 'NEW',
    "created_at" TIMESTAMP DEFAULT NOW(),
    "completed_at" TIMESTAMP 
);

-- +goose Down
DROP TABLE IF EXISTS lessons;





