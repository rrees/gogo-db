-- +goose Up
-- +goose StatementBegin
ALTER TABLE widget RENAME widgets;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE widgets RENAME widget;
-- +goose StatementEnd
