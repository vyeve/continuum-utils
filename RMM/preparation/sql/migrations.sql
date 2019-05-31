CREATE DATABASE endpoints;
-- psql -d endpoints -U postgres
CREATE TABLE IF NOT EXISTS collection (
  createTimeUTC timestamp,
  agentInstalledUTC timestamp,
  createdBy text,
  endpointID text NOT NULL UNIQUE,
  name text,
  type text,
  partnerID text,
  clientID text,
  siteID text,
  friendlyName text,
  remoteAddress text,
  resourceType text,
  endpointType text,
  PRIMARY KEY (endpointID)
);