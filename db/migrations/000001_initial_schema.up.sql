CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS accounts 
(
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4() NOT NULL,

    name varchar NOT NULL,
    email varchar NOT NULL,
    role varchar NOT NULL,
    avatar_url varchar,
    password varchar NOT NULL,

    examples_quantity integer NOT NULL,

    deleted_by uuid REFERENCES accounts (id),

    created_at timestamp DEFAULT now() NOT NULL,
    updated_at timestamp,
    deleted_at timestamp
);

CREATE TABLE IF NOT EXISTS examples 
(
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4() NOT NULL,
    
    name varchar NOT NULL,
    description varchar NOT NULL,
    
    members_quantity integer NOT NULL,

    owner_id uuid REFERENCES accounts (id),

    deleted_by uuid REFERENCES accounts (id),
    created_at timestamp DEFAULT now() NOT NULL,
    updated_at timestamp,
    deleted_at timestamp
);

CREATE TABLE IF NOT EXISTS example_members 
(
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4() NOT NULL,

    example_id uuid REFERENCES examples (id),
    account_id uuid REFERENCES accounts (id),

    deleted_by uuid REFERENCES accounts (id),
    created_at timestamp DEFAULT now() NOT NULL,
    updated_at timestamp,
    deleted_at timestamp
);
