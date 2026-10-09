-- +goose Up
CREATE TABLE IF NOT EXISTS users(
    "id" SERIAL PRIMARY KEY,
    "first_name" VARCHAR(50),
    "last_name" VARCHAR(50),
    "phone" VARCHAR(20),
    "status" VARCHAR(20) DEFAULT 'ACTIVE',
    "rating_score" DECIMAL(10, 2) NOT NULL,
    "role" VARCHAR(100) NOT NULL, -- SUPERADMIN, ADMIN, ADMISTRATOR, OPERATOR, TEACHER, ASSISTANT, ACCOUNTANT 
    "position" VARCHAR(100),  -- 'TEACHER', CEO, ASSISTANT, ACCOUNTANT
    "created_at" TIMESTAMP DEFAULT NOW(),
    "updated_at" TIMESTAMP DEFAULT NOW()
); 

-- +goose Down
DROP TABLE IF EXISTS users;
