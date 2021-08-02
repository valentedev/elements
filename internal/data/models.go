package data

import (
	"database/sql"
)

type Models struct {
	Elements ElementsModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Elements: ElementsModel{DB: db},
	}
}
