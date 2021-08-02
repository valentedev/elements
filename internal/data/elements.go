package data

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

type Elements struct {
	AtomicNumber      string `json:"atomic_number"`
	Symbol            string `json:"symbol"`
	Name              string `json:"name"`
	OriginOfName      string `json:"origin_of_name"`
	Periodo           string `json:"periodo"`
	Grupo             string `json:"grupo"`
	Block             string `json:"block"`
	AtomicWeight      string `json:"atomic_weight"`
	Density           string `json:"density_g_cm3"`
	MeltingPoint      string `json:"melting_point"`
	BoilingPoint      string `json:"boiling_point"`
	SpecificHeat      string `json:"specific_heat_j_g"`
	ElectroNegativity string `json:"electro_negativity"`
	AbundanceInEarth  string `json:"abundance_in_earth_mg_kg"`
	Origin            string `json:"origin"`
	Phase             string `json:"phase"`
}

type ElementsModel struct {
	DB *sql.DB
}

func (m ElementsModel) GetAll(name, atomic_weight string, filters Filters) ([]*Elements, Metadata, error) {
	query := fmt.Sprintf(`
        SELECT count(*) OVER(), name, atomic_weight
        FROM elements_table
        WHERE (to_tsvector('simple', name) @@ plainto_tsquery('simple', $1) OR $1 = '') 
        AND (genres @> $2 OR $2 = '{}')     
        ORDER BY %s %s, id ASC
        LIMIT $3 OFFSET $4`, filters.sortColumn(), filters.sortDirection())

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	args := []interface{}{name, atomic_weight, filters.limit(), filters.offset()}

	rows, err := m.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, Metadata{}, err
	}

	defer rows.Close()

	totalRecords := 0
	elements := []*Elements{}

	for rows.Next() {
		var element Elements

		err := rows.Scan(
			&totalRecords,
			&element.AtomicNumber,
			&element.Symbol,
			&element.Name,
			&element.OriginOfName,
			&element.Periodo,
			&element.Grupo,
			&element.Block,
			&element.AtomicWeight,
			&element.Density,
			&element.MeltingPoint,
			&element.BoilingPoint,
			&element.ElectroNegativity,
			&element.AbundanceInEarth,
			&element.Origin,
			&element.Phase,
		)

		if err != nil {
			return nil, Metadata{}, err
		}

		elements = append(elements, &element)
	}

	if err = rows.Err(); err != nil {
		return nil, Metadata{}, err
	}

	metadata := calculateMetadata(totalRecords, filters.Page, filters.PageSize)

	return elements, metadata, nil
}
