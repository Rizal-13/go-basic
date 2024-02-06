package main

import "fmt"

type Address struct{
	City, Prov, Country string
}

func main(){

address1 := Address{"Surabaya", "East Java", "Indonesia"}
address2 := address1

address2.City = "Gresik" // copy value tanpa merubah value awal (Pass by Value)
fmt.Println(address1)
fmt.Println(address2)

address3 := Address{"Surabaya", "East Java", "Indonesia"}
address4 := &address3

address4.City = "Gresik" // pointer (apabila ada perubahan pada value yang di pointer maka akan merubah value awal) Pass by Reference
fmt.Println(address3)
fmt.Println(address4)



}