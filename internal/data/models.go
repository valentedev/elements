package data

import (
	"database/sql"
)

type Models struct {
	Database DatabaseModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Database: DatabaseModel{DB: db},
	}
}
