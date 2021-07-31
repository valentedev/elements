package data

import "database/sql"

type Database struct{}

type DatabaseModel struct {
	DB *sql.DB
}
