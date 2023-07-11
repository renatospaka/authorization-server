CREATE TABLE IF NOT EXISTS authorizations (
  id                UUID, 
  client_id         UUID, 
  transaction_id    UUID, 
  status      varchar(15), 
  value       numeric(8, 2), 
  approved_at timestamp, 
  denied_at   timestamp, 
  created_at  timestamp, 
  updated_at  timestamp, 
  deleted_at  timestamp,
  PRIMARY KEY (id)
);

ALTER TABLE authorizations
ALTER COLUMN value SET DATA TYPE numeric(8, 2);
