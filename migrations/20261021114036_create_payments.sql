-- +goose Up
CREATE TABLE IF NOT EXISTS payments(
    "id" SERIAL PRIMARY KEY,
    "course_id" INT REFERENCES courses(id) ON DELETE SET NULL,
    "student_id" INT REFERENCES students(id) ON DELETE CASCADE,
    "group_id" INT REFERENCES groups(id) ON DELETE SET NULL,
    "user_id" INT REFERENCES users(id) ON DELETE SET NULL,
    "amount" BIGINT,
    "payment_type" VARCHAR(100),
    "status" VARCHAR(100),
    "paid_at" TIMESTAMP DEFAULT NOW(),
    "comment" VARCHAR(250),
    "created_at" TIMESTAMP DEFAULT NOW(),
    "updated_at" TIMESTAMP DEFAULT NOW()
);

-- +goose Down
DROP TABLE IF EXISTS payments;



