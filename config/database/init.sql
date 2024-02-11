CREATE SCHEMA main;
CREATE SCHEMA kratos;


-- This is main application schema.
-- I do not wnat to mess arround with ORM's because the schema of this 
-- application is preaty simple, and I can use simple raw sql queries

-- Creating the 'issues' table
-- some reports or other errors will be stored here
CREATE TABLE main.issues (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title VARCHAR(100),
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- This table contain all user generated reports
-- I intentionaly does not provide default `id` since it should be provided 
-- outside of db.
CREATE TABLE main.reports (
    id UUID PRIMARY KEY,  
    user_id UUID NOT NULL,
    ticker VARCHAR(255) NOT NULL,
    link VARCHAR(255),
    email_send VARCHAR(255),
    email_recepient VARCHAR(255),
    date_created TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    success BOOLEAN,
    issue_id UUID REFERENCES main.issues(id)
);

-- Creating an index on the 'user_id' column
CREATE INDEX ON main.reports (user_id);

