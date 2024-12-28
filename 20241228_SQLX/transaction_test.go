package main_test

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestTransaction(t *testing.T) {
	tx := db.MustBegin()
	defer tx.Rollback()

	tx.MustExec("INSERT INTO places(country, city, telcode) values ($1, $2, $3)", "India", "Delhi", 63)

	// transaksi lainnya ....

	err := tx.Commit()
	require.NoError(t, err)
}
