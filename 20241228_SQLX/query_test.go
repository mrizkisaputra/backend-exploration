package main_test

import (
	"database/sql"
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetAndSelect(t *testing.T) {
	type Place struct {
		Id            int            `db:"id"`
		Country       string         `db:"country"`
		City          sql.NullString `db:"city"`
		TelephoneCode sql.NullInt32  `db:"telcode"`
	}

	var total int
	err := db.Get(&total, "SELECT COUNT(*) FROM users")
	require.NoError(t, err)
	require.NotNil(t, total)
	fmt.Println(total)

	var place Place
	err = db.Get(&place, "SELECT * FROM places WHERE id = $1", 1)
	require.NoError(t, err)
	fmt.Printf("%+v\n", place)

	var names []string
	err = db.Select(&names, "SELECT name FROM users LIMIT 3")
	require.NoError(t, err)
	fmt.Printf("%+v\n", names)
}

func TestQueryRow(t *testing.T) {
	query := "SELECT * FROM places where id = $1"

	type Place struct {
		Id            int            `db:"id"`
		Country       string         `db:"country"`
		City          sql.NullString `db:"city"`
		TelephoneCode sql.NullInt32  `db:"telcode"`
	}

	place := new(Place)
	err := db.QueryRowx(query, 1).StructScan(place)
	require.NoError(t, err)

	fmt.Printf("%+v", place)
}

func TestQuery(t *testing.T) {

	query := "SELECT * FROM places"

	type Place struct {
		Id            int            `db:"id"`
		Country       string         `db:"country"`
		City          sql.NullString `db:"city"`
		TelephoneCode sql.NullInt32  `db:"telcode"`
	}

	// extension dari sqlx
	rows, err := db.Queryx(query)
	require.NoError(t, err)

	var places []Place
	for rows.Next() {
		place := new(Place)
		if err := rows.StructScan(place); err != nil {
			panic(err)
		}

		places = append(places, *place)
	}

	require.NotNil(t, places)
	fmt.Printf("%+v\n", places)

}
