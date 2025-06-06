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

CREATE TABLE t_songs (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    songerName TEXT NOT NULL,
    clickCount INT DEFAULT 0
);
CREATE TABLE t_playlists (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    userId UUID REFERENCES t_users(id)
);

CREATE TABLE t_playlist_songs (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    playlistId UUID REFERENCES t_playlists(id) ON DELETE CASCADE,
    songId UUID REFERENCES t_songs(id)
);
CREATE TABLE t_cupons (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    discount NUMERIC NOT NULL,
    userId UUID REFERENCES t_users(id)
    );

INSERT INTO t_roles (name) VALUES 
  ('admin'),
  ('user');

-- +goose Down
DROP TABLE IF EXISTS t_users;
DROP TABLE IF EXISTS t_roles;


