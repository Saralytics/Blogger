-- +goose Up
CREATE TABLE
    users (
        id uuid primary key gen_random_uuid (),
        created_at timestamp not null,
        updated_at timestamp not null,
        name string not null
    );

-- +goose Down
DROP TABLE users;