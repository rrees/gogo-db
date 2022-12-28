-- +goose Up
-- +goose StatementBegin
ALTER TABLE widgets CHANGE COLUMN ID ID INTEGER AUTO_INCREMENT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE widgets CHANGE COLUMN ID ID INTEGER;
-- +goose StatementEnd
