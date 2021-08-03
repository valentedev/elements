package data

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"gopkg.in/guregu/null.v3"
)

type Elements struct {
	AtomicNumber      null.String `json:"atomic_number,omitempty"`
	Symbol            null.String `json:"symbol,omitempty"`
	Name              null.String `json:"name"`
	OriginOfName      null.String `json:"origin_of_name,omitempty"`
	Period            null.String `json:"period,omitempty"`
	Groups            null.String `json:"groups,omitempty"`
	Block             null.String `json:"block,omitempty"`
	AtomicWeight      null.String `json:"atomic_weight"`
	Density           null.String `json:"density_g_cm3,omitempty"`
	MeltingPoint      null.String `json:"melting_point,omitempty"`
	BoilingPoint      null.String `json:"boiling_point,omitempty"`
	SpecificHeat      null.String `json:"specific_heat_j_g,omitempty"`
	ElectroNegativity null.String `json:"electro_negativity,omitempty"`
	AbundanceInEarth  null.String `json:"abundance_in_earth_mg_kg,omitempty"`
	Origin            null.String `json:"origin,omitempty"`
	Phase             null.String `json:"phase_at_room_temperature,omitempty"`
}

type ElementsModel struct {
	DB *sql.DB
}

func (m ElementsModel) GetAll(name, atomic_weight string, filters Filters) ([]*Elements, Metadata, error) {
	query := fmt.Sprintf(`
        SELECT count(*) OVER(), atomic_number, symbol, name, origin_of_name, period, groups, block, atomic_weight, density_g_cm3, melting_point, boiling_point, specific_heat_j_g, electro_negativity, abundance_in_earth_mg_kg, origin, phase_at_room_temperature
        FROM elements_table
        ORDER BY %s %s, name ASC
        LIMIT $1 OFFSET $2`, filters.sortColumn(), filters.sortDirection())

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	args := []interface{}{filters.limit(), filters.offset()}

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
			&element.Period,
			&element.Groups,
			&element.Block,
			&element.AtomicWeight,
			&element.Density,
			&element.MeltingPoint,
			&element.BoilingPoint,
			&element.SpecificHeat,
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
