CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    pswd_hash VARCHAR(100) NOT NULL,
    ver_hash VARCHAR(100) NOT NULL,
    timeout_at TIMESTAMP NOT NULL,
    verified_at TIMESTAMP DEFAULT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);

-- Index
-- CREATE INDEX idx_users_deleted_at ON users (deleted_at);
CREATE INDEX idx_users_deleted_at_null ON users (id) WHERE deleted_at IS NULL;
-- CREATE INDEX idx_users_username ON users (username);
-- CREATE INDEX idx_users_deleted_email ON users (deleted_at, email);
-- CREATE INDEX idx_users_updated_at ON users (updated_at DESC);