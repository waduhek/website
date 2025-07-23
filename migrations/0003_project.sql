CREATE TABLE IF NOT EXISTS project (
    id integer PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    name text NOT NULL,
    public_url text NOT NULL DEFAULT '',
    repo_url text NOT NULL DEFAULT '',
    description text[] NOT NULL,
    technologies text[] NOT NULL
);
