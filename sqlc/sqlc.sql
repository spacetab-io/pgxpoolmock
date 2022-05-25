CREATE TABLE authors (
   id INTEGER PRIMARY KEY,
   name TEXT NOT NULL
);

-- name: InsertAuthors :batchone
INSERT INTO authors (name)
VALUES ($1)
RETURNING id;
