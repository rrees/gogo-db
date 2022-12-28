-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS locations (
	ID INTEGER PRIMARY KEY AUTO_INCREMENT NOT NULL,
	name TEXT NOT NULL
);
-- +goose StatementEnd
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS inventory (
	ID INTEGER PRIMARY KEY AUTO_INCREMENT NOT NULL,
	location_id INTEGER NOT NULL,
	widget_id INTEGER NOT NULL,
	quantity INTEGER NOT NULL DEFAULT 0,


	INDEX location_widget_idx (location_id, widget_id),

	CONSTRAINT fk_inv_locations FOREIGN KEY (location_id)
		REFERENCES locations(ID),
	CONSTRAINT fk_inv_widgets FOREIGN KEY (widget_id)
		REFERENCES widgets(ID)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE inventory;
-- +goose StatementEnd
-- +goose StatementBegin
DROP TABLE locations;
-- +goose StatementEnd