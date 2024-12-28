package model

import "time"

type User struct {
	Id        int
	Name      string
	Email     string
	Age       int
	CreatedAt time.Time
}
