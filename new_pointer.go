package main

import "fmt"

type Address struct{
	City, Prov, Country string
}

func main(){

	// address1 := &Address{}
	// membuat pointer kosong
	address1 := new(Address)
	address2 := address1

	address2.City = "Gresik"
	fmt.Println(address1)
	fmt.Println(address2)

}