package postgres

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/vitaliyyevenko/continuum-utils/RMM/preparation/models"
)

// ErrNotFound error if cannot find asset by partnerID and endpointID
var ErrNotFound = errors.New("asset not found")

const (
	insertIntoCollection = `
	INSERT INTO asset (endpointID, partnerID, rawAsset) 
	VALUES ($1, $2, $3)
		ON CONFLICT (endpointID) DO UPDATE 
  		SET	partnerID = EXCLUDED.partnerID,
     		rawAsset = EXCLUDED.rawAsset;`
	selectAsset = `
	SELECT rawAsset 
	FROM asset
	WHERE endpointID=$1 AND partnerID=$2;`
	selectByPartner = `
	SELECT rawAsset 
	FROM asset
	WHERE partnerID=$1;`
	selectAll = `
	SELECT rawAsset 
	FROM asset;`
)

func newDBClient() (DataBase, error) {
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

func (c client) GetAsset(partnerID, endpointID string) (asset *models.AssetCollection, err error) {
	var rawAsset string
	row := c.pg.QueryRow(selectAsset, endpointID, partnerID)
	err = row.Scan(&rawAsset)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		}
		return nil, err
	}
	asset = new(models.AssetCollection)
	err = json.Unmarshal([]byte(rawAsset), asset)
	return
}

func (c client) GetByPartner(partnerID string) (assets []models.AssetCollection, err error) {
	rows, err := c.pg.Query(selectByPartner, partnerID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		}
		return nil, err
	}
	for rows.Next() {
		var rawAsset string
		err = rows.Scan(&rawAsset)
		if err != nil {
			if err == sql.ErrNoRows {
				continue
			}
			return nil, err
		}
		asset := models.AssetCollection{}
		err = json.Unmarshal([]byte(rawAsset), &asset)
		if err != nil {
			return nil, err
		}
		assets = append(assets, asset)
	}
	if len(assets) == 0 {
		return nil, ErrNotFound
	}
	return assets, nil
}

func (c client) GetAll() (assets []models.AssetCollection, err error) {
	rows, err := c.pg.Query(selectAll)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		}
		return nil, err
	}
	for rows.Next() {
		var rawAsset string
		err = rows.Scan(&rawAsset)
		if err != nil {
			if err == sql.ErrNoRows {
				continue
			}
			return nil, err
		}
		asset := models.AssetCollection{}
		err = json.Unmarshal([]byte(rawAsset), &asset)
		if err != nil {
			return nil, err
		}
		assets = append(assets, asset)
	}
	if len(assets) == 0 {
		return nil, ErrNotFound
	}
	return assets, nil
}
