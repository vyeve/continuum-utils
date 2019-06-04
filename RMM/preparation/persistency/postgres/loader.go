package postgres

import (
	"database/sql"

	"github.com/vitaliyyevenko/continuum-utils/RMM/preparation/models"
	"github.com/vitaliyyevenko/continuum-utils/RMM/preparation/writer"
)

// Client to work with Postgresql
var Client DataBase

// DataBase defines methods to work with postgreSQL
type DataBase interface {
	writer.Writer
	GetAsset(partnerID, endpointID string) (asset *models.AssetCollection, err error)
	GetByPartner(partnerID string) (assets []models.AssetCollection, err error)
	GetAll() (assets []models.AssetCollection, err error)
}
type client struct {
	pg *sql.DB
}

// Load initializes Client
func Load() error {
	c, err := newDBClient()
	if err != nil {
		return err
	}
	Client = c
	return nil
}
