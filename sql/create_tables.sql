-- CREATE USER IF NOT EXISTS test;
-- CREATE DATABASE IF NOT EXISTS test;
-- GRANT ALL PRIVILEGES ON DATABASE test TO test;
CREATE TABLE "user" (
    id uuid NOT NULL PRIMARY KEY,
    first_name text,
    last_name text,
    nick_name text,
    password text,
    email text,
    country text,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);