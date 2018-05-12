package main

import (
	"fmt"

	"github.com/faruqisan/go-request-validation/requests"
	"github.com/faruqisan/go-request-validation/validation/validator"
)

func main() {

	validator := validator.NewValidator()

	failUserRequest := requests.CreateUserRequest{
		Name:     "",
		Email:    "wrongemail@",
		Password: "",
	}

	err := validator.ValidateStruct(failUserRequest)
	if err != nil {
		fmt.Println(err)
	}

	passUserRequest := requests.CreateUserRequest{
		Name:     "John",
		Email:    "john@gmail.com",
		Password: "somepassword",
	}

	err = validator.ValidateStruct(passUserRequest)
	if err != nil {
		fmt.Println(err)
	}
}
