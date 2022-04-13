-- Add UUID extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Set timezone
-- For more information, please visit:
-- https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
SET TIMEZONE="Europe/Moscow";

-- Create users table
CREATE TABLE users (
    id UUID DEFAULT uuid_generate_v4 () PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW (),
    updated_at TIMESTAMP NULL,
    email VARCHAR (255) NOT NULL UNIQUE,
    password_hash VARCHAR (255) NOT NULL,
    user_status INT NOT NULL,
    user_role VARCHAR (25) NOT NULL
);

-- Create compute table
CREATE TABLE computes (
    id BIGINT PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW (),
    updated_at TIMESTAMP NULL,
    name VARCHAR (255) NOT NULL,
    json_path VARCHAR NOT NULL,
    project_id VARCHAR NOT NULL,
    zone VARCHAR NOT NULL
);

-- Create instances table
CREATE TABLE instances (
    id BIGINT PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW (),
    updated_at TIMESTAMP NULL,
    name VARCHAR (255) NOT NULL,
    zone VARCHAR (255) NOT NULL,
    status VARCHAR NOT NULL,
    compute_id BIGINT NOT NULL REFERENCES computes (id) ON DELETE CASCADE
);

-- Add indexes
CREATE INDEX active_users ON users (id) WHERE user_status = 1;
