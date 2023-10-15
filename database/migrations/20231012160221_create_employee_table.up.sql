CREATE TABLE IF NOT EXISTS employees (
     id SERIAL PRIMARY KEY,
     first_name TEXT NOT NULL,
     last_name TEXT NOT NULL,
     email TEXT NOT NULL UNIQUE,
     hire_date DATE NOT NULL,
     deleted_at TIMESTAMPTZ,
     created_at TIMESTAMPTZ DEFAULT NOW(),
     updated_at TIMESTAMPTZ DEFAULT NOW()
);
