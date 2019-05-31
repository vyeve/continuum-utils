package postgres

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/vitaliyyevenko/continuum-utils/RMM/preparation/writer"
)

var Client writer.Writer

type client struct {
	pg *sql.DB
}

func Load() error {
	c, err := newDBClient()
	if err != nil {
		return err
	}
	Client = c
	return nil
}
