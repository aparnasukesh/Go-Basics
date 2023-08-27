package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type UserData struct {
	Username string `json:"username" validate:"required,min=2,max=10"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Age      int    `json:"age" validate:"required,gte=13,lte=99"`
}

func main() {
	userData := UserData{
		Username: "johndoe",
		Email:    "johndoe@gmail.com",
		Password: "123",
		Age:      15,
	}

	validate := validator.New()
	err := validate.Struct(userData)

	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, e := range validationErrors {

			fmt.Printf("Field: %s, Tag: %s\n", e.Field(), e.Tag())
			if e.Field() == "Email" {
				fmt.Println("Ivalide Email")
			} else if e.Field() == "Age" {
				fmt.Println("Invalide Age")
			} else if e.Field() == "Password" {
				fmt.Println("invalide password")
			} else if e.Field() == "Username" {
				fmt.Println("Invalide Username")
			}

		}
	} else {
		fmt.Println("Validation passed.")
	}
}
