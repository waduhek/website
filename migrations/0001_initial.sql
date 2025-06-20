CREATE TABLE IF NOT EXISTS experience (
    id serial,
    title text NOT NULL,
    company_name text NOT NULL,
    start_date date NOT NULL,
    end_date date NOT NULL DEFAULT '0001-01-01',
    is_current boolean NOT NULL,
    location text NOT NULL,
    description text[] NOT NULL,
    skills text[] NOT NULL
);

CREATE INDEX experience_start_date_desc ON experience (start_date DESC);
