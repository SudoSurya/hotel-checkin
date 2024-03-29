-- +goose Up
-- +goose StatementBegin
CREATE TABLE hotels (
    id SERIAL PRIMARY KEY,
    hotel_name VARCHAR(255) NOT NULL,
    city VARCHAR(255) NOT NULL,
    state VARCHAR(255) NOT NULL,
    country VARCHAR(255) NOT NULL,
    zip NUMBER NOT NULL,
    landline NUMBER NOT NULL,
    owner_name VARCHAR(255) NOT NULL,
    owner_email VARCHAR(255) NOT NULL,
    password VARCHAR(64) NOT NULL,
    status VARCHAR(64) DEFAULT 'progress' check (status in ('progress', 'accepted', 'rejected')),
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE hotels;
-- +goose StatementEnd
