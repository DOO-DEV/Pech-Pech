-- +migrate Up
CREATE TABLE IF NOT EXISTS "rooms" (
  "name" TEXT NOT NULL,
  "description" TEXT,
  "create_by" TEXT,
  "category" TEXT NOT NULL,
  "created_at" TIMESTAMP DEFAULT NOW() NOT NULL,
  "updated_at" TIMESTAMP DEFAULT NOW() NOT NULL,
  PRIMARY KEY("create_by", "name"),
  CONSTRAINT "fk_user" FOREIGN KEY("create_by") REFERENCES users("id") ON DELETE CASCADE
);

-- +migrate Down
DROP TABLE IF EXISTS "rooms";
