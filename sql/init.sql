CREATE DATEBASE mop;

DROP ROLE IF EXISTS vc;
CREATE USER vc WITH PASSWORD 'admin';

DROP TABLE IF EXISTS users;
CREATE TABLE IF NOT EXISTS users (
  id SERIAL PRIMARY KEY,
  name TEXT NOT NULL,
  age INTEGER NOT NULL
);

GRANT ALL ON users TO vc;
GRANT ALL ON ALL SEQUENCES IN SCHEMA public TO vc;

SELECT * FROM users;