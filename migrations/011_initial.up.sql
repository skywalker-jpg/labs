BEGIN;

CREATE TABLE IF NOT EXISTS jokes (
                       id SERIAL PRIMARY KEY,
                       joke_id VARCHAR(255) UNIQUE,
                       category VARCHAR(100),
                       joke TEXT,
                       created_at TIMESTAMP DEFAULT NOW(),
);

COMMIT;