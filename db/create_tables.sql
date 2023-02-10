CREATE TABLE postgres (
    id uuid NOT NULLPRIMARY KEY,
    first_name text,
    last_name text,
    nick_name text,
    password text,
    email text,
    country text,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);