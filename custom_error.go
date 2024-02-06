package main

import "fmt"

type validationError struct {
	Message string
}

func (v *validationError) Error() string {
	return v.Message
}

type notFoundError struct {
	Message string
}

func (v *notFoundError) Error() string {
	return v.Message
}

func SaveData(id string, data any) error {
	if id == "" {
		return &validationError{Message: "validation error"}
	}
	if id != "insun" {
		return &notFoundError{Message: "data not found"}
	}

	return nil

}

func main() {

	err := SaveData("inasun", nil)
	// fmt.Println(err)

	if err != nil{
		if validationErr, ok := err.(*validationError); ok{
			fmt.Println("validation error:", validationErr.Message)
		}else if validationErr, ok := err.(*notFoundError); ok{
			fmt.Println("not found error:", validationErr.Message)
		} else {
			fmt.Println("unknown error:", err.Error())
		}
	}
}