-- +goose Up
CREATE TABLE IF NOT EXISTS lesson_students(
    "lesson_id" INT REFERENCES lessons(id) ON DELETE CASCADE,
    "student_id" INT REFERENCES students(id) ON DELETE CASCADE,
    "status" VARCHAR(123)
);

-- +goose Down
DROP TABLE IF EXISTS lesson_students;


