-- +goose Up
CREATE TABLE IF NOT EXISTS students(
    "id" SERIAL PRIMARY KEY,
    "course_id" INT REFERENCES courses(id) ON DELETE SET NULL,
    "user_id" INT REFERENCES users(id) ON DELETE SET NULL,
    "group_id" INT REFERENCES groups(id) ON DELETE SET NULL,
    "full_name" VARCHAR(50),
    "phone" VARCHAR(20),
    "parent_phone" VARCHAR(20),
    "status" VARCHAR DEFAULT 'NEW',
    "created_at" TIMESTAMP DEFAULT NOW(),
    "updated_at" TIMESTAMP DEFAULT NOW()
);

-- +goose Down
DROP TABLE IF EXISTS students;



