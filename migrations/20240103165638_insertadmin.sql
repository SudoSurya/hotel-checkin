-- +goose Up
-- +goose StatementBegin
INSERT INTO admin (id,name, email, password, department, isAuthorized, created_at, updated_at) VALUES
                  ('32323121','admin', 'admin@example.com', 'admin', 'FBI', true, '2020-01-03 16:56:38', '2020-01-03 16:56:38');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM admin WHERE email = 'admin@example.com';
-- +goose StatementEnd
