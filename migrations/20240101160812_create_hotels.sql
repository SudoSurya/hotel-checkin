-- +goose Up
-- +goose StatementBegin
CREATE TABLE hotels (
    id SERIAL PRIMARY KEY,
    hotel_name VARCHAR(255) NOT NULL,
    city VARCHAR(255) NOT NULL,
    state VARCHAR(255) NOT NULL,
    country VARCHAR(255) NOT NULL,
    zip VARCHAR(255) NOT NULL,
    landline VARCHAR(255) NOT NULL,
    owner_name VARCHAR(255) NOT NULL,
    owner_email VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE hotels;
-- +goose StatementEnd
