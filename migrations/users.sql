CREATE SCHEMA IF NOT EXISTS goarch;

CREATE TABLE IF NOT EXISTS goarch.users (
  id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    username character varying(255) NOT NULL,
    password TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE,
    deleted_at TIMESTAMP WITH TIME ZONE
);

