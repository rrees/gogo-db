-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS widget (
	id INTEGER AUTO_INCREMENT PRIMARY KEY,
	name TEXT NOT NULL,
	description TEXT
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS widget;
-- +goose StatementEnd
