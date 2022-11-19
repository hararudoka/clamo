DROP TABLE IF EXISTS "user";
DROP TABLE IF EXISTS "message";

CREATE TABLE "user" (
    id uuid PRIMARY KEY,
    username VARCHAR NOT NULL CONSTRAINT name_unique UNIQUE,
    password VARCHAR NOT NULL
);

CREATE TABLE "message" (
    id uuid PRIMARY KEY,
    from_id INTEGER,
    to_id INTEGER,
    text TEXT
);