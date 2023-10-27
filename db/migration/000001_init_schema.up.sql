CREATE TABLE "user" (
    "id" bigserial PRIMARY KEY,
    "username" varchar UNIQUE NOT NULL,
    "hashed_password" varchar NOT NULL,
    "full_name" varchar NOT NULL,
    "email" varchar UNIQUE NOT NULL,
    "update_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
    "password_changed_at" timestamptz NOT NULL DEFAULT('0001-01-01 00:00:00Z'),  
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "todo" (
    "id" bigserial PRIMARY KEY,
    "user_id" bigserial NOT NULL,
    "name" varchar NOT NULL,
    "description" varchar NOT NULL,
    "status" bool NOT NULL DEFAULT(false),
    "deadline" timestamptz NOT NULL,
    "update_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "todo" ("user_id");

ALTER TABLE "todo" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");