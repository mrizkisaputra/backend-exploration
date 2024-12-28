package main_test

import (
	"database/sql"
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNamedQuery(t *testing.T) {
	type Place struct {
		Id            int            `db:"id"`
		Country       string         `db:"country"`
		City          sql.NullString `db:"city"`
		TelephoneCode sql.NullInt32  `db:"telcode"`
	}

	query := "SELECT * FROM places WHERE country = :country"
	rows, err := db.NamedQuery(query, Place{Country: "Indonesia"})
	require.NoError(t, err)

	var places []Place
	for rows.Next() {
		var place Place
		rows.StructScan(&place)
		places = append(places, place)
	}
	fmt.Printf("%+v\n", places)
}
