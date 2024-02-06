package main

import "fmt"

type Address struct{
	City, Prov, Country string
}

func ChangeAddress(address *Address){
	address.Country = "Japan"
}

func main(){
	address := &Address{}
	ChangeAddress(address)
	
	// address := Address{}
	// ChangeAddress(&address)

	fmt.Println(address)
}