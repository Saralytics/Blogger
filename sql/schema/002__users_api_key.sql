-- +goose Up
ALTER TABLE users 
ADD api_key varchar(64) not null default encode(sha256(random()::text::bytea), 'hex');

-- +goose Down
ALTER TABLE users DROP api_key;