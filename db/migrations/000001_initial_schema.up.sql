CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS accounts 
(
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4() NOT NULL,

    name varchar NOT NULL,
    email varchar NOT NULL,
    role varchar NOT NULL,
    avatar_url varchar,
    password varchar NOT NULL,
    aggregates_quantity integer NOT NULL,

    examples_quantity integer NOT NULL,

    deleted_by uuid REFERENCES accounts (id),

    created_at timestamp DEFAULT now() NOT NULL,
    updated_at timestamp,
    deleted_at timestamp
);

CREATE TABLE IF NOT EXISTS aggregates 
(
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4() NOT NULL,
    
    name varchar NOT NULL,
    description varchar NOT NULL,
    
    members_quantity integer NOT NULL,
    examples_quantity integer NOT NULL,

    owner_id uuid REFERENCES accounts (id),

    deleted_by uuid REFERENCES accounts (id),
    created_at timestamp DEFAULT now() NOT NULL,
    updated_at timestamp,
    deleted_at timestamp
);

CREATE TABLE IF NOT EXISTS aggregate_members 
(
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4() NOT NULL,

    examples_quantity integer NOT NULL,

    aggregate_id uuid REFERENCES aggregates (id),
    account_id uuid REFERENCES accounts (id),

    deleted_by uuid REFERENCES accounts (id),
    created_at timestamp DEFAULT now() NOT NULL,
    updated_at timestamp,
    deleted_at timestamp
);

CREATE TABLE IF NOT EXISTS examples 
(
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4() NOT NULL,

    content varchar NOT NULL,    

    aggregate_id uuid REFERENCES aggregates (id),
    member_id uuid REFERENCES aggregate_members (id),

    deleted_by uuid REFERENCES accounts (id),
    created_at timestamp DEFAULT now() NOT NULL,
    updated_at timestamp,
    deleted_at timestamp
);

CREATE TABLE IF NOT EXISTS notifications 
(
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4() NOT NULL,

    subject varchar NOT NULL,
    message varchar NOT NULL,

    recipient_id uuid REFERENCES accounts (id),
  
    sent_at timestamp DEFAULT now() NOT NULL
)