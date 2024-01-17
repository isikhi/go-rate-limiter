-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS rate_limit_options
(
    id bigserial,
    client_id varchar(255) not null,
    token_count integer not null,
    duration integer not null,
    created_at timestamp with time zone default current_timestamp,
    primary key (id)
);

CREATE OR REPLACE FUNCTION update_updated_at_column()
    RETURNS TRIGGER AS $$
BEGIN
    RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_rate_limit_options_updated_at BEFORE UPDATE
    ON rate_limit_options FOR EACH ROW EXECUTE PROCEDURE
    update_updated_at_column();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table rate_limit_options;
-- +goose StatementEnd
