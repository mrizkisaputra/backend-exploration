package main_test

import (
	"belajar-sqlx/pkg/db/postgres"
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

var db, err = postgres.NewPostgresConn()

func TestExecStandar(t *testing.T) {
	query := "INSERT INTO places (country, city, telcode) VALUES($1, $2, $3)"

	// standar package database/sql
	result, err := db.Exec(query, "Indonesia", "", 62)
	require.NoError(t, err)
	fmt.Println(result.LastInsertId())
}

func TestExecExtensionSqlx(t *testing.T) {
	query := "INSERT INTO place (country, city, telcode) VALUES($1, $2, $3)"

	// extension dari sqlx
	result := db.MustExec(query, "Indonesia", "Palembang", 62)
	fmt.Println(result.LastInsertId())
	fmt.Println(result.RowsAffected())
}
