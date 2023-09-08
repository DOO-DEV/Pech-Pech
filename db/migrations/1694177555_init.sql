-- +migrate Up
CREATE TABLE IF NOT EXISTS users (
    "id"            TEXT NOT NULL,
    "username"      TEXT NOT NULL,
    "password"      TEXT NOT NULL,
    "room_limit"    TEXT NOT NULL,
    CONSTRAINT "sers_pkey" PRIMARY KEY ("id"),
    CONSTRAINT "users_username_key" UNIQUE ("username")
);

-- +migrate Down
DROP TABLE IF EXISTS "users";