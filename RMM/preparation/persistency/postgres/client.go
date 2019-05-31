package postgres

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/vitaliyyevenko/continuum-utils/RMM/preparation/models"
	"github.com/vitaliyyevenko/continuum-utils/RMM/preparation/writer"
)

/*
CREATE TABLE IF NOT EXISTS collection (
  createTimeUTC timestamp,
  agentInstalledUTC timestamp,
  createdBy text,
  endpointID serial NOT NULL UNIQUE,
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
*/
const (
	insertIntoCollection = `insert into collection (createTimeUTC, agentInstalledUTC, createdBy, endpointID, name, type, partnerID, clientID,
		siteID, friendlyName, remoteAddress, resourceType, endpointType) values ($1, $2, $3, 44, $5, $6, $7, $8, $9, $10, $11, $12, $13)`
)

func newDBClient() (writer.Writer, error) {
	connStr := "user=postgres dbname=endpoints sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return client{pg: db}, nil
}

func (c client) WriteAll(assets []models.AssetCollection) (err error) {
	return
}

func (c client) Write(a models.AssetCollection) (err error) {
	_, err = c.pg.Exec(insertIntoCollection, a.CreateTimeUTC, a.AgentInstalledUTC, a.CreatedBy, a.EndpointID, a.Name, a.Type,
		a.PartnerID, a.ClientID, a.SiteID, a.FriendlyName, a.RemoteAddress, a.ResourceType, a.EndpointType)
	return err
}
