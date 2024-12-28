package main

import "belajar-sqlx/pkg/db/postgres"

func main() {
	_, err := postgres.NewPostgresConn()
	if err != nil {
		panic(err)
	}

}
