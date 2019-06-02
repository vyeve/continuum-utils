CREATE TABLE IF NOT EXISTS asset (
  endpointID varchar(256) NOT NULL UNIQUE,
  partnerID varchar(32) NOT NULL,
  rawAsset json,
  PRIMARY KEY (endpointID)
);
COPY asset FROM '/migrations/assets.csv' WITH CSV DELIMITER ',' HEADER