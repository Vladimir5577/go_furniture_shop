-- +goose Up
-- +goose StatementBegin
CREATE TABLE category (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) UNIQUE NOT NULL,
    description TEXT,
    image VARCHAR(255),
    is_active BOOlEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP default current_timestamp
);

CREATE TABLE furniture (
    id SERIAL PRIMARY KEY,
    category_id INT,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    price numeric CHECK (price > 0),
    image VARCHAR(255),
    is_active BOOlEAN DEFAULT TRUE,
    created_at TIMESTAMP default current_timestamp,
    updated_at TIMESTAMP default current_timestamp,

    FOREIGN KEY (category_id) REFERENCES category(id)
        ON DELETE SET NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS furniture;
DROP TABLE IF EXISTS category;
-- +goose StatementEnd
