CREATE TABLE IF NOT EXISTS users (
    id UInt64 DEFAULT generateUUIDv4(),  -- Генерация уникальных значений для id
    username String,
    email String,
    pswd_hash String,
    ver_hash String,
    timeout_at DateTime,
    verified_at DateTime DEFAULT NULL,
    created_at DateTime DEFAULT now(),
    updated_at DateTime DEFAULT now(),
    deleted_at DateTime DEFAULT NULL,
    PRIMARY KEY (id)  -- primary key на id
) ENGINE = MergeTree()
ORDER BY id;
