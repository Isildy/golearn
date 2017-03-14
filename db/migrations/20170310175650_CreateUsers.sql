-- +goose Up
CREATE TABLE user (
    id integer,
    first_name text,
    PRIMARY KEY(id)
);


-- +goose Down
DROP TABLE user;


