-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION "uuid-ossp";
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
