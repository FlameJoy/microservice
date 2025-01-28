CREATE TABLE IF NOT EXISTS orders (
    id UInt64 DEFAULT generateUUIDv4(),  -- Генерация уникальных значений для id
    created_at DateTime DEFAULT now(),
    updated_at DateTime DEFAULT now(),
    deleted_at DateTime DEFAULT NULL,
    PRIMARY KEY (id)  -- primary key на id
) ENGINE = MergeTree()
ORDER BY id;