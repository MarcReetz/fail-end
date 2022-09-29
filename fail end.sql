CREATE TABLE security.user (
  "id" SERIAL PRIMARY KEY,
  "username" text UNIQUE,
  "password" text,
  "email" text UNIQUE
);

CREATE TABLE security.fail (
  "id" SERIAL UNIQUE,
  "title" text NOT NULL,
  "description" text,
  "user_id" int,
  "hits" int DEFAULT 0,
  "tags" int[],
  PRIMARY KEY ("user_id", "title")
);

CREATE TABLE security.tag (
  "id" SERIAL UNIQUE,
  "title" text NOT NULL,
  "user_id" int,
  "type" text[],
  PRIMARY KEY ("user_id", "title")
);

CREATE UNIQUE INDEX ON security.tag ("title","user_id");

CREATE UNIQUE INDEX ON security.tag ("title","user_id");

ALTER TABLE security.fail ADD FOREIGN KEY ("user_id") REFERENCES security.user ("id");

ALTER TABLE security.tag ADD FOREIGN KEY ("user_id") REFERENCES security.user ("id");
