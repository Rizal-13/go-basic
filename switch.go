package main

import "fmt"


func main(){

	name := "insun"

	switch name{
	case "insun":
		fmt.Println("Hello Insun")
	case "insun123":
		fmt.Println("Hello Insun")
	case "insuncrypto":
		fmt.Println("Hello Insun")
	default :
		fmt.Println("Hello Master")
	}

	switch length := len(name) ; length <= 5{
	case true:
		fmt.Println("Nama Sudah Ok !")
	case false:
		fmt.Println("Nama Terlalu Panjang")
	}

	length := len(name)
	switch{
	case length <= 5:
		fmt.Println("Nama Sudah Ok !")
	default:
		fmt.Println("Nama Terlalu Panjang")

	}
}