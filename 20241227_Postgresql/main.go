package main

import (
	"context"
	"fmt"
	"golang-postgresql/cfg"
	"golang-postgresql/repositories"
	"log"
)

func main() {

	db, err := cfg.NewPostgresConn()
	if err != nil {
		log.Fatal(err)
	}

	userRepo := repositories.NewUserRepository(db)

	//err = userRepo.Insert(context.Background(), model.User{
	//	Name:  "mrizkisaputra",
	//	Email: "mrizkisaputra@gmail.com",
	//	Age:   23,
	//})
	//if err != nil {
	//	log.Fatal(err)
	//}

	users, err := userRepo.FindAll(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(users)

}
