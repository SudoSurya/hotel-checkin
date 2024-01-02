-- +goose Up
-- +goose StatementBegin
CREATE TABLE feds (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    department VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(64) NOT NULL,
    status VARCHAR(64) DEFAULT 'progress' check (status in ('progress', 'accepted', 'rejected')),
    number NUMBER NOT NULL,
    isAuthorized BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE feds;
-- +goose StatementEnd

