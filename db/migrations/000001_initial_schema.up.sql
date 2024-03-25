CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS accounts 
(
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4() NOT NULL,

    name varchar NOT NULL,
    email varchar NOT NULL,
    role varchar NOT NULL,
    avatar_url varchar,
    password varchar NOT NULL,

    deleted_by uuid REFERENCES accounts (id),

    created_at timestamp DEFAULT now() NOT NULL,
    updated_at timestamp,
    deleted_at timestamp
);