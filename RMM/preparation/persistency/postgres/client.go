package postgres

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/vitaliyyevenko/continuum-utils/RMM/preparation/models"
	"github.com/vitaliyyevenko/continuum-utils/RMM/preparation/writer"
)

const (
	insertIntoCollection = `insert into asset (endpointID, partnerID, rawAsset) 
	values ($1, $2, $3)`
)

func newDBClient() (writer.Writer, error) {
	connStr := "user=postgres dbname=dg sslmode=disable password=1488"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return client{pg: db}, nil
}

func (c client) WriteAll(assets []models.AssetCollection) (err error) {
	for _, asset := range assets {
		errOne := c.Write(asset)
		if errOne != nil {
			if err == nil {
				err = errOne
			} else {
				err = fmt.Errorf("%s; %s", err, errOne)
			}
		}
	}
	return err
}

func (c client) Write(asset models.AssetCollection) (err error) {
	p, err := json.Marshal(&asset)
	if err != nil {
		return err
	}
	_, err = c.pg.Exec(insertIntoCollection, asset.EndpointID, asset.PartnerID, string(p))
	return err
}
