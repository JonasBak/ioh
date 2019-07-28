CREATE DATABASE ioh;

\c ioh

CREATE TABLE clients (
  id VARCHAR(8) PRIMARY KEY
);
CREATE TABLE configs (
  id SERIAL PRIMARY KEY,
  plant VARCHAR(24),
  water INT,
  clientid VARCHAR(8) UNIQUE NOT NULL,
  timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE configs
ADD FOREIGN KEY (clientid) REFERENCES clients(id);

GRANT ALL PRIVILEGES ON TABLE clients TO postgres;
GRANT ALL PRIVILEGES ON TABLE configs TO postgres;
GRANT USAGE, SELECT ON ALL SEQUENCES IN SCHEMA public TO postgres;

\dt
\d clients
\d configs
