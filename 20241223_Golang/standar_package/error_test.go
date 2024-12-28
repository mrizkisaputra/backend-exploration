package standar_package

import (
	"fmt"
	"github.com/pkg/errors"
	"testing"
)

type CustomError struct {
	message string
}

func NewCustomError(message string) *CustomError {
	return &CustomError{message: message}
}

func (e *CustomError) Error() string {
	return e.message
}

var (
	BadRequestError     = errors.New("Bad Request")
	NotFoundError       = errors.New("Not Found")
	InternalServerError = errors.New("Internal Server Error")
)

func TestError1(t *testing.T) {
	err := login("", "")
	if err != nil {
		switch {
		case errors.Is(err, BadRequestError):
			fmt.Println(err.Error())
		case errors.Is(err, NotFoundError):
			fmt.Println(err.Error())
		case errors.Is(err, InternalServerError):
			fmt.Println(err.Error())
		default:
			fmt.Println("Unknown error")
		}
	}
}

func TestError2(t *testing.T) {
	err := doSomething()
	if err != nil {
		var customError *CustomError
		if errors.As(err, &customError) {
			fmt.Println(err.Error())
		}
	}
}

func doSomething() error {
	return NewCustomError("Something error")
}

func login(email, password string) error {
	if email == "" || password == "" {
		return fmt.Errorf("error nih: %w", BadRequestError)
	}
	return nil
}
