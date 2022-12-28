-- +goose Up
-- +goose StatementBegin
ALTER TABLE widget CHANGE COLUMN id ID INTEGER;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE widget CHANGE COLUMN ID id INTEGER;
-- +goose StatementEnd
