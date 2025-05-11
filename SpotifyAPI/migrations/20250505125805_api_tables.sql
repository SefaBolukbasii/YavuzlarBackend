-- +goose Up
CREATE TABLE t_roles (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL
);

CREATE TABLE t_users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    username TEXT NOT NULL,
    password TEXT NOT NULL,
    roleId UUID REFERENCES t_roles(id),
    accountType TEXT DEFAULT 'Free',
    cash NUMERIC DEFAULT 100
);

INSERT INTO t_roles (name) VALUES 
  ('admin'),
  ('user');

-- +goose Down
DROP TABLE IF EXISTS t_users;
DROP TABLE IF EXISTS t_roles;


