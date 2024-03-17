CREATE TABLE users (
  id INTEGER PRIMARY KEY,
  name TEXT NOT NULL
);

INSERT INTO users (id, name)
VALUES (100, 'Peter'),
  (200, 'Julia');