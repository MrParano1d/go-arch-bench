CREATE SCHEMA IF NOT EXISTS goarch;

CREATE TABLE IF NOT EXISTS goarch.books (
    id serial PRIMARY KEY,
    title character varying(255) NOT NULL,
    author_id uuid NOT NULL REFERENCES goarch.users(id),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE,
    deleted_at TIMESTAMP WITH TIME ZONE
);
