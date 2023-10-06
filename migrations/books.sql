CREATE TABLE IF NOT EXISTS books (
    id serial PRIMARY KEY,
    title character varying(255) NOT NULL,
    author_id uuid NOT NULL REFERENCES users(id),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE,
    deleted_at TIMESTAMP WITH TIME ZONE
);
