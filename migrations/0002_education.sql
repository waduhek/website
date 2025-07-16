CREATE TABLE IF NOT EXISTS education (
    id serial,
    institute text NOT NULL,
    degree text NOT NULL,
    major text NOT NULL,
    grade text NOT NULL,
    location text NOT NULL,
    start_date date NOT NULL,
    end_date date NOT NULL
);

CREATE INDEX education_start_date_desc ON education (start_date DESC);
