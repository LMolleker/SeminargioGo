package database

import (
	"SeminarioGo/internal/config"
	"errors"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3" //sqlite driver support
)

//NewDataBase ...
func NewDataBase(conf *config.Config) (*sqlx.DB, error) {
	switch conf.DbCfg.Type {
	case "sqlite3":
		db, err := sqlx.Open(conf.DbCfg.Driver, ":memory:")
		if err != nil {
			return nil, err
		}
		err = db.Ping()
		if err != nil {
			return nil, err
		}
		return db, nil
	default:
		return nil, errors.New("invalid db type")
	}
}

//CreateSchema ...
func CreateSchema(s *sqlx.DB) error {
	schema := "CREATE TABLE IF NOT EXISTS 'Beer' ( " +
		"id integer PRIMARY KEY AUTOINCREMENT, " +
		"name varchar(50) NOT NULL, " +
		"alcohol_content FLOAT(2,2) NOT NULL, " +
		"price FLOAT(4,2) NOT NULL)"
	_, err := s.Exec(schema)
	if err != nil {
		return err
	}
	return nil
}
