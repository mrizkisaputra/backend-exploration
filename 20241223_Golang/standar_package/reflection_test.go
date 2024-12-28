package standar_package

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

type LoginRequest struct {
	Username string `wajib:"true"`
	Password string `wajib:"true"`
}

func TestReflect(t *testing.T) {
	loginRequest := LoginRequest{Username: "kiki@gmail.com"}
	if err := validate(loginRequest); err != nil {
		fmt.Println(err.Error())
	}

}

func validate(obj any) error {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)
		if value.String() == "" {
			return errors.New(field.Name + "is wajib true")
		}
	}

	return nil
}
